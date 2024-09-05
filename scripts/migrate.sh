source scripts/lib.sh || { echo "Are you at repo root?"; exit 1; }

usage() {
  cat <<EOUSAGE
Usage: $0 [up|down|force|version] {#}"
EOUSAGE
}

database_host="localhost"
if [[ $WINGRAM_DB_HOST != "" ]]; then
  database_host=$WINGRAM_DB_HOST
fi
database_port="5432"
if [[ $WINGRAM_DB_PORT != "" ]]; then
  database_port=$WINGRAM_DB_PORT
fi
database_user="testuser"
if [[ $WINGRAM_DB_USER != "" ]]; then
  database_user=$WINGRAM_DB_USER
fi
database_name='testdb'
if [[ $WINGRAM_DB_NAME != "" ]]; then
  database_name=$WINGRAM_DB_NAME
fi
database_password="testpassword"
if [[ $WINGRAM_DB_PASS != "" ]]; then
  database_password=$WINGRAM_DB_PASS
fi
ssl_mode='disable'
if [[ $WINGRAM_DB_SSL != "" ]]; then
  ssl_mode=$WINGRAM_DB_SSL
fi

# Redirect stderr to stdout because migrate outputs to stderr, and we want
# to be able to use ordinary output redirection.
case "$1" in
  up|down|force|version)
    migrate \
      -source file:migrations \
      -database "postgresql://$database_user:$database_password@$database_host:$database_port/$database_name?sslmode=$ssl_mode" \
      "$@" 2>&1
    ;;
  *)
    usage
    exit 1
    ;;
esac
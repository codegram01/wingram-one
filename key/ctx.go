package key

type CtxKey string

const (
	CtxIdentity CtxKey = "identity"
	CtxAuthorization CtxKey = "Authorization"
)
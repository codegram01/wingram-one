package crypto

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	password := "Vip1OnlyCreate"

	hashPass, err := Hash(password)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(hashPass)
}

func TestCheckHash(t *testing.T) {
	password := "Vip1OnlyCreate"
	passwordHash := "$2a$10$dvMyvcF9t/L8tGt/bx0Q.ugKq4QLcQyJvdSJcKALLhTaWwnY7iwFy"

	val := CheckHash(password, passwordHash)

	fmt.Println(val)
}

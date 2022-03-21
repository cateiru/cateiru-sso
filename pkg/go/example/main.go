package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cateiru/cateiru-sso/pkg/go/sso"
)

const CLIENT_ID = "6a13bdd63607a796d7f3e1e3c52c3d"
const TOKEN_SECRET = "03634b89e83f9d711dcf55952d51f9256ba92ee035d18004160d0f78ea4c8939"

func main() {
	loginUri := sso.CreateURI(CLIENT_ID, "direct")

	fmt.Printf("----> %s\n", loginUri)
	fmt.Print("code: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	code := scanner.Text()

	// get token
	token, err := sso.GetToken(code, "direct", TOKEN_SECRET)
	if err != nil {
		panic(err)
	}

	claims, err := sso.ValidateIDToken(token.IDToken)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi! %s\n", claims.NickName)

	// claimsJ, err := json.Marshal(claims)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(claimsJ))
}

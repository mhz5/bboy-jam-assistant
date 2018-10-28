// Invocation: go run generate_secure_cookiestore_key.go
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gorilla/securecookie"
)

const (
	sessionSecretKeyFilename = "/tmp/session_secret_key"
)

func main() {
	secret := securecookie.GenerateRandomKey(32)
	err := ioutil.WriteFile(sessionSecretKeyFilename, secret, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Saved session store secret key to " + sessionSecretKeyFilename)
}

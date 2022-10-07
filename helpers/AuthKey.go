package helpers

import (
	"context"
	"errors"
	"fmt"
	jwt "github.com/golang-jwt/jwt"
	jwk "github.com/lestrrat-go/jwx/jwk"
)

func GetKey(token *jwt.Token) (interface{}, error) {

	// For a demonstration purpose, Google Sign-in is used.
	// https://developers.google.com/identity/sign-in/web/backend-auth
	//
	// This user-defined KeyFunc verifies tokens issued by Google Sign-In.
	//
	// Note: In this example, it downloads the keyset every time the restricted route is accessed.
	keySet, err := jwk.Fetch(context.Background(), "https://www.googleapis.com/oauth2/v3/certs")
	if err != nil {
		return nil, err
	}

	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have a key ID in the kid field")
	}

	key, found := keySet.LookupKeyID(keyID)

	if !found {
		return nil, fmt.Errorf("unable to find key %q", keyID)
	}

	var pubkey interface{}
	if err := key.Raw(&pubkey); err != nil {
		return nil, fmt.Errorf("unable to get the public key. Error: %s", err.Error())
	}

	return pubkey, nil
}

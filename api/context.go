package api

import (
	"encoding/base64"
	"flag"
	"net/http"

	auth0 "github.com/auth0-community/go-auth0"
	jose "gopkg.in/square/go-jose.v2"
)

var (
	// JWKS_URI            = flag.String("jwks_url", "https://samples.auth0.com/.well-known/jwks.json", "")
	AUTH0_API_ISSUER    = flag.String("auth0_api_issuer", "https://samples.auth0.com/", "")
	AUTH0_API_CLIENT    = flag.String("auth0_api_client", "kbyuFDidLLm280LIwVFiazOqjO3ty8KH", "")
	AUTH0_CLIENT_SECRET = flag.String("auth0_client_secret", "60Op4HFM0I8ajz0WdiStAbziZ-VFQttXuxixHHs2R7r7-CW8GR79l-mmLqMhc-Sa", "")
)

type Context struct {
	JWTValidator *auth0.JWTValidator
	OIDCClient   *http.Client
}

func NewContext() (*Context, error) {
	var c = Context{}
	var err error

	// create jwt validator
	secret, _ := base64.URLEncoding.DecodeString(*AUTH0_CLIENT_SECRET)
	configuration := auth0.NewConfiguration(auth0.NewKeyProvider(secret), []string{*AUTH0_API_CLIENT}, *AUTH0_API_ISSUER, jose.HS256)
	c.JWTValidator = auth0.NewValidator(configuration)

	// create http client to get userinfo
	c.OIDCClient = http.DefaultClient

	return &c, err
}

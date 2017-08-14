package api

import (
	"fmt"
	"net/http"
	"strings"

	"bluecanvas.io/auth0-api/models"
	auth0 "github.com/auth0-community/go-auth0"
	"github.com/rcrowley/go-tigertonic"
	jwt "gopkg.in/square/go-jose.v2/jwt"
)

func wrap(i interface{}, metricsName string) http.Handler {
	return tigertonic.Timed(tigertonic.Marshaled(i), metricsName, nil)
}

func authWrap(c *Context, i interface{}, metricsName string) http.Handler {
	return tigertonic.First(&AuthMiddleware{c}, wrap(i, metricsName))
}

type AuthMiddleware struct {
	Context *Context
}

func (m AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := m.Context.checkJWT(r); err != nil {
		tigertonic.ResponseErrorWriter.WriteError(r, w, err)
	}
}

func (c *Context) checkJWT(r *http.Request) error {
	_, err := c.JWTValidator.ValidateRequest(r)
	if err != nil {
		fmt.Println(err)
		return models.NewHTTPError(http.StatusUnauthorized, "missing or invalid token")
	}
	return nil
}

func checkScope(r *http.Request, validator *auth0.JWTValidator, token *jwt.JSONWebToken) bool {
	claims := map[string]interface{}{}
	err := validator.Claims(r, token, &claims)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if strings.Contains(claims["scope"].(string), "read:messages") {
		return true
	} else {
		return false
	}
}

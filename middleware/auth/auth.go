package auth

import (
	"crypto/subtle"
	"net/http"

	"github.com/ssoor/webapi"
)

const (
	AuthenticatedUsername = "authenticated_username"
)

// Equal performs a constant time compare of two strings to limit timing attacks.
func Equal(v1 string, v2 string) bool {
	return subtle.ConstantTimeCompare([]byte(v1), []byte(v2)) == 1
}

// Credential represents a credential information.
type Credential struct {
	Username string
	Password string
}

// New returns a new HTTP basic authorization middleware.
func New(credentials ...*Credential) *Middleware {
	instance := &Middleware{
		credentials: make(map[string]string),
	}

	for _, v := range credentials {
		instance.credentials[v.Username] = v.Password
	}

	return instance
}

// Middleware represents a HTTP basic authorization middleware.
type Middleware struct {
	credentials        map[string]string         // key: username, value: password
	AuthenticationFunc func(string, string) bool // authentication function for authenticating
}

func (this Middleware) OnRequest(req *http.Request, writer *http.ResponseWriter) *http.ResponseWriter {
	var authenticated bool
	username, password, ok := req.BasicAuth()
	if ok {
		if len(username) > 0 {
			if v, ok := this.credentials[username]; ok {
				if Equal(v, password) {
					authenticated = true
				}
			}

			if !authenticated && this.AuthenticationFunc != nil {
				authenticated = this.AuthenticationFunc(username, password)
			}
		}
	}

	if !authenticated {
		writer.Write(http.StatusUnauthorized)
		writer.Header().Set("WWW-Authenticate", "Basic realm=\"Authorization Required\"")
	} else {
		// set username in the context, the username can be read later using context.Data[authentication.AuthenticatedUsername]
		context.Data[AuthenticatedUsername] = username
	}
}

func (this Middleware) OnResponse(req *http.Request, writer *http.ResponseWriter) *http.ResponseWriter {
	// do nothing
}

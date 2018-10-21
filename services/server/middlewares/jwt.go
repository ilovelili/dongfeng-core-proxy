package middleware

import (
	"net/http"

	"github.com/boj/redistore"
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	"github.com/ilovelili/dongfeng-shared-lib"
)

// sessionkey string used as Redis session store key
const sessionkey = "blacklisttoken"

var (
	jwks         string
	sessionstore *redistore.RediStore
)

func init() {
	config := utils.GetConfig()
	jwks = config.JWKS
	sessionstore, _ = redistore.NewRediStore(config.Redis.GetMaxConnectionCount(), "tcp", config.Redis.Host, config.Redis.Password, []byte("session-store"))
}

// JwtAuthenticate JWT auth middleware. go-restful has poor support for middleware injection
func JwtAuthenticate(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain) {
	idtoken, valid := utils.ResolveIDToken(req)
	if !valid {
		rsp.WriteErrorString(http.StatusUnauthorized, "401: Not Authorized")
		return
	}

	// check if idtoken in blacklist
	session, err := sessionstore.Get(req.Request, "login-session")
	if err != nil {
		rsp.WriteError(http.StatusInternalServerError, err)
		return
	}
	flashes := session.Flashes(sessionkey)

	// token in blacklist
	if containString(flashes, idtoken) {
		rsp.WriteErrorString(http.StatusUnauthorized, "401: Not Authorized")
		return
	}

	// parse and validate the token
	_, token, err := sharedlib.ParseJWT(idtoken, jwks)
	if err != nil || !token.Valid {
		rsp.WriteErrorString(http.StatusUnauthorized, "401: Not Authorized")
		return
	}

	chain.ProcessFilter(req, rsp)
}

// containString slice contains specified string element or not
func containString(s []interface{}, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

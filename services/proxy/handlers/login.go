package handlers

import (
	"github.com/boj/redistore"
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng/sharedlib"
	"github.com/micro/go-micro/metadata"

	api "github.com/ilovelili/dongfeng/core-proxy/services/proto"
	"github.com/ilovelili/dongfeng/core-proxy/services/utils"
)

// sessionkey string used as Redis session store key
const sessionkey = "blacklisttoken"

var (
	config       *utils.Config
	sessionstore *redistore.RediStore
)

func init() {
	config = utils.GetConfig()
	sessionstore, _ = redistore.NewRediStore(config.Redis.GetMaxConnectionCount(), "tcp", config.Redis.Host, config.Redis.Password, []byte("session-store"))
}

// Login login
func Login(req *restful.Request, rsp *restful.Response) {
	idtoken, _ := utils.ResolveIDToken(req)
	ip := sharedlib.ResolveRemoteIP(req.Request)
	jwks := config.Auth.JWKS

	// Set arbitrary headers in context
	ctx := metadata.NewContext(req.Request.Context(), map[string]string{
		sharedlib.MetaDataToken: idtoken,
		sharedlib.MetaDataIP:    ip,
		sharedlib.MetaDataJwks:  jwks,
	})

	response, err := newclient().Login(ctx, &api.LoginRequest{})
	if err != nil {
		rsp.WriteEntity(err)
	} else {
		rsp.WriteEntity(response)
	}
}

// Logout logout
func Logout(req *restful.Request, rsp *restful.Response) {
	idtoken, _ := utils.ResolveIDToken(req)

	// save token to blacklist
	// Get a session.
	session, err := sessionstore.Get(req.Request, "login-session")
	if err != nil {
		rsp.WriteError(500, err)
		return
	}

	// Add flash
	session.AddFlash(idtoken, sessionkey)
	if err = session.Save(req.Request, rsp.ResponseWriter); err != nil {
		rsp.WriteError(500, err)
	} else {
		rsp.Write([]byte("OK"))
	}
}

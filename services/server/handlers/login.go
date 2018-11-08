package handlers

import (
	"github.com/boj/redistore"
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	protobuf "github.com/ilovelili/dongfeng-protobuf"
)

// sessionkey string used as Redis session store key
const sessionkey = "blacklisttoken"

var (
	sessionstore *redistore.RediStore
)

func init() {
	sessionstore, _ = redistore.NewRediStore(config.Redis.GetMaxConnectionCount(), "tcp", config.Redis.Host, config.Redis.Password, []byte("session-store"))
}

// Login login
func Login(req *restful.Request, rsp *restful.Response) {
	response, err := newclient().Login(ctx(req), &protobuf.LoginRequest{})
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
		writeError(rsp, errorcode.CoreProxyFailedToGetSession)
		return
	}

	// Add flash
	session.AddFlash(idtoken, sessionkey)
	if err = session.Save(req.Request, rsp.ResponseWriter); err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToSaveSession)
	} else {
		rsp.Write([]byte("OK"))
	}
}

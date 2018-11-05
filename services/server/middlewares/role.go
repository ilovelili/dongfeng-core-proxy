package middlewares

import (
	"encoding/json"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
)

// AdminRoleAuthenticate requires admin role
func AdminRoleAuthenticate(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain) {
	AdminRole.roleAuthenticate(req, rsp, chain)
}

// UserRoleAuthenticate requires user role
func UserRoleAuthenticate(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain) {
	UserRole.roleAuthenticate(req, rsp, chain)
}

// Role role restriction for each endpoint
type Role []string

type user struct {
	RoleSettings string `json:"https://moneyhatch.com/role"`
}

var (
	// AdminRole admin role
	AdminRole Role = []string{"admin"}
	// UserRole user role
	UserRole Role = []string{"user", ""}
)

// JwtAuthenticate JWT auth middleware. go-restful has poor support for middleware injection
func (r *Role) roleAuthenticate(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain) {
	idtoken, valid := utils.ResolveIDToken(req)
	if !valid {
		writeError(rsp, errorcode.GenericNotAuthorized)
		return
	}

	claims, token, err := sharedlib.ParseJWT(idtoken, jwks)
	if err != nil || !token.Valid {
		writeError(rsp, errorcode.GenericNotAuthorized)
		return
	}

	// Unmarshal user info
	userprofilestream, _ := json.Marshal(claims)
	var userprofile *user
	err = json.Unmarshal(userprofilestream, &userprofile)
	if err != nil {
		writeError(rsp, errorcode.GenericInsufficientPrivileges)
		return
	}

	if !containString(r.resolveRole(), userprofile.RoleSettings) {
		writeError(rsp, errorcode.GenericInsufficientPrivileges)
		return
	}

	chain.ProcessFilter(req, rsp)
}

func (r *Role) resolveRole() []interface{} {
	result := make([]interface{}, 0)
	for _, role := range *r {
		result = append(result, role)
	}

	return result
}

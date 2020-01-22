package utils

import (
	"strings"

	restful "github.com/emicklei/go-restful"
)

// ResolveHeaderInfo resolve id_token saved in header
func ResolveHeaderInfo(req *restful.Request) (idtoken, pid, email string, valid bool) {
	valid = true
	idtokensegments := strings.Split(req.HeaderParameter("Authorization"), "Bearer ")
	if len(idtokensegments) != 2 {
		valid = false
		return
	}

	idtoken = idtokensegments[1]
	pid = req.HeaderParameter("x-pid")
	email = req.HeaderParameter("x-email")
	return
}

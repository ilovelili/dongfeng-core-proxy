package utils

import (
	"strings"

	restful "github.com/emicklei/go-restful"
)

// ResolveIDToken resolve id_token saved in header
func ResolveIDToken(req *restful.Request) (idtoken string, valid bool) {
	valid = true
	idtokensegments := strings.Split(req.HeaderParameter("Authorization"), "Bearer ")
	if len(idtokensegments) != 2 {
		valid = false
		return
	}

	idtoken = idtokensegments[1]
	return
}

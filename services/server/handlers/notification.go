package handlers

import (
	"encoding/json"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// UpdateNotificationsRequest notification update request
type UpdateNotificationsRequest struct {
	IDs []int64 `json:"ids"`
}

// UpdateNotifications update notifications
func UpdateNotifications(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *UpdateNotificationsRequest
	err := decoder.Decode(&updatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidUpdateNotificationRequestBody)
		return
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateNotifications(ctx(req), &proto.UpdateNotificationsRequest{
		Token:         idtoken,
		Notifications: updatereq.IDs,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

package handlers

import (
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetMenus get menus
func GetMenus(req *restful.Request, rsp *restful.Response) {
	junior_or_senior, breakfast_or_lunch, from, to := req.QueryParameter("junior_or_senior"), req.QueryParameter("breakfast_or_lunch"), req.QueryParameter("from"), req.QueryParameter("to")
	if from > to {
		writeError(rsp, errorcode.CoreProxyInvalidGetMenuRequest)
		return
	}

	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetMenus(ctx(req), &proto.GetMenuRequest{
		Pid:              pid,
		From:             from,
		To:               to,
		BreakfastOrLunch: resolveBreakfastOrLunch(breakfast_or_lunch),
		JuniorOrSenior:   resolveJuniorOrSenior(junior_or_senior),
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

func resolveJuniorOrSenior(junior_or_senior string) (result int64) {
	result = -1

	if junior_or_senior == "junior" {
		result = 0
	} else if junior_or_senior == "senior" {
		result = 1
	}

	return
}

func resolveBreakfastOrLunch(breakfast_or_lunch string) (result int64) {
	result = -1

	if breakfast_or_lunch == "snack" {
		result = 2
	} else if breakfast_or_lunch == "lunch" {
		result = 1
	} else if breakfast_or_lunch == "breakfast" {
		result = 0
	}

	return
}

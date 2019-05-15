package handlers

import (
	"encoding/json"
	"strings"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// UpdateIngredientRequest ingredient update request
type UpdateIngredientRequest struct {
	Ingredient        string  `json:"ingredient" csv:"ingredient"`
	Alias             string  `json:"alias" csv:"alias"`
	Protein100g       float64 `json:"protein_100g" csv:"protein_100g"`
	ProteinDaily      float64 `json:"protein_daily" csv:"protein_daily"`
	Fat100g           float64 `json:"fat_100g" csv:"fat_100g"`
	FatDaily          float64 `json:"fat_daily" csv:"fat_daily"`
	Carbohydrate100g  float64 `json:"carbohydrate_100g" csv:"carbohydrate_100g"`
	CarbohydrateDaily float64 `json:"carbohydrate_daily" csv:"carbohydrate_daily"`
	Heat100g          float64 `json:"heat_100g" csv:"heat_100g"`
	HeatDaily         float64 `json:"heat_daily" csv:"heat_daily"`
	Calcium100g       float64 `json:"calcium_100g" csv:"calcium_100g"`
	CalciumDaily      float64 `json:"calcium_daily" csv:"calcium_daily"`
	Iron100g          float64 `json:"iron_100g" csv:"iron_100g"`
	IronDaily         float64 `json:"iron_daily" csv:"iron_daily"`
	Zinc100g          float64 `json:"zinc_100g" csv:"zinc_100g"`
	ZincDaily         float64 `json:"zinc_daily" csv:"zinc_daily"`
	VA100g            float64 `json:"va_100g" csv:"va_100g"`
	VADaily           float64 `json:"va_daily" csv:"va_daily"`
	VB1100g           float64 `json:"vb1_100g" csv:"vb1_100g"`
	VB1Daily          float64 `json:"vb1_daily" csv:"vb1_daily"`
	VB2100g           float64 `json:"vb2_100g" csv:"vb2_100g"`
	VB2Daily          float64 `json:"vb2_daily" csv:"vb2_daily"`
	VC100g            float64 `json:"vc_100g" csv:"vc_100g"`
	VCDaily           float64 `json:"vc_daily" csv:"vc_daily"`
	Category          string  `json:"category" csv:"category"`
}

// GetIngredients get ingredients
func GetIngredients(req *restful.Request, rsp *restful.Response) {
	ingredients := strings.Split(req.QueryParameter("ingredients"), ",")
	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetIngredients(ctx(req), &proto.GetIngredientRequest{
		Token:       idtoken,
		Ingredients: ingredients,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UpdateIngredient update ingredient
func UpdateIngredient(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *UpdateIngredientRequest
	err := decoder.Decode(&updatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidUpdateIngredientRequestBody)
		return
	}

	ingredients := []*proto.Ingredient{&proto.Ingredient{
		Ingredient:        updatereq.Ingredient,
		Alias:             updatereq.Alias,
		Protein_100G:      updatereq.Protein100g,
		ProteinDaily:      updatereq.ProteinDaily,
		Fat_100G:          updatereq.Fat100g,
		FatDaily:          updatereq.FatDaily,
		Carbohydrate_100G: updatereq.Carbohydrate100g,
		CarbohydrateDaily: updatereq.CarbohydrateDaily,
		Heat_100G:         updatereq.Heat100g,
		HeatDaily:         updatereq.HeatDaily,
		Calcium_100G:      updatereq.Calcium100g,
		CalciumDaily:      updatereq.CalciumDaily,
		Iron_100G:         updatereq.Iron100g,
		IronDaily:         updatereq.IronDaily,
		Zinc_100G:         updatereq.Zinc100g,
		ZincDaily:         updatereq.ZincDaily,
		Va_100G:           updatereq.VA100g,
		VaDaily:           updatereq.VADaily,
		Vb1_100G:          updatereq.VB1100g,
		Vb1Daily:          updatereq.VB1Daily,
		Vb2_100G:          updatereq.VB2100g,
		Vb2Daily:          updatereq.VB2Daily,
		Vc_100G:           updatereq.VC100g,
		VcDaily:           updatereq.VCDaily,
		Category:          updatereq.Category,
	}}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateIngredient(ctx(req), &proto.UpdateIngredientRequest{
		Token:       idtoken,
		Ingredients: ingredients,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

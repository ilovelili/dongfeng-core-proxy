package handlers

import (
	"encoding/json"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// IngredientNutrition ingredient nutrition
type IngredientNutrition struct {
	Protein100g       float64 `json:"protein_100g"`
	ProteinDaily      float64 `json:"protein_daily"`
	Fat100g           float64 `json:"fat_100g"`
	FatDaily          float64 `json:"fat_daily"`
	Carbohydrate100g  float64 `json:"carbohydrate_100g"`
	CarbohydrateDaily float64 `json:"carbohydrate_daily"`
	Heat100g          float64 `json:"heat_100g"`
	HeatDaily         float64 `json:"heat_daily"`
	Calcium100g       float64 `json:"calcium_100g"`
	CalciumDaily      float64 `json:"calcium_daily"`
	Iron100g          float64 `json:"iron_100g"`
	IronDaily         float64 `json:"iron_daily"`
	Zinc100g          float64 `json:"zinc_100g"`
	ZincDaily         float64 `json:"zinc_daily"`
	VA100g            float64 `json:"va_100g"`
	VADaily           float64 `json:"va_daily"`
	VB1100g           float64 `json:"vb1_100g"`
	VB1Daily          float64 `json:"vb1_daily"`
	VB2100g           float64 `json:"vb2_100g"`
	VB2Daily          float64 `json:"vb2_daily"`
	VC100g            float64 `json:"vc_100g"`
	VCDaily           float64 `json:"vc_daily"`
}

// UpdateIngredientRequest ingredient update request
type UpdateIngredientRequest struct {
	Ingredient          string `json:"ingredient"`
	IngredientNutrition `json:"nutrition"`
}

// GetIngredient get ingredient
func GetIngredient(req *restful.Request, rsp *restful.Response) {
	ingredientname := req.PathParameter("ingredient")
	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newnutritionclient().GetIngredient(ctx(req), &proto.GetIngredientRequest{
		Token:      idtoken,
		Ingredient: ingredientname,
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

	in := updatereq.IngredientNutrition
	ingredients := []*proto.Ingredient{&proto.Ingredient{
		Ingredient: updatereq.Ingredient,
		Nutrition: &proto.IngredientNutrition{
			Protein_100G:      in.Protein100g,
			ProteinDaily:      in.ProteinDaily,
			Fat_100G:          in.Fat100g,
			FatDaily:          in.FatDaily,
			Carbohydrate_100G: in.Carbohydrate100g,
			CarbohydrateDaily: in.CarbohydrateDaily,
			Heat_100G:         in.Heat100g,
			HeatDaily:         in.HeatDaily,
			Calcium_100G:      in.Calcium100g,
			CalciumDaily:      in.CalciumDaily,
			Iron_100G:         in.Iron100g,
			IronDaily:         in.IronDaily,
			Zinc_100G:         in.Zinc100g,
			ZincDaily:         in.ZincDaily,
			Va_100G:           in.VA100g,
			VaDaily:           in.VADaily,
			Vb1_100G:          in.VB1100g,
			Vb1Daily:          in.VB1Daily,
			Vb2_100G:          in.VB2100g,
			Vb2Daily:          in.VB2Daily,
			Vc_100G:           in.VC100g,
			VcDaily:           in.VCDaily,
		},
	}}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newnutritionclient().UpdateIngredient(ctx(req), &proto.UpdateIngredientRequest{
		Token:       idtoken,
		Ingredients: ingredients,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

package handlers

import (
	"encoding/json"
	"strings"

	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// IngredientRequestItem ingredient update request
type IngredientRequestItem struct {
	Ingredient        string  `json:"ingredient" csv:"原料"`
	Alias             string  `json:"alias" csv:"-"`
	Protein100g       float64 `json:"protein_100g" csv:"蛋白质(100g)"`
	ProteinDaily      float64 `json:"protein_daily" csv:"-"`
	Fat100g           float64 `json:"fat_100g" csv:"脂肪(100g)"`
	FatDaily          float64 `json:"fat_daily" csv:"-"`
	Carbohydrate100g  float64 `json:"carbohydrate_100g" csv:"碳水化合物(100g)"`
	CarbohydrateDaily float64 `json:"carbohydrate_daily" csv:"-"`
	Heat100g          float64 `json:"heat_100g" csv:"热量(100g)"`
	HeatDaily         float64 `json:"heat_daily" csv:"-"`
	Calcium100g       float64 `json:"calcium_100g" csv:"钙(100g)"`
	CalciumDaily      float64 `json:"calcium_daily" csv:"-"`
	Iron100g          float64 `json:"iron_100g" csv:"铁(100g)"`
	IronDaily         float64 `json:"iron_daily" csv:"-"`
	Zinc100g          float64 `json:"zinc_100g" csv:"锌(100g)"`
	ZincDaily         float64 `json:"zinc_daily" csv:"-"`
	VA100g            float64 `json:"va_100g" csv:"VA(100g)"`
	VADaily           float64 `json:"va_daily" csv:"-"`
	VB1100g           float64 `json:"vb1_100g" csv:"VB1(100g)"`
	VB1Daily          float64 `json:"vb1_daily" csv:"-"`
	VB2100g           float64 `json:"vb2_100g" csv:"VB2(100g)"`
	VB2Daily          float64 `json:"vb2_daily" csv:"-"`
	VC100g            float64 `json:"vc_100g" csv:"VC(100g)"`
	VCDaily           float64 `json:"vc_daily" csv:"-"`
	Category          string  `json:"category" csv:"类别"`
}

// GetIngredients get ingredients
func GetIngredients(req *restful.Request, rsp *restful.Response) {
	ingredients := strings.Split(req.QueryParameter("ingredients"), ",")
	_, pid, email, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetIngredients(ctx(req), &proto.GetIngredientRequest{
		Pid:         pid,
		Email:       email,
		Ingredients: ingredients,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// GetIngredientNames get ingredient names without JWT
func GetIngredientNames(req *restful.Request, rsp *restful.Response) {
	query := req.QueryParameter("q")
	response, err := newcoreclient().GetIngredientNames(ctx(req), &proto.GetIngredientNameRequest{Pattern: query})
	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	if len(response.GetNames()) == 0 {
		rsp.WriteAsJson([]string{})
		return
	}

	rsp.WriteAsJson(response.Names)
}

// UpdateIngredient update ingredient
func UpdateIngredient(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *IngredientRequestItem
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

	_, pid, email, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().UpdateIngredients(ctx(req), &proto.UpdateIngredientRequest{
		Pid:         pid,
		Email:       email,
		Ingredients: ingredients,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UpdateIngredients update ingredients
func UpdateIngredients(req *restful.Request, rsp *restful.Response) {
	file, _, err := req.Request.FormFile("file")
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidUpdateIngredientRequestBody)
		return
	}
	defer file.Close()

	ingredients := []*IngredientRequestItem{}
	if err := gocsv.Unmarshal(file, &ingredients); err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidUpdateIngredientRequestBody)
		return
	}

	_ingredients := []*proto.Ingredient{}
	for _, ingredient := range ingredients {
		_ingredients = append(_ingredients, &proto.Ingredient{
			Ingredient:        ingredient.Ingredient,
			Protein_100G:      ingredient.Protein100g,
			Fat_100G:          ingredient.Fat100g,
			Carbohydrate_100G: ingredient.Carbohydrate100g,
			Heat_100G:         ingredient.Heat100g,
			Calcium_100G:      ingredient.Calcium100g,
			Iron_100G:         ingredient.Iron100g,
			Zinc_100G:         ingredient.Zinc100g,
			Va_100G:           ingredient.VA100g,
			Vb1_100G:          ingredient.VB1100g,
			Vb2_100G:          ingredient.VB2100g,
			Vc_100G:           ingredient.VC100g,
			Category:          ingredient.Category,
		})
	}

	_, pid, email, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().UpdateIngredients(ctx(req), &proto.UpdateIngredientRequest{
		Pid:         pid,
		Email:       email,
		Ingredients: _ingredients,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

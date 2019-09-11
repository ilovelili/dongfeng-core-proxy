package handlers

import (
	"strings"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// IngredientUnitAmount ingredient unit amount
type IngredientUnitAmount struct {
	Ingredient string  `json:"ingredient"`
	UnitAmount float64 `json:"unit_amount,omitempty"`
}

// RecipeNutrition recipe nutrition
type RecipeNutrition struct {
	Carbohydrate float64 `json:"carbohydrate"`
	Dietaryfiber float64 `json:"dietaryfiber"`
	Protein      float64 `json:"protein"`
	Fat          float64 `json:"fat"`
	Heat         float64 `json:"heat"`
}

// UpdateRecipeRequest recipe update request
type UpdateRecipeRequest struct {
	Recipes []*UpdateRecipeItem `json:"recipes"`
}

// UpdateRecipeItem update recipe item
type UpdateRecipeItem struct {
	Recipe          string                  `json:"recipe"`
	Ingredients     []*IngredientUnitAmount `json:"ingredients"`
	RecipeNutrition `json:"nutrition"`
}

// GetRecipes get recipes
func GetRecipes(req *restful.Request, rsp *restful.Response) {
	names := strings.Split(req.QueryParameter("recipes"), ",")
	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetRecipes(ctx(req), &proto.GetRecipeRequest{
		Pid:   pid,
		Names: names,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

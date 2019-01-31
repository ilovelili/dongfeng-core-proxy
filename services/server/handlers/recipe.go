package handlers

import (
	"encoding/json"

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

// GetRecipe get recipes
func GetRecipe(req *restful.Request, rsp *restful.Response) {
	recipename := req.PathParameter("recipe")
	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newnutritionclient().GetRecipe(ctx(req), &proto.GetRecipeRequest{
		Token:  idtoken,
		Recipe: recipename,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UpdateRecipe update recipe
func UpdateRecipe(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *UpdateRecipeRequest
	err := decoder.Decode(&updatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidUpdateRecipeRequestBody)
		return
	}

	recipes := make([]*proto.Recipe, 0)
	for _, r := range updatereq.Recipes {
		rn := r.RecipeNutrition
		ingredientunitamounts := make([]*proto.IngredientUnitAmount, 0)
		for _, ua := range r.Ingredients {
			ingredientunitamounts = append(ingredientunitamounts, &proto.IngredientUnitAmount{
				Ingredient: ua.Ingredient,
				UnitAmount: ua.UnitAmount,
			})
		}

		recipes = append(recipes, &proto.Recipe{
			Recipe:      r.Recipe,
			Ingredients: ingredientunitamounts,
			Nutrition: &proto.RecipeNutrition{
				Carbohydrate: rn.Carbohydrate,
				Dietaryfiber: rn.Dietaryfiber,
				Protein:      rn.Protein,
				Fat:          rn.Fat,
				Heat:         rn.Heat,
			},
		})
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newnutritionclient().UpdateRecipe(ctx(req), &proto.UpdateRecipeRequest{
		Token:   idtoken,
		Recipes: recipes,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

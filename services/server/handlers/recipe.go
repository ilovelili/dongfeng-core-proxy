package handlers

import (
	excelize "github.com/360EntSecGroup-Skylar/excelize"
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	protobuf "github.com/ilovelili/dongfeng-protobuf"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
)

// UploadRecipe upload recipe list
func UploadRecipe(req *restful.Request, rsp *restful.Response) {
	if err := req.Request.ParseMultipartForm(32 << 20); err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToReadRecipeFile)
		return
	}

	file, _, err := req.Request.FormFile("recipe")
	defer file.Close()
	if err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToReadRecipeFile)
		return
	}

	excel, err := excelize.OpenReader(file)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyUnsupportedMimeType)
		return
	}

	recipeingredientmap := make(map[string][]string)
	for _, sheet := range excel.WorkBook.Sheets.Sheet {
		// get the "原料" sheet
		if sheet.Name != "原料" {
			continue
		}

		rows := excel.GetRows(sheet.Name)
		recipecolindex, ingredientcolindex := 0, 0
		for rindex, row := range rows {
			if rindex == 0 {
				for cindex, col := range row {
					if col == "菜品名称" {
						recipecolindex = cindex
					} else if col == "原料名称" {
						ingredientcolindex = cindex
					}
				}
			} else {
				if recipecolindex == 0 && ingredientcolindex == 0 {
					writeError(rsp, errorcode.CoreProxyBadFormatRecipeFile)
					return
				}

				recipe, ingredient := row[recipecolindex], row[ingredientcolindex]
				if v, ok := recipeingredientmap[recipe]; !ok {
					recipeingredientmap[recipe] = []string{ingredient}
				} else {
					if !sharedlib.ContainString(v, ingredient) {
						recipeingredientmap[recipe] = append(v, ingredient)
					}
				}
			}
		}
	}

	recipes := make([]*protobuf.Recipe, 0)
	for k, v := range recipeingredientmap {
		recipes = append(recipes, &protobuf.Recipe{
			Recipe:      k,
			Ingredients: v,
		})
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newnutritionclient().UpdateRecipe(ctx(req), &protobuf.UpdateRecipeRequest{
		Token:   idtoken,
		Recipes: recipes,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

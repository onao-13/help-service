package api

import (
	"github.com/gorilla/mux"
	"help-service/internals/app/controller"
)

const helpApiPath = "/help/"

func CreateRoute(helpController *controller.HelpController) *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc(helpApiPath, helpController.ShowMainPageDoc)
	route.HandleFunc(helpApiPath+"home-page", helpController.ShowHomePageDoc)
	route.HandleFunc(helpApiPath+"articles", helpController.ShowArticlesPageDoc)
	route.HandleFunc(helpApiPath+"categories", helpController.ShowCategoriesPageDoc)
	route.HandleFunc(helpApiPath+"products", helpController.ShowProductsPageDoc)
	return route
}

//helpApiPath = "/help/"

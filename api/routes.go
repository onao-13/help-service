package api

import (
	"help-service/internals/app/controller"

	"github.com/gorilla/mux"
)

func CreateRoute(helpController *controller.HelpController) *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", helpController.ShowMainPageDoc)
	route.HandleFunc("/home-page", helpController.ShowHomePageDoc)
	route.HandleFunc("/articles", helpController.ShowArticlesPageDoc)
	route.HandleFunc("/categories", helpController.ShowCategoriesPageDoc)
	route.HandleFunc("/products", helpController.ShowProductsPageDoc)
	return route
}

//helpApiPath = "/help/"

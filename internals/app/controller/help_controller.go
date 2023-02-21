package controller

import (
	"help-service/internals/app/handler"
	"help-service/internals/utils/md"
	"net/http"
)

// HelpController Контроллер для отправки документации по работе с API
type HelpController struct {
	md *md.MdWorkerImpl
}

// New Создание нового контроллера
func New() *HelpController {
	mdWorker := md.New()
	return &HelpController{md: mdWorker}
}

// ShowHomePageDoc Отправка документации с API главной страницы
func (controller *HelpController) ShowHomePageDoc(w http.ResponseWriter, r *http.Request) {
	html := controller.md.GetPage("doc/page/doc.md", "Документация главной страницы")
	handler.WrapHtmlPage(w, r, html)
}

func (controller *HelpController) ShowMainPageDoc(w http.ResponseWriter, r *http.Request) {
	html := controller.md.GetPage("doc/main/doc.md", "Главная")
	handler.WrapHtmlPage(w, r, html)
}

func (controller *HelpController) ShowCategoriesPageDoc(w http.ResponseWriter, r *http.Request) {
	html := controller.md.GetPage("doc/categories/doc.md", "Документация категорий")
	handler.WrapHtmlPage(w, r, html)
}

func (controller *HelpController) ShowArticlesPageDoc(w http.ResponseWriter, r *http.Request) {
	html := controller.md.GetPage("doc/articles/doc.md", "Документация статей")
	handler.WrapHtmlPage(w, r, html)
}

func (controller *HelpController) ShowProductsPageDoc(w http.ResponseWriter, r *http.Request) {
	html := controller.md.GetPage("doc/products/doc.md", "Документация продуктов")
	handler.WrapHtmlPage(w, r, html)
}

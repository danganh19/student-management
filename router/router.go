package router

import (
	"github.com/go-chi/chi"
	"github.com/swaggo/http-swagger"
	"ims-be/controller"
	_ "ims-be/docs"
	"net/http"
)

const basePath string = "/api/tda/student-management"

func Router() http.Handler {
	r := chi.NewRouter()
	productController := controller.NewProductController()
	inventoryController := controller.NewInventoryController()
	importController := controller.NewImportController()
	exportController := controller.NewExportController()
	r.Get(basePath+"/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/api/tda/student-management/swagger/doc.json"),
	))

	r.Route(basePath, func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/export", exportController.CreateExportRecord)

			r.Get("/export/{Id}", exportController.GetExportRecordById)
			r.Get("/exports", exportController.GetExportRecords)

			r.Put("/export/deactive/{Id}", exportController.DeActiveExportRecord)
			r.Put("/export/active/{Id}", exportController.ActiveExportRecord)
			r.Put("/export/update/{Id}", exportController.UpdateExportRecord)
		})

		r.Group(func(r chi.Router) {
			r.Post("/import", importController.CreateImportRecord)

			r.Get("/import/{Id}", importController.GetImportRecordById)
			r.Get("/imports", importController.GetImportRecords)

			r.Put("/import/deactive/{Id}", importController.DeActiveImportRecord)
			r.Put("/import/active/{Id}", importController.ActiveImportRecord)
			r.Put("/import/update/{Id}", importController.UpdateImportRecord)
		})

		r.Group(func(r chi.Router) {
			r.Get("/inventory/{Id}", inventoryController.GetInventoryById)
			r.Get("/inventories", inventoryController.GetInventories)
		})

		r.Group(func(r chi.Router) {
			r.Post("/product", productController.CreateProduct)

			r.Get("/product/{Id}", productController.GetProductById)
			r.Get("/products", productController.GetProducts)

			r.Put("/product/deactive/{Id}", productController.DeActiveProduct)
			r.Put("/product/active/{Id}", productController.ActiveProduct)
			r.Put("/product/update/{Id}", productController.UpdateProduct)
		})
	})
	return r
}

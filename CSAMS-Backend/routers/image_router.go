package routers

import (
	"CSAMS-Backend/api"
	"CSAMS-Backend/middleware"
)

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImageApi
	router.POST("image", middleware.JwtAuth(), app.ImageUploadDataView)
}

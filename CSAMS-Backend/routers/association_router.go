package routers

import (
	"CSAMS-Backend/api"
	"CSAMS-Backend/middleware"
)

func (router RouterGroup) AssociationRouter() {
	app := api.ApiGroupApp.AssociationApi
	router.GET("association_info", middleware.JwtAuth(), app.AssociationInfoView)
}

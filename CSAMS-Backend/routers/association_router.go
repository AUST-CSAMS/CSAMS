package routers

import (
	"CSAMS-Backend/api"
	"CSAMS-Backend/middleware"
)

func (router RouterGroup) AssociationRouter() {
	app := api.ApiGroupApp.AssociationApi
	router.GET("association_info", middleware.JwtAuth(), app.AssociationInfoView)
	router.POST("associations", middleware.JwtTeacher(), app.AssociationCreateView)
	router.GET("associations", app.AssociationListView)
	router.GET("associations/member", middleware.JwtAuth(), app.AssociationMemberListView)
	router.PUT("associations/member", middleware.JwtTeacher(), app.AssociationManageView)
	router.POST("associations/member", middleware.JwtStudentAdmin(), app.AssociationJoinView)
	router.DELETE("associations/member", middleware.JwtStudentAdmin(), app.AssociationQuitView)
}

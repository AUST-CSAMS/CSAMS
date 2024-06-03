package routers

import (
	"CSAMS-Backend/api"
	"CSAMS-Backend/middleware"
)

func (router RouterGroup) AssignmentRouter() {
	app := api.ApiGroupApp.AssignmentApi
	router.POST("assignment_create", middleware.JwtTeacher(), app.AssignmentCreateView)
	router.GET("assignment_list", middleware.JwtAuth(), app.AssignmentListView)
	router.POST("assignment_score", middleware.JwtTeacher(), app.AssignmentScoreView)
}

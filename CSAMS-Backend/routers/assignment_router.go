package routers

import (
	"CSAMS-Backend/api"
	"CSAMS-Backend/middleware"
)

func (router RouterGroup) AssignmentRouter() {
	app := api.ApiGroupApp.AssignmentApi
	router.GET("assignments", middleware.JwtTeacher(), app.AssignmentListView)
	router.PUT("assignments/submit", middleware.JwtAuth(), app.AssignmentSubmitView)
	router.PUT("assignments/correct", middleware.JwtTeacher(), app.AssignmentCorrectView)
}

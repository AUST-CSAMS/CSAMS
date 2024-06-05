package routers

import (
	"CSAMS-Backend/api"
	"CSAMS-Backend/middleware"
)

func (router RouterGroup) AssignmentRouter() {
	app := api.ApiGroupApp.AssignmentApi
	router.GET("assignments", middleware.JwtTeacher(), app.AssignmentListView)
}

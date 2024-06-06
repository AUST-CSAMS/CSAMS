package routers

import (
	"CSAMS-Backend/api"
	"CSAMS-Backend/middleware"
)

func (router RouterGroup) ActivityRouter() {
	app := api.ApiGroupApp.ActivityApi
	router.GET("activities/finished", app.ActivityFinishedListView)
	router.GET("activities/ongoing", app.ActivityOngoingListView)
	router.GET("activities/upcoming", app.ActivityUpcomingListView)
	router.POST("activities", middleware.JwtTeacher(), app.ActivityCreateView)
	router.GET("activities", middleware.JwtAuth(), app.ActivityListView)
	router.GET("activities_create", middleware.JwtTeacher(), app.ActivityCreatedListView)
	router.GET("activity_info/:id", app.ActivityInfoView)
	router.POST("activities/join", middleware.JwtAuth(), app.ActivityJoinView)
	router.PUT("activities/end", middleware.JwtTeacher(), app.ActivityEndView)
}

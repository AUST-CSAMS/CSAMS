package routers

import "CSAMS-Backend/api"

func (router RouterGroup) ActivityRouter() {
	app := api.ApiGroupApp.ActivityApi
	router.GET("activities/finished", app.ActivityFinishedListView)
	router.GET("activities/ongoing", app.ActivityOngoingListView)
	router.GET("activities/upcoming", app.ActivityUpcomingListView)
}

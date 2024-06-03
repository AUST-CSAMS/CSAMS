package api

import (
	"CSAMS-Backend/api/activity_api"
	"CSAMS-Backend/api/assignment_api"
	"CSAMS-Backend/api/association_api"
	"CSAMS-Backend/api/image_api"
	"CSAMS-Backend/api/user_api"
)

type ApiGroup struct {
	UserApi        user_api.UserApi
	AssociationApi association_api.AssociationApi
	ActivityApi    activity_api.ActivityApi
	AssignmentApi  assignment_api.AssignmentApi
	ImageApi       image_api.ImageApi
}

var ApiGroupApp = new(ApiGroup)

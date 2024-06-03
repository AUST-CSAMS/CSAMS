package service

import "CSAMS-Backend/service/user_ser"

type ServiceGroup struct {
	UserService user_ser.UserService
}

var ServiceApp = new(ServiceGroup)

package user_api

import (
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/service"
	"CSAMS-Backend/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	token := c.Request.Header.Get("token")
	err := service.ServiceApp.UserService.Logout(claims, token)
	log.Print(fmt.Sprintf("用户 %s 注销登录", claims.Name))
	if err != nil {
		log.Print(err)
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OkWithMessage("注销成功", c)
}

package user_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"github.com/gin-gonic/gin"
	"log"
)

type UpdateAvatarRequest struct {
	Avatar string `json:"avatar" binding:"required"`
}

func (UserApi) UserUpdateAvatar(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr UpdateAvatarRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	err = global.DB.Model(&user).Update("avatar", cr.Avatar).Error
	if err != nil {
		log.Print(err)
		res.FailWithMessage("头像修改失败", c)
		return
	}
	res.OkWithMessage("头像修改成功", c)
	return
}

package user_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"CSAMS-Backend/utils/pwd"
	"github.com/gin-gonic/gin"
	"log"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"` // 旧密码
	Pwd    string `json:"pwd" binding:"required" msg:"请输入新密码"`     // 新密码
}

func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr UpdatePasswordRequest
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
	// 判断密码是否一致
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithMessage("密码错误", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Pwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		log.Print(err)
		res.FailWithMessage("密码修改失败", c)
		return
	}
	res.OkWithMessage("密码修改成功", c)
	return
}

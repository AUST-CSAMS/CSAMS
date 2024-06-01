package user_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils"
	"CSAMS-Backend/utils/jwts"
	"CSAMS-Backend/utils/pwd"
	"github.com/gin-gonic/gin"
	"log"
)

type UserLoginRequest struct {
	ID       uint64 `json:"id" binding:"required" msg:"请输入学工号"`     // 学工号
	Password string `json:"password" binding:"required" msg:"请输入密码"` // 密码
}

func (UserApi) UserLoginView(c *gin.Context) {
	var cr UserLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "id = ?", cr.ID).Error
	if err != nil {
		// 对内
		log.Print("用户名不存在")
		//对外
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		log.Print("用户名密码错误")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		ID:   userModel.ID,
		Name: userModel.Name,
		Role: int(userModel.Role),
	})
	if err != nil {
		log.Print(err)
		res.FailWithMessage("token生成失败", c)
		return
	}
	//获取 ip  地址
	ip, addr := utils.GetAddrByGin(c)
	//在请求头中设置token
	c.Request.Header.Set("token", token)

	global.DB.Create(&models.LoginDataModel{
		UserID: userModel.ID,
		IP:     ip,
		Name:   userModel.Name,
		Token:  token,
		Device: "",
		Addr:   addr,
	})

	res.OkWithData(token, c)
}

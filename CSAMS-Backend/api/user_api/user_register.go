package user_api

import (
	"CSAMS-Backend/models/ctype"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/service/user_ser"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type UserRegisterRequest struct {
	ID       uint64      `json:"id" binding:"required" msg:"请输入学工号"`      // 学工号
	Password string      `json:"password" binding:"required" msg:"请输入密码"` // 密码
	Name     string      `json:"name" binding:"required" msg:"请输入姓名"`     // 姓名
	Age      int         `json:"age" binding:"required" msg:"请输入年龄"`      // 年龄
	Gender   string      `json:"gender" binding:"required" msg:"请选择性别"`   // 性别
	Role     ctype.Role  `json:"role" binding:"required" msg:"请选择身份"`     // 权限 1 教师 2 学生
	Major    ctype.Major `json:"major" binding:"required" msg:"请输入专业"`    // 专业
	Tel      uint64      `json:"tel" binding:"required" msg:"请输入电话号码"`    // 电话号码
}

func (UserApi) UserRegisterView(c *gin.Context) {
	var cr UserRegisterRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	err := user_ser.UserService{}.Register(cr.ID, cr.Password, cr.Name, cr.Age, cr.Gender, cr.Role, cr.Major, cr.Tel)
	if err != nil {
		log.Print(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("用户%s创建成功!", cr.Name), c)
	return
}

package association_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type AssociationJoinRequest struct {
	ID uint64 `json:"id" binding:"required" msg:"请输入学号"` // 学号
}

func (AssociationApi) AssociationJoinView(c *gin.Context) {
	var cr AssociationJoinRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var AssociationModel models.AssociationModel

	err := global.DB.Take(&AssociationModel, "teacher_id = ?", claims.UserID).Error
	if err != nil {
		log.Print(err)
		res.FailWithMessage("还未创建协会", c)
		return
	}
	// 更新协会成员表
	err = global.DB.Create(&models.AssociationMemberModel{
		ID:          AssociationModel.ID,             // 写入协会id，这里表上面没关联，但是后端代码关联了
		Users:       []models.UserModel{{ID: cr.ID}}, // 将此学生加入到关联表中
		Posts:       "成员",                            // 设置职位信息
		JoiningTime: time.Now(),                      // 设置加入时间
	}).Error
	if err != nil {
		res.FailWithMessage("添加关联成员失败", c)
		return
	}

	// 成功响应
	res.OkWithMessage("添加关联成员成功", c)
}

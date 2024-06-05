package assignment_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"github.com/gin-gonic/gin"
	"log"
)

type AssignmentSubmitRequest struct {
	ID      uint64 `json:"id" binding:"required"` // 活动id
	Content string `json:"content"  binding:"required" msg:"请输入作业内容"`
}

func (AssignmentApi) AssignmentSubmitView(c *gin.Context) {
	var cr AssignmentSubmitRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var AssignmentModel models.AssignmentModel
	err := global.DB.Where("activity_id = ? AND user_id = ?", cr.ID, claims.UserID).Find(&AssignmentModel).Error
	if err != nil {
		log.Print(err)
		res.FailWithMessage("作业不存在", c)
		return
	}

	err = global.DB.Model(&AssignmentModel).Update("content", cr.Content).Error
	if err != nil {
		log.Print(err)
		res.FailWithMessage("作业提交失败", c)
		return
	}
	res.OkWithMessage("作业提交成功", c)
}

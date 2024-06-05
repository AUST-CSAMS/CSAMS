package activity_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ActivityEndRequest struct {
	ID uint64 `json:"id" bind:"required"` // 活动id
}

func (ActivityApi) ActivityEndView(c *gin.Context) {
	var cr ActivityEndRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var activity models.ActivityModel
	err := global.DB.Take(&activity, cr.ID).Error
	if err != nil {
		res.FailWithMessage("活动不存在", c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	if activity.ResponsiblePerson != claims.Name {
		res.FailWithMessage("没有停止该活动的权限", c)
		return
	}

	// 未提交作业，扣诚信度
	var assignmentsNotSubmit []models.AssignmentModel
	var userIDsNotSubmit []uint64
	// 查找所有IsSubmit为false且ActivityID匹配的记录
	global.DB.Where("activity_id = ? AND is_submit = ?", cr.ID, false).Find(&assignmentsNotSubmit)
	// 将对应的UserID存放到数组中
	for _, assignment := range assignmentsNotSubmit {
		userIDsNotSubmit = append(userIDsNotSubmit, assignment.UserID)
	}
	// 更新对应的UserModel中的Integrity字段的值
	if len(userIDsNotSubmit) > 0 {
		global.DB.Model(&models.UserModel{}).Where("id IN (?)", userIDsNotSubmit).Update("integrity", gorm.Expr("integrity - ?", 5))
	}

	// 完成作业，给活动积分
	var assignments []models.AssignmentModel
	var userIDs []uint64
	// 查找所有IsFinish为true且ActivityID匹配的记录
	global.DB.Where("activity_id = ? AND is_finish = ?", cr.ID, true).Find(&assignments)
	// 将对应的UserID存放到数组中
	for _, assignment := range assignments {
		userIDs = append(userIDs, assignment.UserID)
	}
	// 更新对应的UserModel中的Score字段的值
	if len(userIDs) > 0 {
		global.DB.Model(&models.UserModel{}).Where("id IN (?)", userIDs).Update("score", gorm.Expr("score + ?", activity.Score))

	}

	res.OkWithMessage("作业截止完成", c)
}

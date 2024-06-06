package activity_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/ctype"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"github.com/gin-gonic/gin"
	"time"
)

type ActivityJoinModelRequest struct {
	ID uint64 `json:"id" bind:"required"` // 活动id
}

func (ActivityApi) ActivityJoinView(c *gin.Context) {
	var cr ActivityJoinModelRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var activity models.ActivityModel
	err := global.DB.Take(&activity, cr.ID).Error
	if err != nil {
		print(err)
		res.FailWithMessage("活动不存在", c)
		return
	}

	if activity.EndTime.Before(time.Now()) {
		res.FailWithMessage("活动已结束", c)
		return
	}

	if activity.StartTime.After(time.Now()) {
		res.FailWithMessage("活动未开始", c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("该用户不存在", c)
		return
	}

	// 限制筛选
	if !contains(activity.Limit, user.Major) {
		res.FailWithMessage("没有报名资格", c)
		return
	}

	// 是否已经报名
	err = global.DB.Take(&models.ActivityLogModel{}, "activity_id = ? AND user_id = ?", cr.ID, claims.UserID).Error
	if err == nil {
		res.FailWithMessage("用户已经报名", c)
		return
	}

	// 使用ORM保存ActivityLogModel实例并进行关联
	err = global.DB.Create(&models.ActivityLogModel{
		ActivityID: cr.ID,
		UserID:     claims.UserID,
	}).Error
	if err != nil {
		res.FailWithMessage("添加活动记录失败", c)
		return
	}
	res.OkWithMessage("添加活动记录成功", c)

	// 使用ORM保存AssignmentModel实例并进行关联
	err = global.DB.Create(&models.AssignmentModel{
		ActivityID: cr.ID,
		UserID:     claims.UserID,
		IsSubmit:   false,
		IsFinish:   false,
		IsCorrect:  false,
	}).Error
	if err != nil {
		res.FailWithMessage("添加作业失败", c)
		return
	}
	res.OkWithMessage("添加作业成功", c)
	return
}

func contains(arr ctype.MajorArray, target ctype.Major) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

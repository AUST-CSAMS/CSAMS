package assignment_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

type AssignmentCreateRequest struct {
	UserID     uint64 `json:"user_id" binding:"required" msg:"请输入用户id"`     // 用户id
	ActivityID uint64 `json:"activity_id" binding:"required" msg:"请输入活动id"` // 活动id
	Content    string `json:"content" binding:"required" msg:"请输入作业内容"`     // 作业内容
}

func (AssignmentApi) AssignmentCreateView(c *gin.Context) {
	var acr AssignmentCreateRequest
	if err := c.ShouldBindJSON(&acr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	err := createAssignment(acr.UserID, acr.ActivityID, acr.Content)
	if err != nil {
		log.Print(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithMessage("作业上传成功!", c)
	return
}

func createAssignment(userID, activityID uint64, content string) error {
	var UserModel models.UserModel
	err := global.DB.Take(&UserModel, "id = ?", userID).Error
	if err != nil {
		return errors.New("用户不存在")
	}

	var ActivityModel models.ActivityModel
	err = global.DB.Take(&ActivityModel, "id = ?", activityID).Error
	if err != nil {
		return errors.New("活动不存在")
	}

	err = global.DB.Create(&models.AssignmentModel{
		UserID:     userID,
		ActivityID: activityID,
		Content:    content,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

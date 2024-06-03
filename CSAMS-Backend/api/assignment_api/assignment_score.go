package assignment_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"github.com/gin-gonic/gin"
)

func (AssignmentApi) AssignmentScoreView(c *gin.Context) {
	var assignment models.AssignmentModel
	var activity models.ActivityModel
	if err := c.ShouldBindJSON(&assignment); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 老师已经验证权限，并从请求中获取当前登录老师的ID
	userID, _ := c.Get("user_id")
	teacherID := userID.(uint64)

	// 查询数据库获取作业信息
	err := global.DB.First(&assignment, assignment.ID).Error
	if err != nil {
		res.FailWithMessage("未查询到作业信息", c)
		return
	}

	// 查询老师信息，信息存储在User表中
	var teacher models.UserModel
	if err := global.DB.First(&teacher, teacherID).Error; err != nil {
		res.FailWithMessage("未查询到老师信息", c)
		return
	}

	// 检查作业是否已经评分过
	if teacher.Score > 0 {
		res.FailWithMessage("作业已经评分过了", c)
		return
	}

	// 假设前端传入了评分值，存储在activity.Score字段中(?)

	// 更新数据库中作业信息，存储评分结果
	teacher.Score = activity.Score
	if err := global.DB.Save(&assignment).Error; err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithData("作业评分完成", c)
}

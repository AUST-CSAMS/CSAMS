package assignment_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type AssignmentCorrectRequest struct {
	ID       uint64 `json:"id" bind:"required"`        // 作业id
	IsFinish bool   `json:"is_finish" bind:"required"` // 是否完成
}

func (AssignmentApi) AssignmentCorrectView(c *gin.Context) {
	var cr AssignmentCorrectRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	maps := structs.Map(&struct {
		IsFinish  bool
		IsCorrect bool
	}{
		IsFinish:  cr.IsFinish,
		IsCorrect: true,
	})

	err := global.DB.Take(&models.AssignmentModel{}, cr.ID).Updates(maps).Error
	if err != nil {
		res.FailWithMessage("作业不存在", c)
		return
	}
	res.OkWithMessage("作业批改成功", c)

}

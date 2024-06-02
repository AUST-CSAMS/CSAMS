package activity_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type ActivityCreateRequest struct {
	ID           uint64 `json:"id" binding:"required" msg:"请输入活动id"`              // 活动id
	ActivityName string `json:"activity_name" binding:"required" msg:"请输入活动名称"` // 活动名称
	StartTime    string `json:"startTime" binding:"required" msg:"请输入开始时间"`     // 开始时间
	EndTime      string `json:"endTime" binding:"required" msg:"请输入结束时间"`       // 结束时间
	Location     string `json:"location" binding:"required" msg:"请输入活动地点"`      // 活动地点
	Introduction string `json:"introduction" binding:"required" msg:"请输入活动简介"`  // 活动简介
	BannerID     uint64 `json:"banner_id" binding:"required"`
	Score        uint64 `json:"score" binding:"required" msg:"请输入活动积分"` // 活动积分
}

func (ActivityApi) ActivityCreateView(c *gin.Context) {
	var cr ActivityCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	err := createActivity(cr.ID, cr.ActivityName, cr.StartTime, cr.EndTime, cr.Location, cr.Introduction, claims.Name, cr.Score)
	if err != nil {
		log.Print(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("活动%s创建成功!", cr.ActivityName), c)
	return
}

func createActivity(id uint64, activityName, startTime, endTime, location, introduction, responsiblePerson string, score uint64) error {
	var ActivityModel models.ActivityModel

	err := global.DB.Take(&ActivityModel, "id = ?", id).Error
	if err == nil {
		return errors.New("活动id已存在")
	}

	_startTime, dateTimeErr := time.Parse("2006-01-02", startTime)
	if dateTimeErr != nil {
		return errors.New("时间格式错误")
	}

	_endTime, dateTimeErr := time.Parse("2006-01-02", endTime)
	if dateTimeErr != nil {
		return errors.New("时间格式错误")
	}

	// 入库
	err = global.DB.Create(&models.ActivityModel{
		ID:           id,
		ActivityName: activityName,
		StartTime:    _startTime,
		EndTime:      _endTime,
		Location:     location,
		Introduction: introduction,
		//	Image:             image,
		ResponsiblePerson: responsiblePerson,
		//	Tel:               tel,
		Score: score,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

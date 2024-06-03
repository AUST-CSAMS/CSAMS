package activity_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/ctype"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type ActivityCreateRequest struct {
	ID           uint64           `json:"id" binding:"required" msg:"请输入活动id"`              // 活动id
	ActivityName string           `json:"activity_name" binding:"required" msg:"请输入活动名称"` // 活动名称
	StartTime    string           `json:"startTime" binding:"required" msg:"请输入开始时间"`     // 开始时间
	EndTime      string           `json:"endTime" binding:"required" msg:"请输入结束时间"`       // 结束时间
	Location     string           `json:"location" binding:"required" msg:"请输入活动地点"`      // 活动地点
	Introduction string           `json:"introduction" binding:"required" msg:"请输入活动简介"`  // 活动简介
	Image        string           `json:"image" binding:"required"`                              //活动图片路径
	Score        uint64           `json:"score" binding:"required" msg:"请输入活动积分"`         // 活动积分
	Limit        ctype.MajorArray `json:"limit"`                                                 // 专业限制
}

func (ActivityApi) ActivityCreateView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr ActivityCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var ActivityModel models.ActivityModel

	err := global.DB.Take(&ActivityModel, "id = ?", cr.ID).Error
	if err == nil {
		res.FailWithMessage("活动id已存在", c)
		return
	}

	startTime, dateTimeErr := time.Parse("2006-01-02", cr.StartTime)
	if dateTimeErr != nil {
		res.FailWithMessage("时间格式错误", c)
		return
	}

	endTime, dateTimeErr := time.Parse("2006-01-02", cr.EndTime)
	if dateTimeErr != nil {
		res.FailWithMessage("时间格式错误", c)
		return
	}

	var userInfo models.UserModel
	err = global.DB.Take(&userInfo, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	// 入库
	err = global.DB.Create(&models.ActivityModel{
		ID:                cr.ID,
		ActivityName:      cr.ActivityName,
		StartTime:         startTime,
		EndTime:           endTime,
		Location:          cr.Location,
		Introduction:      cr.Introduction,
		Image:             cr.Image,
		ResponsiblePerson: claims.Name,
		Tel:               userInfo.Tel,
		Score:             cr.Score,
		Limit:             cr.Limit,
	}).Error

	if err != nil {
		log.Print(err)
		res.FailWithMessage(fmt.Sprintf("活动%s创建失败!", cr.ActivityName), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("活动%s创建成功!", cr.ActivityName), c)

	return
}

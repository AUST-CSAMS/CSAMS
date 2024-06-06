package activity_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"github.com/gin-gonic/gin"
)

type ActivityInfoRequest struct {
	ID uint64 `json:"id" form:"id" uri:"id"` // 活动id
}

func (ActivityApi) ActivityInfoView(c *gin.Context) {
	var cr ActivityInfoRequest
	if err := c.ShouldBindUri(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var activityInfo models.ActivityModel
	err := global.DB.Take(&activityInfo, cr.ID).Error
	if err != nil {
		res.FailWithMessage("活动不存在", c)
		return
	}
	res.OkWithData(activityInfo, c)
}

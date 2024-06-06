package activity_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/service/common"
	"CSAMS-Backend/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (ActivityApi) ActivityListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 查询 activityLogs 表获取含有相同 UserID 的 ActivityID 列表
	var activityLogs []models.ActivityLogModel
	err := global.DB.Where("user_id = ?", claims.UserID).Find(&activityLogs).Error
	if err != nil {
		res.FailWithMessage("用户没有参加活动", c)
		return
	}
	var activityIDs []uint64
	for _, aa := range activityLogs {
		activityIDs = append(activityIDs, aa.ActivityID)
	}

	// ActivityModel 表获取所有含有相同 ActivityID 的记录
	query := global.DB.Where("id IN (?)", activityIDs)

	list, count, _ := common.ComList(models.ActivityModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
		Where:    query,
	})

	res.OkWithList(list, count, c)
}

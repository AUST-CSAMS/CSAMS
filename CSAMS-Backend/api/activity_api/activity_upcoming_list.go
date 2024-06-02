package activity_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/service/common"
	"github.com/gin-gonic/gin"
	"time"
)

func (ActivityApi) ActivityUpcomingListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	date := time.Now().Format("2006-01-02")

	query := global.DB.Where("start_time > ?", date)

	list, count, _ := common.ComList(models.ActivityModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
		Where:    query,
	})

	res.OkWithList(list, count, c)
}

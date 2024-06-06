package activity_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/service/common"
	"CSAMS-Backend/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (ActivityApi) ActivityCreatedListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var association models.AssociationModel
	err := global.DB.Take(&association, "teacher_id", claims.UserID).Error
	if err != nil {
		res.FailWithMessage("未创建协会", c)
		return
	}

	var activityAssociations []models.ActivityAssociationModel
	err = global.DB.Where("association_id = ?", association.ID).Find(&activityAssociations).Error
	if err != nil {
		res.FailWithMessage("没有活动记录", c)
		return
	}

	var activityIDs []uint64
	for _, aa := range activityAssociations {
		activityIDs = append(activityIDs, aa.ActivityID)
	}

	query := global.DB.Where("activity_id IN (?)", activityIDs)

	list, count, _ := common.ComList(models.ActivityModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
		Where:    query,
	})

	res.OkWithList(list, count, c)
}

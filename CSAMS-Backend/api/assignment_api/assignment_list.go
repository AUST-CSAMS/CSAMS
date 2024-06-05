package assignment_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/service/common"
	"CSAMS-Backend/utils/jwts"
	"github.com/gin-gonic/gin"
)

// AssignmentListView 负责教师批改作业的列表
func (AssignmentApi) AssignmentListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 查询 AssociationModel 表获取 AssociationID 列表
	var associations []models.AssociationModel
	err := global.DB.Where("teacher_id = ?", claims.UserID).Find(&associations).Error
	if err != nil {
		res.FailWithMessage("教师没有创建协会", c)
		return
	}

	// 查询 ActivityAssociationModel 表获取含有相同 AssociationID 的 ActivityID 列表
	var activityIDs []uint64
	for _, association := range associations {
		var activityAssociations []models.ActivityAssociationModel
		err := global.DB.Where("association_id = ?", association.ID).Find(&activityAssociations).Error
		if err != nil {
			res.FailWithMessage("协会没有发布活动", c)
			return
		}
		for _, aa := range activityAssociations {
			activityIDs = append(activityIDs, aa.ActivityID)
		}
	}

	// AssignmentModel 表获取所有含有相同 ActivityID 的记录
	query := global.DB.Where("activity_id IN (?)", activityIDs)

	list, count, _ := common.ComList(models.AssignmentModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
		Where:    query,
	})

	res.OkWithList(list, count, c)
}

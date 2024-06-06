package association_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/service/common"
	"CSAMS-Backend/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (AssociationApi) AssociationMemberListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var associationMember models.AssociationMemberModel
	err := global.DB.Take(&associationMember, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户未加入协会", c)
		return
	}
	query := global.DB.Where("association_id = ?", associationMember.AssociationID)
	list, count, _ := common.ComList(models.AssociationMemberModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
		Where:    query,
	})

	res.OkWithList(list, count, c)
}

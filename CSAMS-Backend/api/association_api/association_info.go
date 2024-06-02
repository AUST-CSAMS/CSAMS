package association_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (AssociationApi) AssociationInfoView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 查询用户所属协会ID
	var userInfo models.UserModel
	err := global.DB.Take(&userInfo, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	// 如果用户未加入任何协会，返回提示信息
	if userInfo.AssociationID == nil {
		res.OkWithMessage("用户未加入协会", c)
		return
	}

	// 查询用户所属协会信息
	var associationInfo models.AssociationModel
	err = global.DB.Take(&associationInfo, *userInfo.AssociationID).Error
	if err != nil {
		res.FailWithMessage("协会不存在", c)
		return
	}

	res.OkWithData(associationInfo, c)
}

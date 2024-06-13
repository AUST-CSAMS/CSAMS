package association_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type AssociationUpdateRequest struct {
	AssociationName string `json:"association_name" binding:"required" msg:"请输入协会名称"` // 协会名称
	Introduction    string `json:"introduction" binding:"required" msg:"请输入协会简介"`     // 协会简介
}

func (AssociationApi) AssociationUpdateView(c *gin.Context) {
	var cr AssociationUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var AssociationModel models.AssociationModel

	err := global.DB.Take(&AssociationModel, "teacher_id = ?", claims.UserID).Error
	if err != nil {
		res.FailWithMessage("未创建协会", c)
		return
	}

	maps := structs.Map(&struct {
		AssociationName string
		Introduction    string
	}{
		AssociationName: cr.AssociationName,
		Introduction:    cr.Introduction,
	})

	err = global.DB.Model(&AssociationModel).Updates(maps).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("协会信息修改成功", c)
	return

}

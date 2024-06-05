package association_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"github.com/gin-gonic/gin"
	"log"
)

type AssociationQuitRequest struct {
	ID uint64 `json:"id" binding:"required"` // 协会id
}

func (AssociationApi) AssociationQuitView(c *gin.Context) {
	var cr AssociationQuitRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var AssociationModel models.AssociationModel

	err := global.DB.Take(&AssociationModel, "teacher_id = ?", claims.UserID).Error
	if err != nil {
		log.Print(err)
		res.FailWithMessage("还未创建协会", c)
		return
	}

	var associationMember models.AssociationMemberModel

	err = global.DB.Take(&associationMember, cr.ID).Delete(&associationMember).Error
	if err != nil {
		print(err)
		res.FailWithMessage("成员删除失败", c)
	}
	res.OkWithMessage("成员删除成功", c)
}

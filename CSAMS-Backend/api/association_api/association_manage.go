package association_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"github.com/gin-gonic/gin"
	"log"
)

type AssociationManageRequest struct {
	ID    uint64 `json:"id" binding:"required"`                // 学号
	Posts string `json:"posts" binding:"required" msg:"请选择职位"` // 职位
}

func (AssociationApi) AssociationManageView(c *gin.Context) {
	var cr AssociationManageRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var association models.AssociationModel

	err := global.DB.Take(&association, "teacher_id = ?", claims.UserID).Error
	if err != nil {
		log.Print(err)
		res.FailWithMessage("还未创建协会", c)
		return
	}
	var associationMember models.AssociationMemberModel
	// 使用 db.Update 方法进行更新
	err = global.DB.Take(&associationMember, cr.ID).Update("Posts", cr.Posts).Error
	if err != nil {
		print(err)
		res.FailWithMessage("成员职位修改失败", c)
	}
	res.OkWithMessage("成员职位修改成功", c)
}

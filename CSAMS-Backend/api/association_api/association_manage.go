package association_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/ctype"
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

	if cr.Posts == "会长" {
		var user models.UserModel
		err := global.DB.Take(&user, cr.ID).Update("role", ctype.PermissionStudentAdmin).Error
		if err != nil {
			res.FailWithMessage("用户不存在", c)
			return
		}
		err = global.DB.Model(&association).Update("president", user.Name).Error
		if err != nil {
			res.FailWithMessage("协会不存在", c)
			return
		}
		res.OkWithMessage("会长信息同步成功", c)
	}
}

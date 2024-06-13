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

type AssociationQuitRequest struct {
	IDs []uint64 `json:"id" binding:"required"` // 学号
}

func (AssociationApi) AssociationQuitView(c *gin.Context) {
	var cr AssociationQuitRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var associationModel models.AssociationModel

	err := global.DB.Take(&associationModel, "teacher_id = ?", claims.UserID).Error
	if err != nil {
		log.Print(err)
		res.FailWithMessage("还未创建协会", c)
		return
	}

	for _, id := range cr.IDs {
		if id == associationModel.TeacherID {
			res.FailWithMessage("不能将教师踢出协会", c)
			continue
		}

		var associationMember models.AssociationMemberModel
		err = global.DB.Take(&associationMember, id).Error
		if err != nil {
			print(err)
			res.FailWithMessage("成员查找失败", c)
		}

		if associationMember.Posts == "会长" {
			err := global.DB.Model(&associationModel).Update("president", nil).Error
			if err != nil {
				res.FailWithMessage(err.Error(), c)
				return
			}
			err = global.DB.Model(&models.UserModel{}).Where("id = ?", id).Update("role", ctype.PermissionStudent).Error
			if err != nil {
				res.FailWithMessage(err.Error(), c)
				return
			}
		}

		err = global.DB.Delete(&associationMember).Error
		if err != nil {
			print(err)
			res.FailWithMessage("成员删除失败", c)
		}
	}
	res.OkWithMessage("成员删除成功", c)
}

package association_api

import (
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/service/association_ser"
	"CSAMS-Backend/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type AssociationCreateRequest struct {
	ID              uint64 `json:"id" binding:"required" msg:"请输入协会id"`               // 协会id
	AssociationName string `json:"association_name" binding:"required" msg:"请输入协会名称"` // 协会名称
	Introduction    string `json:"introduction" binding:"required" msg:"请输入协会简介"`     // 协会简介
}

func (AssociationApi) AssociationCreateView(c *gin.Context) {
	var cr AssociationCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	err := association_ser.AssociationService{}.CreateAssociation(cr.ID, cr.AssociationName, claims.Name, claims.UserID, cr.Introduction)
	if err != nil {
		log.Print(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("协会%s创建成功!", cr.AssociationName), c)
	return
}

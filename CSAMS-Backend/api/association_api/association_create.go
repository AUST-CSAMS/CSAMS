package association_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
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

	err := createAssociation(cr.ID, cr.AssociationName, claims.Name, claims.UserID, cr.Introduction)
	if err != nil {
		log.Print(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("协会%s创建成功!", cr.AssociationName), c)
	return
}

func createAssociation(id uint64, associationName, teacherName string, teacherID uint64, introduction string) error {
	var AssociationModel models.AssociationModel

	err := global.DB.Take(&AssociationModel, "teacher_id = ?", teacherID).Error
	if err == nil {
		return errors.New("已经拥有协会")
	}

	err = global.DB.Take(&AssociationModel, "id = ?", id).Error
	if err == nil {
		return errors.New("协会id已存在")
	}

	err = global.DB.Take(&AssociationModel, "association_name = ?", associationName).Error
	if err == nil {
		return errors.New("协会已存在")
	}

	date := time.Now().Format("2006-01-02")
	creatAt, _ := time.Parse("2006-01-02", date)

	// 入库
	err = global.DB.Create(&models.AssociationModel{
		ID:              id,
		AssociationName: associationName,
		CreateAt:        creatAt,
		TeacherName:     teacherName,
		TeacherID:       teacherID,
		Introduction:    introduction,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

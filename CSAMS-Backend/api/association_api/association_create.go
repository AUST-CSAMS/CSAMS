package association_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
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

	var AssociationModel models.AssociationModel

	err := global.DB.Take(&AssociationModel, "teacher_id = ?", claims.UserID).Error
	if err == nil {
		log.Print(err)
		res.FailWithMessage("已经拥有协会", c)
		return
	}

	err = global.DB.Take(&AssociationModel, "id = ?", cr.ID).Error
	if err == nil {
		log.Print(err)
		res.FailWithMessage("协会id已存在", c)
		return
	}

	err = global.DB.Take(&AssociationModel, "association_name = ?", cr.AssociationName).Error
	if err == nil {
		log.Print(err)
		res.FailWithMessage("协会已存在", c)
		return
	}

	// 入库
	err = global.DB.Create(&models.AssociationModel{
		ID:              cr.ID,
		AssociationName: cr.AssociationName,
		CreateAt:        time.Now(),
		TeacherName:     claims.Name,
		TeacherID:       claims.UserID,
		Introduction:    cr.Introduction,
	}).Error
	if err != nil {
		log.Print(err)
		res.FailWithMessage(fmt.Sprintf("协会%s创建失败!", cr.AssociationName), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("协会%s创建成功!", cr.AssociationName), c)

	// 更新协会成员表
	err = global.DB.Create(&models.AssociationMemberModel{
		ID:          cr.ID,                                   // 写入协会id，这里表上面没关联，但是后端代码关联了
		Users:       []models.UserModel{{ID: claims.UserID}}, // 将当前用户加入到关联表中
		Posts:       "负责老师",                                  // 设置职位信息
		JoiningTime: time.Now(),                              // 设置加入时间
	}).Error
	if err != nil {
		res.FailWithMessage("添加关联成员失败", c)
		return
	}

	// 成功响应
	res.OkWithMessage("添加关联成员成功", c)

	return
}

package association_ser

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"errors"
	"time"
)

func (AssociationService) CreateAssociation(id uint64, associationName, teacherName string, teacherID uint64, introduction string) error {
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

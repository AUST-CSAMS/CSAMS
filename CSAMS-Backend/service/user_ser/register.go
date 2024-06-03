package user_ser

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/ctype"
	"CSAMS-Backend/utils/pwd"
	"errors"
)

const Avatar = "/uploads/avatar/default.png"
const Integrity = 100

func (UserService) Register(id uint64, password, name string, age int, gender string, role ctype.Role, major string, tel uint64) error {
	// 判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "id = ?", id).Error
	if err == nil {
		return errors.New("学工号已存在")
	}
	// 对密码进行hash
	hashPwd := pwd.HashPwd(password)
	// 入库
	err = global.DB.Create(&models.UserModel{
		ID:        id,
		Password:  hashPwd,
		Name:      name,
		Age:       age,
		Gender:    gender,
		Avatar:    Avatar,
		Role:      role,
		Major:     major,
		Tel:       tel,
		Integrity: Integrity,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

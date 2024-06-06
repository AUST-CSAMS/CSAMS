package ctype

import (
	"encoding/json"
	"log"
)

type Role int

const (
	PermissionTeacher      Role = 1 // 教师
	PermissionStudent      Role = 2 // 学生
	PermissionStudentAdmin Role = 3 // 学生管理员
)

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *Role) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Print("解析失败:", err)
		return err
	}
	*r = toRole(s)
	return nil
}
func (r Role) String() string {
	var str string
	switch r {
	case PermissionTeacher:
		str = "教师"
	case PermissionStudent:
		str = "学生"
	case PermissionStudentAdmin:
		str = "学生管理员"
	default:
		str = "其他"
	}
	return str
}

func toRole(s string) Role {
	var r Role
	switch s {
	case "教师":
		r = PermissionTeacher
	case "学生":
		r = PermissionStudent
	case "学生管理员":
		r = PermissionStudentAdmin
	default:
		r = 0
	}
	return r
}

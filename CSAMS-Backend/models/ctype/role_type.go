package ctype

import "encoding/json"

type Role int

const (
	PermissionTeacher      Role = 1 // 教师
	PermissionStudentAdmin Role = 2 // 学生管理员
	PermissionStudent      Role = 3 // 学生
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case PermissionTeacher:
		str = "教师"
	case PermissionStudentAdmin:
		str = "学生管理员"
	case PermissionStudent:
		str = "学生"
	default:
		str = "其他"
	}
	return str
}

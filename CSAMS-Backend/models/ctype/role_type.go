package ctype

import "encoding/json"

type Role int

const (
	PermissionTeacher      Role = 1 // 教师
	PermissionStudent      Role = 2 // 学生
	PermissionStudentAdmin Role = 3 // 学生管理员
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
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

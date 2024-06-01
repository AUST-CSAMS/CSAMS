package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	ID   uint64 `json:"id"`   // 学工号
	Name string `json:"name"` // 姓名
	Role int    `json:"role"` // 权限 1 教师 2 学生 3 学生管理员
}

var MySecret []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

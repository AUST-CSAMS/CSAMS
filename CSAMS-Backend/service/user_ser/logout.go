package user_ser

import (
	"CSAMS-Backend/service/redis_ser"
	"CSAMS-Backend/utils/jwts"
	"time"
)

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	//获取过期时间
	exp := claims.ExpiresAt
	//获取现在时间
	now := time.Now()
	//相减
	diff := exp.Time.Sub(now)
	return redis_ser.Logout(token, diff)
}

package assignment_api

import (
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/utils/jwts"
	"github.com/gin-gonic/gin"
)

type AssignmentCreateRequest struct {
	ID      uint64 `json:"id" binding:"required"` // 活动id
	Content string `json:"content"  binding:"required" msg:"请输入作业内容"`
}

func (AssignmentApi) AssignmentCreateView(c *gin.Context) {
	var cr AssignmentCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var AssignmentModel models.AssignmentModel

}

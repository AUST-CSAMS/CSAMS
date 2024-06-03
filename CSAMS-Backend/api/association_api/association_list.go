package association_api

import (
	"CSAMS-Backend/models"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/service/common"
	"github.com/gin-gonic/gin"
)

func (AssociationApi) AssociationListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, _ := common.ComList(models.AssociationModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})

	res.OkWithList(list, count, c)
}

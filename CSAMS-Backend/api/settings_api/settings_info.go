package settings_api

import (
	"github.com/gin-gonic/gin"
	"sql_sever/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.FailWithCode(2, c)
}

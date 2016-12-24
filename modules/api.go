package components

import (
	"github.com/gin-gonic/gin"
	"github.com/itimofeev/vhustle/modules/gsheet"
	"github.com/itimofeev/vhustle/modules/util"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.LoggerWithWriter(util.GinLog.Writer()), gin.RecoveryWithWriter(util.RecLog.Writer()))

	api := r.Group("/api/v1")

	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to vhustle!")
	})

	api.GET("/contests", convertContext(gsheet.HandleGetContestList))

	admin := api.Group("/admin")
	admin.POST("/gsheet", func(c *gin.Context) {
		gsheet.UpdateContests()
	})

	return r
}

func convertContext(f func(ctx *util.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		newCtx := &util.Context{Context: c}
		f(newCtx)
	}
}

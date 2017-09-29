package forex

import (
	"net/http"

	"github.com/gin-gonic/gin"
	truefx "github.com/tonkla/gotruefx"
)

func GetAll(c *gin.Context) {
	result := truefx.NewFeed().Get()
	if len(result) > 0 {
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusNotFound, []truefx.Tick{})
	}
}

func GetBySymbol(c *gin.Context) {
	symbol := c.Param("symbol")
	result := truefx.NewFeed().GetBySymbol(symbol)
	if len(result) > 0 {
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusNotFound, []truefx.Tick{})
	}
}

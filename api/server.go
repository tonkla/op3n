package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	truefx "github.com/tonkla/gotruefx"
)

func main() {
	r := gin.Default()
	r.GET("/", home)
	r.GET("/forex", getForexAll)
	r.GET("/forex/:symbol", getForexBySymbol)
	r.Run(":3000")
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Please visit https://www.op3n.ga for document."})
}

func getForexAll(c *gin.Context) {
	result := truefx.NewFeed().Get()
	if len(result) > 0 {
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusNotFound, []truefx.Tick{})
	}
}

func getForexBySymbol(c *gin.Context) {
	symbol := c.Param("symbol")
	result := truefx.NewFeed().GetBySymbol(symbol)
	if len(result) > 0 {
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusNotFound, []truefx.Tick{})
	}
}

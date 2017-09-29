package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"github.com/tonkla/op3n/api/forex"
	"github.com/tonkla/op3n/api/province"
)

func main() {
	r := gin.Default()

	r.GET("/", home)

	r.GET("/forex", forex.GetAll)
	r.GET("/forex/:symbol", forex.GetBySymbol)

	r.GET("/thailand-provinces", province.GetAllProvinces)
	r.GET("/thailand-provinces/:province", province.GetProvince)
	r.GET("/thailand-provinces/:province/districts", province.GetAllDistricts)
	r.GET("/thailand-provinces/:province/districts/:district", province.GetDistrict)
	r.GET("/thailand-provinces/:province/districts/:district/subdistricts", province.GetAllSubdistricts)
	r.GET("/thailand-provinces/:province/districts/:district/subdistricts/:subdistrict", province.GetSubdistrict)
	r.GET("/thailand-provinces/postcodes", province.GetPostcodes)
	r.GET("/thailand-provinces/postcodes/:postcode", province.GetByPostcode)

	r.Run(":3000")
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Please visit https://www.op3n.ga for documentation."})
}

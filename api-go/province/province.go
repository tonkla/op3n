package province

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Province struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	NameEN   string  `json:"name_en"`
	Slug     string  `json:"slug"`
	Code     int     `json:"code"`
	RegionID int     `json:"region_id"`
	Lat      float32 `json:"lat"`
	Lng      float32 `json:"lng"`
}

type District struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	NameEN       string  `json:"name_en"`
	Code         int     `json:"code"`
	ProvinceID   int     `json:"province_id"`
	ProvinceCode int     `json:"province_code"`
	Lat          float32 `json:"lat"`
	Lng          float32 `json:"lng"`
}

type Subdistrict struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	NameEN       string  `json:"name_en"`
	Code         int     `json:"code"`
	Postcode     string  `json:"postcode"`
	DistrictID   int     `json:"district_id"`
	DistrictCode int     `json:"district_code"`
	ProvinceID   int     `json:"province_id"`
	ProvinceCode int     `json:"province_code"`
	Lat          float32 `json:"lat"`
	Lng          float32 `json:"lng"`
}

func GetAllProvinces(c *gin.Context) {
	var repo Repository
	provinces, err := repo.FindAllProvinces()
	if err != nil {
		c.JSON(http.StatusNotFound, []Province{})
	} else {
		c.JSON(http.StatusOK, provinces)
	}
}

func GetProvince(c *gin.Context) {
	p, _ := strconv.ParseInt(c.Param("province"), 10, 0)
	var repo Repository
	province, err := repo.FindProvince(p)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.JSON(http.StatusOK, province)
	}
}

func GetAllDistricts(c *gin.Context) {
	p, _ := strconv.ParseInt(c.Param("province"), 10, 0)
	var repo Repository
	districts, err := repo.FindDistrictsByProvince(p)
	if err != nil {
		c.JSON(http.StatusNotFound, []District{})
	} else {
		c.JSON(http.StatusOK, districts)
	}
}

func GetDistrict(c *gin.Context) {
	p, _ := strconv.ParseInt(c.Param("province"), 10, 0)
	d, _ := strconv.ParseInt(c.Param("district"), 10, 0)
	var repo Repository
	district, err := repo.FindDistrict(p, d)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.JSON(http.StatusOK, district)
	}
}

func GetAllSubdistricts(c *gin.Context) {
	p, _ := strconv.ParseInt(c.Param("province"), 10, 0)
	d, _ := strconv.ParseInt(c.Param("district"), 10, 0)
	var repo Repository
	subdistricts, err := repo.FindSubdistrictsByProvinceDistrict(p, d)
	if err != nil {
		c.JSON(http.StatusNotFound, []Subdistrict{})
	} else {
		c.JSON(http.StatusOK, subdistricts)
	}
}

func GetSubdistrict(c *gin.Context) {
	p, _ := strconv.ParseInt(c.Param("province"), 10, 0)
	d, _ := strconv.ParseInt(c.Param("district"), 10, 0)
	s, _ := strconv.ParseInt(c.Param("subdistrict"), 10, 0)
	var repo Repository
	subdistrict, err := repo.FindSubdistrict(p, d, s)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.JSON(http.StatusOK, subdistrict)
	}
}

func GetPostcodes(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func GetByPostcode(c *gin.Context) {
	p := c.Param("postcode")
	var repo Repository
	subdistricts, err := repo.FindSubdistrictsByPostcode(p)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.JSON(http.StatusOK, subdistricts)
	}
}

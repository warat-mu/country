package main

import (
	"country/addDB"
	"country/models"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	addDB.DBcheck()

	// addDB.NewCountry()
	// models.Test()
	// log.Print("start")
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", getCountrydata)
	router.GET("/desc", getCountrydataDesc)
	router.GET("/asc", getCountrydataAsc)
	router.GET("/:region", getCountryRegiondata)
	router.Run()

	// fmt.Println(models.GetCountry())
	// getCountryRegiondata()
}

func getCountrydata(c *gin.Context) {
	countrylist := models.GetCountry()
	// log.Println(countrylist)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func getCountrydataDesc(c *gin.Context) {

	countrylist := models.GetCountryDESC()

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func getCountrydataAsc(c *gin.Context) {

	countrylist := models.GetCountryASC()

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func getCountryRegiondata(c *gin.Context) {

	region := c.Param("region")
	countrylist := models.GetCountryByReion(region)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

// func getCountryRegiondata() {

// 	countrylist := models.GetCountryByReion()

// 	// if countrylist == nil || len(countrylist) == 0 {
// 	// 	c.AbortWithStatus(http.StatusNotFound)
// 	// } else {
// 	// 	c.IndentedJSON(http.StatusOK, countrylist)
// 	// }
// 	fmt.Println(len(countrylist))
// }

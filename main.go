package main

import (
	"country/addDB"
	"country/models"
	"database/sql"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	models.Db = testsql()

	addDB.DBcheck(models.Db)

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

func mysqlcon() *sql.DB {
	db, err := sql.Open("mysql", "tester:secret@tcp(host.docker.internal:3306)/api")
	if err != nil {
		return nil
	}

	return db
}
func getCountrydata(c *gin.Context) {
	countrylist := models.GetCountry(*models.Db)
	// log.Println(countrylist)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func getCountrydataDesc(c *gin.Context) {

	countrylist := models.GetCountryDESC(*models.Db)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func getCountrydataAsc(c *gin.Context) {

	countrylist := models.GetCountryASC(*models.Db)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func getCountryRegiondata(c *gin.Context) {

	region := c.Param("region")
	countrylist := models.GetCountryByReion(region, *models.Db)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func testsql() *sql.DB {
	db, err := sql.Open("mysql", "tester:secret@tcp(host.docker.internal:3306)/api")
	if err != nil {
		return nil
	}
	return db
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

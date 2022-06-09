package main

import (
	"context"
	"country/addDB"
	"country/env"
	"country/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ctx context.Context

func main() {

	ctx := context.Background()
	models.Db = testsql()

	addDB.DBcheck(models.Db)

	// addDB.NewCountry()
	// models.Test()
	// log.Print("start")
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(env.VerifyCache(ctx))
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

func setCache(c *gin.Context, cl []models.Data) {
	path := c.Request.URL.Path
	byte, err := json.Marshal(cl)
	if err != nil {
		log.Println(err)
	}
	cacheErr := env.Cache.Set(ctx, path, byte, 5*time.Second)
	if cacheErr != nil {
		log.Println(cacheErr.Err())
	}
}

func getCountrydata(c *gin.Context) {
	countrylist := models.GetCountry(*models.Db)
	// log.Println(countrylist)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		setCache(c, countrylist)
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func getCountrydataDesc(c *gin.Context) {

	countrylist := models.GetCountryDESC(*models.Db)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		setCache(c, countrylist)
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func getCountrydataAsc(c *gin.Context) {

	countrylist := models.GetCountryASC(*models.Db)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		setCache(c, countrylist)
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func getCountryRegiondata(c *gin.Context) {

	region := c.Param("region")
	countrylist := models.GetCountryByReion(region, *models.Db)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		setCache(c, countrylist)
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

package main

import (
	"context"
	"country/addDB"
	"country/env"
	"country/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx context.Context
var Cache *redis.Client
var Db *sql.DB

func main() {

	Db = Connectmysql()
	ctx := context.Background()
	Cache = env.Redisinit()
	addDB.DBcheck(Db)
	pong, err := Cache.Ping(ctx).Result()
	if err != nil {
		log.Println("ERROR")
		log.Println(err)

	}
	log.Println("OK======================")
	log.Println(pong)
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(env.VerifyCache(ctx, Cache))
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
	// path := c.Request.URL.Path
	// data, err := json.Marshal(cl)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(data)
	// Cache = env.Redisinit()
	// cacheErr := Cache.Set(ctx, path, data, 100*time.Second)
	// if cacheErr != nil {
	// 	panic(cacheErr)
	// }
}

func getCountrydata(c *gin.Context) {
	countrylist := models.GetCountry(Db)
	// log.Println(countrylist)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		// setCache(c, countrylist)
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func getCountrydataDesc(c *gin.Context) {
	countrylist := models.GetCountryDESC(Db)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		// setCache(c, countrylist)
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func getCountrydataAsc(c *gin.Context) {
	countrylist := models.GetCountryASC(Db)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		// setCache(c, countrylist)
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func getCountryRegiondata(c *gin.Context) {
	region := c.Param("region")
	countrylist := models.GetCountryByReion(region, Db)

	if countrylist == nil || len(countrylist) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		// setCache(c, countrylist)
		c.IndentedJSON(http.StatusOK, countrylist)
	}
}

func Connectmysql() *sql.DB {
	db, err := sql.Open("mysql", "tester:secret@tcp(host.docker.internal:3306)/api")
	if err != nil {
		return nil
	}
	return db
}

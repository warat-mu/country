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

func main() {
  dep := env.New(
    Connectmysql(),
    context.Background(),
    env.Redisinit(),
    )
	addDB.DBcheck(dep.Db)
	// pong, err := Cache.Ping(ctx).Result()
	// if err != nil {
	// 	log.Println("ERROR")
	// 	log.Println(err)

	// }
	// log.Println(pong)
	router := gin.Default()
	router.Use(cors.Default())

	router.Use(env.VerifyCache(dep))
	router.GET("/", getCountrydata(dep))
	router.GET("/desc", getCountrydataDesc(dep))
	router.GET("/asc", getCountrydataAsc(dep))
	router.GET("/:region", getCountryRegiondata(dep))
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

func setCache(c *gin.Context, cl []*models.Data, dep *env.Dep) {
  path := c.Request.URL.Path
  data, err := json.Marshal(cl)
  if err != nil {
    log.Println(err)
  }
  log.Println(data)
  cacheErr := dep.Cache.Set(dep.Ctx, path, data, 100*time.Second)
  if cacheErr != nil {
    panic(cacheErr)
  }
}


func getCountrydata(dep *env.Dep) gin.HandlerFunc {
  return func(c *gin.Context) {
      countrylist := models.GetCountry(dep.Db)
      // log.Println(countrylist)

      if countrylist == nil || len(countrylist) == 0 {
        c.AbortWithStatus(http.StatusNotFound)
      } else {
setCache(c, countrylist, dep)
        c.IndentedJSON(http.StatusOK, countrylist)
      }
  }
}

func getCountrydataDesc(dep *env.Dep) gin.HandlerFunc {
  return func (c *gin.Context) {
    countrylist := models.GetCountryDESC(dep.Db)

    if countrylist == nil || len(countrylist) == 0 {
      c.AbortWithStatus(http.StatusNotFound)
    } else {
      setCache(c, countrylist, dep)
      c.IndentedJSON(http.StatusOK, countrylist)
    }
  }
}

func getCountrydataAsc(dep *env.Dep) gin.HandlerFunc {
  return func(c *gin.Context) {
    countrylist := models.GetCountryASC(dep.Db)

    if countrylist == nil || len(countrylist) == 0 {
      c.AbortWithStatus(http.StatusNotFound)
    } else {
setCache(c, countrylist, dep)
      c.IndentedJSON(http.StatusOK, countrylist)
    }
  }
}

func getCountryRegiondata(dep *env.Dep) gin.HandlerFunc {
  return func(c *gin.Context) {
    region := c.Param("region")
    countrylist := models.GetCountryByReion(region, dep.Db)

    if countrylist == nil || len(countrylist) == 0 {
      c.AbortWithStatus(http.StatusNotFound)
    } else {
setCache(c, countrylist, dep)
      c.IndentedJSON(http.StatusOK, countrylist)
    }
  }

}

func Connectmysql() *sql.DB {
	db, err := sql.Open("mysql", "tester:secret@tcp(host.docker.internal:3306)/api")
	if err != nil {
		return nil
	}
	return db
}

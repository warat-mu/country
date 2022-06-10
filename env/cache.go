package env

import (
	"country/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func Redisinit() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func VerifyCache(dep *Dep) gin.HandlerFunc {
  return func(c *gin.Context) {
    path := c.Request.URL.Path
    data, err := dep.Cache.Get(dep.Ctx, path).Bytes()
    if err != nil {
      c.Next()
    } else {
      var countries []*models.Data
      // To json
      json.Unmarshal(data, &countries)
      c.JSON(http.StatusOK, countries)
      c.Abort()
    }
  }
}

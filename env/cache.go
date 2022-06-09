package env

import (
	"context"
	"country/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func Redisinit() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:6379",
		Password: "",
		DB:       0,
	})
}

func VerifyCache(ctx context.Context, cache *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		data, err := cache.Get(ctx, path).Bytes()
		if err != nil {
			c.Next()
		} else {
			var countries []models.Data
			// To json
			json.Unmarshal(data, &countries)
			c.JSON(http.StatusOK, countries)
			c.Abort()
		}
	}
}

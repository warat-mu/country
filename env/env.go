package env

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v8"
)

func New(db *sql.DB, ctx context.Context, cache *redis.Client) *Dep {
  return &Dep {
    Db: db,
    Ctx: ctx,
    Cache: cache,
  }
}

type Dep struct {
  Db *sql.DB
  Ctx context.Context
  Cache *redis.Client
}

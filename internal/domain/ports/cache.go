package ports

import (
	"time"
)

type CacheRepository interface {
	Set(key string, data any, timeExpiration time.Duration) error
	Get(key string, src any) error
	Delete(key string) error
}

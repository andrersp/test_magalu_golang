package cache

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/andrersp/favorites/internal/domain/ports"
	dbCache "github.com/patrickmn/go-cache"
)

const (
	defaultExpiration = 5
	maxExpiration     = 10
)

type cacheRepository struct {
	db *dbCache.Cache
}

func NewCacheRepository() ports.CacheRepository {
	repository := new(cacheRepository)
	repository.db = dbCache.New(defaultExpiration*time.Minute, maxExpiration*time.Minute)

	return repository
}

// Delete implements ports.CacheRepository.
func (c *cacheRepository) Delete(key string) error {
	c.db.Delete(key)
	return nil
}

// Get implements ports.CacheRepository.
func (c *cacheRepository) Get(key string, src any) error {
	data, found := c.db.Get(key)

	if !found {
		return errors.New("data not found")
	}

	bts, _ := json.Marshal(data)

	return json.Unmarshal(bts, src)
}

// Set implements ports.CacheRepository.
func (c *cacheRepository) Set(key string, data any, timeExpiration time.Duration) error {
	c.db.Set(key, data, timeExpiration)

	return nil
}

package storage

import (
	"context"
	"time"

	"github.com/patrickmn/go-cache"
)

type StorageService struct {
	db *cache.Cache
}

var (
	storageService = &StorageService{}
	ctx            = context.Background
)

const CacheDuration = 6 * time.Hour

func InitializeStorage() *StorageService {
	storageService.db = cache.New(CacheDuration, 10*time.Minute)
	return storageService
}

func SaveUrlMapping(shortUrl, originalUrl string) {
	storageService.db.Set(shortUrl, originalUrl, cache.DefaultExpiration)
}

func RetrieveInitialUrl(shortUrl string) string {
	item, ok := storageService.db.Get(shortUrl)
	if !ok {
		return ""
	}
	return item.(string)
}

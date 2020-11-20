package store

import (
	"time"

	"github.com/fioncat/go-gendb/misc/log"
)

var cacheTTL = time.Hour

var enableCache bool

func SetCacheTTL(ttl time.Duration) {
	if ttl == 0 {
		log.Errorf("can not set cache-ttl to zero.")
		return
	}
	cacheTTL = ttl
}

func EnableCache() {
	enableCache = true
}

func SaveCache(base, key string, v interface{}) error {
	if !enableCache {
		return nil
	}
	err := Save("cache_"+base, key, v, cacheTTL)
	if err != nil {
		return err
	}
	log.Infof("save cache base=%s key=%s ttl=%s",
		base, key, cacheTTL.String())
	return nil
}

func GetCache(base, key string, v interface{}) bool {
	if !enableCache {
		return false
	}
	found, err := Get("cache_"+base, key, v)
	if !found {
		return false
	}
	if err != nil {
		log.Errorf("read cache %s/%s failed: %v", base, key, err)
		return false
	}
	return true
}

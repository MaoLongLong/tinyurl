package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStorageService *StorageService

func init() {
	testStorageService = InitializeStorage()
}

func TestStorageInit(t *testing.T) {
	assert.NotNil(t, testStorageService.db)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortURL := "Jsz4k57oAX"

	SaveUrlMapping(shortURL, initialLink)
	retrievedUrl := RetrieveInitialUrl(shortURL)
	assert.Equal(t, initialLink, retrievedUrl)
}

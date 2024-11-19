package handlers

import (
	"github.com/bradfitz/gomemcache/memcache"
	"log"
	"os"
)

var MemClient *memcache.Client

func InitMemcached() {
	// Get Memcached host and port from environment variables
	host := os.Getenv("MEMCACHED_HOST")
	port := os.Getenv("MEMCACHED_PORT")
	if host == "" || port == "" {
		log.Fatal("MEMCACHED_HOST or MEMCACHED_PORT not set")
	}

	// Connect to Memcached
	MemClient = memcache.New(host + ":" + port)

	// Test connection
	err := MemClient.Set(&memcache.Item{
		Key:   "test",
		Value: []byte("connected"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to Memcached: %v", err)
	}

	log.Println("Connected to Memcached successfully")
}

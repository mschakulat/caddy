package tools

import (
	"caddy/src/config"
	"encoding/gob"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Filecache struct {
	cacheItemName string
	ttlFile       string
	ttl           time.Duration
}

func VersionCache() *Filecache {
	return &Filecache{
		cacheItemName: "item",
		ttlFile:       "timestamp",
		ttl:           60 * time.Second,
	}
}

func (c *Filecache) Has(key string) bool {
	if c.isExpired(key) {
		return false
	}

	dirPath := filepath.Join(config.SystemPaths.Cache, key, c.cacheItemName)
	_, err := os.Stat(dirPath)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func (c *Filecache) Set(key string, item []string) {
	path := c.getItemPathByName(key)

	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create directories: %v", err)
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(item); err != nil {
		fmt.Println(err)
	}

	c.setTTL(key)
}

func (c *Filecache) Get(key string) *[]string {
	if c.isExpired(key) {
		return nil
	}

	file, err := os.Open(c.getItemPathByName(key))
	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}
	defer file.Close()

	var data []string
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}

	return &data
}

func (c *Filecache) getItemPathByName(key string) string {
	return filepath.Join(config.SystemPaths.Cache, key, c.cacheItemName)
}

func (c *Filecache) setTTL(key string) {
	timestamp := time.Now().Add(c.ttl).Format(time.RFC3339)
	ttlFilePath := filepath.Join(config.SystemPaths.Cache, key, c.ttlFile)
	err := os.WriteFile(ttlFilePath, []byte(timestamp), 0644)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func (c *Filecache) isExpired(key string) bool {
	ttlFilePath := filepath.Join(config.SystemPaths.Cache, key, c.ttlFile)

	_, err := os.Stat(ttlFilePath)
	if os.IsNotExist(err) {
		return false
	}

	file, _ := os.ReadFile(ttlFilePath)

	ttl, _ := time.Parse(time.RFC3339, string(file))

	if time.Now().After(ttl) {
		_ = os.RemoveAll(filepath.Join(config.SystemPaths.Cache, key))
		return true
	}

	return false
}

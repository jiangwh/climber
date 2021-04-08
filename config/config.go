package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
)

type DbConfig struct {
	Source, Driver string
}

var (
	cfg     *DbConfig
	once    sync.Once
	cfgLock = new(sync.RWMutex)
)

func GetDbConfig() *DbConfig {
	once.Do(ReloadConfig)
	cfgLock.RLock()
	defer cfgLock.RUnlock()
	return cfg
}

func GetDbFilePath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal("获取目录失败")
	}
	fmt.Println(currentPath)
	if strings.HasSuffix(currentPath, "climber") {
		return strings.ReplaceAll(currentPath, "climber", "climber/config") + "/sqlite3.toml"
	}

	return currentPath + "/sqlite3.toml"
}

func ReloadConfig() {
	filePath, err := filepath.Abs(GetDbFilePath())
	if err != nil {
		panic(err)
	}
	fmt.Printf("parse db config once. filePath: %s\n", filePath)
	config := new(DbConfig)
	if _, err := toml.DecodeFile(filePath, config); err != nil {
		panic(err)
	}
	cfgLock.Lock()
	defer cfgLock.Unlock()
	cfg = config
}

package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	MongoDB    string
	ServerPort string
}

func Load() (Config, error) {
	// godotenv.Load 读取 .env 环境文件
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("Failed to load .env")
	}

	// 分别获取环境变量
	mongoURI, err := extractEnv("MONGO_URI")
	if err != nil {
		return Config{}, err
	}

	mongoDB, err := extractEnv("MONGO_DB_NAME")
	if err != nil {
		return Config{}, err
	}

	port, err := extractEnv("PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{
		MongoURI:   mongoURI,
		MongoDB:    mongoDB,
		ServerPort: port,
	}, nil
}

// 封装函数检查env文件是否为空
func extractEnv(key string) (string, error) {

	// os.Getenv 读取 env 文件
	val := os.Getenv(key)

	if val == "" {
		return "", fmt.Errorf("Missing request env")
	}

	return val, nil
}

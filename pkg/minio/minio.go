package minio

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"auth-service-backend/config"
	"auth-service-backend/internal/database/redis"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	redisDB        *redis.Redis
	minioClient    *minio.Client
	lifetime       = 24 * 60 * 60 * time.Second
	minioBucket    = ""
	minioHost      = ""
	minioAccessKey = ""
	minioSecretKey = ""
)

func Initialization(rDB *redis.Redis) bool {
	minioBucket = config.AppConfig.MinioBucket
	minioHost = config.AppConfig.MinioHost
	minioAccessKey = config.AppConfig.MinioAccessKey
	minioSecretKey = config.AppConfig.MinioSecretKey
	redisDB = rDB

	var err error
	minioClient, err = minio.New(minioHost, &minio.Options{
		Creds: credentials.NewStaticV4(
			minioAccessKey,
			minioSecretKey,
			"",
		),
		Secure: true,
	})
	if err != nil {
		log.Fatalf("Unable to initialize MinIO client: %v", err)
	}

	return true
}

func GetDataUrl(ctx context.Context, dataName string) (string, error) {
	cache, err := redisDB.Get(dataName)
	if err == nil && cache != "" {
		return cache, nil
	}

	presignedURL, err := minioClient.PresignedGetObject(ctx, minioBucket, dataName, lifetime, nil)
	if err != nil {
		return "", fmt.Errorf("something went wrong with MinIO: %w", err)
	}

	err = redisDB.Set(dataName, presignedURL.String())
	if err != nil {
		return "", fmt.Errorf("failed to set cache: %w", err)
	}

	return presignedURL.String(), nil
}

func GetFile(ctx context.Context, dataName string) ([]byte, error) {
	object, err := minioClient.GetObject(ctx, minioBucket, dataName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("something went wrong with MinIO: %w", err)
	}
	defer object.Close()

	fileContent := []byte{}
	buffer := make([]byte, 1024)
	for {
		n, err := object.Read(buffer)
		if n > 0 {
			fileContent = append(fileContent, buffer[:n]...)
		}
		if err != nil {
			break
		}
	}

	return fileContent, nil
}

func PutData(ctx context.Context, file []byte, userId string, filename string) (string, error) {
	objectName := fmt.Sprintf("%s/%s", userId, filename)
	_, err := minioClient.PutObject(ctx, minioBucket, objectName, bytes.NewReader(file), int64(len(file)), minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return "The image has been uploaded", nil
}

package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"auth-service-backend/config"
)

func RunMigrations() error {
	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DB,
	)

	cmd := exec.Command("migrate", "-path", "./internal/database/postgres/migrations", "-database", databaseURL, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return errors.New("failed to apply migrations")
	} else {
		log.Println("Migrations applied successfully")
		return nil
	}
}

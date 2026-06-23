package main

import (
	"context"
	"fmt"
	"log"
	"message-platform/configs"
	"message-platform/ent"

	_ "github.com/lib/pq"
)

func init() {

}

func main() {
	systemConfig := configs.LoadConfig()
	setupDatabase(systemConfig)

}

func setupDatabase(systemConfig configs.Config) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		systemConfig.Database.Host, systemConfig.Database.Port, systemConfig.Database.Username, systemConfig.Database.Name, systemConfig.Database.Password)
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

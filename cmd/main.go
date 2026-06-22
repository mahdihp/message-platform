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
	system_Config := configs.LoadConfig()
	setupDB(system_Config)

}

func setupDB(system_Config configs.Config) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		system_Config.Database.Host, system_Config.Database.Port, system_Config.Database.Username, system_Config.Database.Name, system_Config.Database.Password)
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

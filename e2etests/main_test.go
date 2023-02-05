// go:build e2e
package e2etests

import (
	"log"
	"os"
	"testing"

	"github.com/Pochirify/pochirify-backend/internal/domain/repository"
	"github.com/joho/godotenv"
)

var (
	port         string
	projectID    string
	instanceID   string
	databaseID   string
	repositories repository.Repositories
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("failed to load env file")
		os.Exit(1)
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	projectID = os.Getenv("GCP_PROJECT_ID")
	if port == "" {
		projectID = "pochirify-dev"
	}
	instanceID = os.Getenv("SPANNER_INSTANCE_ID")
	if instanceID == "" {
		instanceID = "pochirify"
	}
	databaseID = os.Getenv("SPANNER_DATABASE_ID")
	if databaseID == "" {
		databaseID = "pochirify-server"
	}

	// // TODO: we should manipulate env based on e2etest mode
	// credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	// if err := exec.Command("sh", "-c", fmt.Sprintf("export GOOGLE_APPLICATION_CREDENTIALS=.%s", credentials)).Run(); err != nil {
	// 	log.Printf("failed to set GOOGLE_APPLICATION_CREDENTIALS: %s", err.Error())
	// 	os.Exit(1)
	// }

	repositories = initRepositories()

	os.Exit(m.Run())
}
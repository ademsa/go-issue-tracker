package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	infraPersistence "go-issue-tracker/pkg/infrastructure/persistence"
	"go-issue-tracker/pkg/interfaces/externalapi"
	"go-issue-tracker/pkg/interfaces/externalapimock"
	"go-issue-tracker/pkg/interfaces/http"
	"go-issue-tracker/pkg/interfaces/persistence"
	"go-issue-tracker/pkg/usecases"
	"log"
	netHttp "net/http"
	"time"
)

// EndpointsBaseAddress is base path
var EndpointsBaseAddress = "127.0.0.1:8000"

func main() {
	// Get db path
	dbPath, err := infraPersistence.GetDefaultSQLiteDBFilePath()
	if err != nil {
		log.Fatal(err)
	}

	// Connecting to SQLite database
	db, err := infraPersistence.GetSQLiteDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// SQLite repositories
	ir := persistence.NewSQLiteIssueRepository(db)
	lr := persistence.NewSQLiteLabelRepository(db)
	pr := persistence.NewSQLiteProjectRepository(db)

	// HTTP Client for making calls to color repository
	httpClient := &netHttp.Client{
		Timeout: 30 * time.Second,
	}

	// External API repository
	cr := externalapi.NewColorRepository("http://"+EndpointsBaseAddress+externalapimock.ExternalAPIMockPath, httpClient)

	// Use Cases
	iuc := usecases.NewIssueUseCase(ir)
	luc := usecases.NewLabelUseCase(lr)
	puc := usecases.NewProjectUseCase(pr)
	cuc := usecases.NewColorUseCase(cr)

	httpServer, ruc := http.PrepareHTTPServer(iuc, luc, puc, cuc)

	http.PrepareEndpoints(httpServer, ruc)

	// Preparing mock endpoint to simulate External Api
	externalapimock.PrepareEndpoints(httpServer)

	// Start server
	httpServer.Logger.Fatal(httpServer.Start(EndpointsBaseAddress))
}

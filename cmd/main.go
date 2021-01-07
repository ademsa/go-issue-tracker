package main

import (
	"flag"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go-issue-tracker/pkg/domain"
	"go-issue-tracker/pkg/infrastructure/database"
	"go-issue-tracker/pkg/infrastructure/helpers"
	"go-issue-tracker/pkg/interfaces/externalapi"
	"go-issue-tracker/pkg/interfaces/externalapimock"
	"go-issue-tracker/pkg/interfaces/gql"
	"go-issue-tracker/pkg/interfaces/grpc"
	"go-issue-tracker/pkg/interfaces/persistence"
	"go-issue-tracker/pkg/interfaces/rest"
	"go-issue-tracker/pkg/usecases"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

// EndpointBaseAddress is base path
var EndpointBaseAddress = "0.0.0.0:3001"

// GRPCEndpointAddress is path to gRPC Color Service
var GRPCEndpointAddress = "0.0.0.0:3002"

func main() {
	grpcStatus := flag.Bool("grpc", false, "Use gRPC Color Service")
	flag.Parse()

	// Get db path
	dbPath, err := database.GetDefaultSQLiteDBFilePath()
	if err != nil {
		log.Fatal(err)
	}

	// Connecting to SQLite database
	db, err := database.GetSQLiteDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// SQLite repositories
	ir := persistence.NewSQLiteIssueRepository(db)
	lr := persistence.NewSQLiteLabelRepository(db)
	pr := persistence.NewSQLiteProjectRepository(db)

	// Use Cases
	iuc := usecases.NewIssueUseCase(ir)
	luc := usecases.NewLabelUseCase(lr)
	puc := usecases.NewProjectUseCase(pr)

	var cr domain.ColorRepository
	if *grpcStatus == true {
		// External gRPC Color Service
		cr = grpc.NewColorRepository(GRPCEndpointAddress)
	} else {
		// HTTP Client for making calls to color repository
		httpClient := &http.Client{
			Timeout: 30 * time.Second,
		}

		// External API repository
		cr = externalapi.NewColorRepository("http://"+EndpointBaseAddress+externalapimock.ExternalAPIMockPath, httpClient)
	}

	cuc := usecases.NewColorUseCase(cr)

	// Prepare HTTP Server
	httpServer := rest.PrepareServer()

	// Preparing mock endpoint to simulate External Api
	externalapimock.PrepareEndpoints(httpServer)

	// REST
	restManager := rest.NewManager(iuc, luc, puc, cuc)
	rootDirPath, err := helpers.GetProjectDirPath()
	uiDirPath := filepath.Join(rootDirPath, "ui")
	if err != nil {
		log.Fatal(err)
	}
	rest.PrepareEndpoints(httpServer, restManager, uiDirPath)

	// GraphQL
	gqlSchema := gql.PrepareGraphQL(iuc, luc, puc, cuc)
	gqlManager := gql.NewRequestManager(gqlSchema)
	gql.PrepareEndpoints(httpServer, gqlManager)

	// Start HTTP server
	httpServer.Logger.Fatal(httpServer.Start(EndpointBaseAddress))
}

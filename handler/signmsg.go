package handler

import (
	"fmt"
	"net/http"
	"os"
	"server/logger"
)

func SignAuthwalletHandler(w http.ResponseWriter, r *http.Request) {
	config := logger.NewConfigFromEnv()

	// Initialize logger
	log, err := logger.NewLogger(config)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()
	log.Info("NFT retrieved from ")

	

}

//remove nft handler
func SignReqwalletHandler(w http.ResponseWriter, r *http.Request) {
	config := logger.NewConfigFromEnv()

	// Initialize logger
	log, err := logger.NewLogger(config)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()
	log.Info("NFT removed from ")
	
}

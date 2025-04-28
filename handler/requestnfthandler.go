package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"server/logger"
	"server/util"
)

/*****************
Req Blockchain NFT
*******************/

//store nft handler
func StoreReqNFTHandler(w http.ResponseWriter, r *http.Request) {

	config := logger.NewConfigFromEnv()

    // Initialize logger
    log, err := logger.NewLogger(config)
    if err != nil {
        fmt.Printf("Failed to initialize logger: %v\n", err)
        os.Exit(1)
    }
    defer log.Sync()
	//log.Info("NFT stored in ")
	var req struct {
		RequestWallPubAddr string `json:"requestwallPubAddr"`
		Nft string `json:"nft_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	conAddr, err := util.GetConAddrByReqPubKey(req.RequestWallPubAddr)
	if err != nil {
		http.Error(w, "Error retrieving wallet address: "+err.Error(), http.StatusInternalServerError)
		log.Error("Error retrieving wallet address: " + err.Error())
		return
	}
	if conAddr == "" {
		http.Error(w, "No wallet found for the given request address", http.StatusNotFound)
		log.Error("No wallet found for the given request address")
		return
	}
	SendMessageToClient(conAddr, "setreqnft "+req.Nft)
	log.Info("NFT stored in ")
	
}
//get nft handler
func GetReqNFTHandler(w http.ResponseWriter, r *http.Request) {
	config := logger.NewConfigFromEnv()

	// Initialize logger
	log, err := logger.NewLogger(config)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()
	log.Info("NFT retrieved from ")

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Handle the request to get the NFT
	var req struct {
		RequestWallPubAddr string `json:"requestwallPubAddr"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	conAddr, err := util.GetConAddrByReqPubKey(req.RequestWallPubAddr)
	if err != nil {
		http.Error(w, "Error retrieving wallet address: "+err.Error(), http.StatusInternalServerError)
		log.Error("Error retrieving wallet address: " + err.Error())
		return
	}
	if conAddr == "" {
		http.Error(w, "No wallet found for the given request address", http.StatusNotFound)
		log.Error("No wallet found for the given request address")
		return
	}
	response, err  := SendMessageToClient(conAddr, "getreqnft")
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Preprocess the response to extract the NFT ID
	var nftID string
	_, err = fmt.Sscanf(response, "NFT_AUTH %s", &nftID)
	if err != nil {
		http.Error(w, "Failed to parse NFT ID", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"nft": nftID,
	})
}

//remove nft handler
func RemoveReqNFTHandler(w http.ResponseWriter, r *http.Request) {
	config := logger.NewConfigFromEnv()

	// Initialize logger
	log, err := logger.NewLogger(config)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Handle the request to get the NFT
	var req struct {
		RequestWallPubAddr string `json:"requestwallPubAddr"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	conAddr, err := util.GetConAddrByReqPubKey(req.RequestWallPubAddr)
	if err != nil {
		http.Error(w, "Error retrieving wallet address: "+err.Error(), http.StatusInternalServerError)
		log.Error("Error retrieving wallet address: " + err.Error())
		return
	}
	if conAddr == "" {
		http.Error(w, "No wallet found for the given request address", http.StatusNotFound)
		log.Error("No wallet found for the given request address")
		return
	}
	response, err :=SendMessageToClient(conAddr, "removereqnft")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"server/util"
	"strings"
	"sync"
)


var (
	clients   = make(map[string]net.Conn) // Map to store active clients
	clientsMu sync.Mutex                  // Mutex to synchronize access to the clients map
)

func HandleConnection(conn net.Conn) {
	defer func() {
		// Remove client on disconnect
		clientsMu.Lock()
		delete(clients, conn.RemoteAddr().String())
		clientsMu.Unlock()
		conn.Close()
	}()

	fmt.Println("Socket server: Connection established with", conn.RemoteAddr())
	scanner := bufio.NewScanner(conn)
	fmt.Println("Socket server: Waiting for authentication data from", conn.RemoteAddr())
	// First, authenticate the client
	if scanner.Scan() {
		authDataString := scanner.Text()
		fmt.Println("Socket server received from", conn.RemoteAddr(), ":", authDataString)
		parts := strings.Split(authDataString, ",")

		if len(parts) != 2 {
			fmt.Println("Socket server: Invalid auth data format from", conn.RemoteAddr())
			conn.Write([]byte("Invalid authentication data format. Closing connection.\n"))
			return
		}

		requestPubAddr := strings.TrimSpace(parts[0])
		authPubAddr := strings.TrimSpace(parts[1])

		if util.ValidateAuthaddr(authPubAddr) {
			if util.ValidateReqaddr(requestPubAddr) {
				conn.Write([]byte("VALIDATED\n"))
				fmt.Println("Authentication successful for", conn.RemoteAddr())
				util.AddWallet(conn.RemoteAddr().String(), requestPubAddr, authPubAddr)

				// Add the client to the map AFTER authentication
				clientsMu.Lock()
				clients[conn.RemoteAddr().String()] = conn
				clientsMu.Unlock()

				// Maintain the connection with the client
				for scanner.Scan() {
					message := scanner.Text()
					fmt.Printf("Received message from %s: %s\n", conn.RemoteAddr(), message)
					// Echo the message back to the client
					_, err := conn.Write([]byte("Echo: " + message + "\n"))
					if err != nil {
						fmt.Println("Error writing to connection:", err)
						util.DeleteWallet(conn.RemoteAddr().String())
						break
					}
				}
			} else {
				conn.Write([]byte("Invalid request address. Closing connection.\n"))
				fmt.Println("Invalid request address for", conn.RemoteAddr())
				return
			}
		} else {
			conn.Write([]byte("Invalid authentication address. Closing connection.\n"))
			fmt.Println("Invalid authentication address for", conn.RemoteAddr())
			return
		}
	} else {
		fmt.Println("Socket server: Failed to read authentication message from", conn.RemoteAddr())
		return
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from connection:", err)
	}
	fmt.Println("Socket server: Connection closed with", conn.RemoteAddr())
}

// HTTP handler to send a message to a specific client
func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ClientAddr string `json:"clientAddr"`
		Message    string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := SendMessageToClient(req.ClientAddr, req.Message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message sent successfully"))
}

// Function to send a message to a specific client


func SendMessageToClient(clientAddr string, command string) error {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	conn, exists := clients[clientAddr]
	if !exists {
		return fmt.Errorf("client %s not found", clientAddr)
	}

	_, err := fmt.Fprintf(conn, "%s\n", command)
	return err
}

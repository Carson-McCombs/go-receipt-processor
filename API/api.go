package api

import (
	points "GoReceiptProcessor/Points"
	receipt "GoReceiptProcessor/Receipt"

	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
	receiptMap map[string]receipt.Receipt
	pointsMap  map[string]int64
}

func NewServer() *Server {
	server := &Server{
		Router:     mux.NewRouter(),
		receiptMap: make(map[string]receipt.Receipt),
		pointsMap:  make(map[string]int64),
	}
	server.routes()
	return server
}

// All possible ways of interacting with the server
// API Routes
func (s *Server) routes() {
	s.HandleFunc("/receipts/process", s.processReceipt).Methods("POST")
	s.HandleFunc("/receipts/{id}", s.getReceiptPoints).Methods("GET")
}

func (s *Server) processReceipt(w http.ResponseWriter, r *http.Request) {
	var unparsedReceipt receipt.UnparsedReceipt
	err := json.NewDecoder(r.Body).Decode(&unparsedReceipt)
	if err != nil {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}
	id := uuid.New().String()
	receipt, _ := receipt.ParseReceipt(id, unparsedReceipt, true)
	points := points.CalculatePoints(receipt)
	s.receiptMap[id] = receipt
	s.pointsMap[id] = points
	err = json.NewEncoder(w).Encode(receipt.Id)
	if err != nil {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

}

type pointsResponse struct {
	Points int64 `json:"points"`
}

// On GET HTTP Request, tries to parse the ID given within the request and iterates over the list of ReceiptPoints objects until it finds it or finishes iterating over the list.
// It then outputs the receipt points value in JSON format.
func (s *Server) getReceiptPoints(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	points, containsKey := s.pointsMap[id]
	if !containsKey {
		http.Error(w, "No receipt found for that id", http.StatusNotFound)
		return
	}
	//pointsOutput := map[string]int64{"points": points}
	pointsOutput := pointsResponse{Points: points}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(pointsOutput)
	if err != nil {
		http.Error(w, "No receipt found for that id", http.StatusNotFound)
		return
	}
}

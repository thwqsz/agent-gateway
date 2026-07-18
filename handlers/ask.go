package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thwqsz/agent-gateway/client"
)

type AskRequest struct {
	Text string `json:"text"`
}

type AskResponse struct {
	Answer string `json:"answer"`
}

var agentClient = client.NewAgentClient("http://localhost:8000")

func AskHandler(w http.ResponseWriter, r *http.Request) {
	var req AskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	answer, err := agentClient.Ask(req.Text)
	if err != nil {
		http.Error(w, "agent service error: "+err.Error(), http.StatusBadGateway)
		return
	}

	response := AskResponse{Answer: answer}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

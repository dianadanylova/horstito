package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"horstito/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var sessions = make(map[string]models.Session)

func CreateSession(w http.ResponseWriter, r *http.Request) {
	var session models.Session
	_ = json.NewDecoder(r.Body).Decode(&session)

	session.ID = uuid.New().String()
	session.StartTime = time.Now()
	session.EndTime = session.StartTime.Add(25 * time.Minute)

	sessions[session.ID] = session
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

func GetAllSessions(w http.ResponseWriter, r *http.Request) {
	var all []models.Session
	for _, s := range sessions {
		all = append(all, s)
	}
	json.NewEncoder(w).Encode(all)
}

func GetSessionByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	session, ok := sessions[id]
	if !ok {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(session)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	delete(sessions, id)
	w.WriteHeader(http.StatusNoContent)
}

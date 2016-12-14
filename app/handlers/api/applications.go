package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/generationtux/brizo/database"
	"github.com/generationtux/brizo/resources"
)

// Applications provides a listing of all Applications
func Applications(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		log.Printf("Database error: '%s'\n", err)
		http.Error(w, "there was an error when attempting to connect to the database", http.StatusInternalServerError)
		return
	}
	apps, err := resources.AllApplications(db)
	if err != nil {
		log.Printf("Error when retrieving applications: '%s'\n", err)
		http.Error(w, "there was an error when retrieving applications", http.StatusInternalServerError)
		return
	}
	for i := range apps {
		if len(apps[i].Pods) == 0 {
			apps[i].Pods = make([]resources.Pod, 0)
		}
	}
	json.NewEncoder(w).Encode(apps)
}
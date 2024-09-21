package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lgustavopalmieri/ultrafactory-oee/internal/adapters/web/server"
)

func SetupTestRoutes(s *server.WebServer, db *sql.DB) {
	s.AddHandler("/test", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Ultrafactory OEE on Kubernetees VERY NICEE!!!!IT REALLY WORTHS!")
}
package main

import (
	"encoding/json"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
	_ "restapi/docs"
)

// getAllScansByCustomerID godoc
// @Summary      Get scans by customer ID
// @Description  Retrieves scan results filtered by a customer ID
// @Tags         scans
// @Accept       json
// @Produce      json
// @Param        customerID  query     string  true  "Customer ID"
// @Success      200  {array}  Scan
// @Failure      400  {string}  string  "Missing parameter"
// @Failure      500  {string}  string  "Internal error"
// @Router       /api/scansByCustomerID [get]
func getAllScansByCustomerID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Scans); err != nil {
		http.Error(w, "Failed to encode resources", http.StatusInternalServerError)
	}
}

// getScanDetails godoc
// @Summary      Get scan details by scan ID
// @Description  Retrieves scan results filtered by a customer ID
// @Tags         scans
// @Accept       json
// @Produce      json
// @Param        scanID  query     string  true  "Scan ID"
// @Success      200  {array}  Scan
// @Failure      400  {string}  string  "Missing parameter"
// @Failure      500  {string}  string  "Internal error"
// @Router       /api/scanDetails [get]
func getScanDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Scans); err != nil {
		http.Error(w, "Failed to encode resources", http.StatusInternalServerError)
	}
}

// @title           TT Rest API
// @version         1.0
// @description     Rest API for TT APP
// @host            localhost:8080
// @BasePath        /
func main() {
	http.HandleFunc("/api/scansByCustomerID", getAllScansByCustomerID)
	http.HandleFunc("/api/scanDetails", getScanDetails)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	log.Println("Server starting at :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

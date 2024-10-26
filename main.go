package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

var receipts = make(map[string]int)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/receipts/process", processReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", getReceiptPoints).Methods("GET")

	http.ListenAndServe(":8081", r)
}

func processReceipt(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /receipts/process endpoint hit")
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "JSON is invalid", http.StatusBadRequest)
		return
	}
	fmt.Println(receipt)

	id := uuid.New().String()
	points := calculatePoints(receipt)
	receipts[id] = points

	response := map[string]string{"id": id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getReceiptPoints(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	log.Println("GET /receipts/%s/points endpoint hit\n", id)

	points, exists := receipts[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := map[string]int{"points": points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func calculatePoints(receipt Receipt) int {
	points := 0

	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(re.FindAllString(receipt.Retailer, -1))

	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Floor(total) {
		points += 50
	}

	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	timeOfPurchase, _ := time.Parse("15:04", receipt.PurchaseTime)
	hour := timeOfPurchase.Hour()
	if hour == 14 || hour == 15 {
		points += 10
	}

	return points
}

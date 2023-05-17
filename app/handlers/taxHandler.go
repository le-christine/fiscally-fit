package handlers

import (
	"encoding/json"
	"github.com/le-christine/fiscally-fit/app/services"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TaxHandler struct {
	taxService *services.TaxService
}

func NewTaxHandler(taxService *services.TaxService) *TaxHandler {
	return &TaxHandler{taxService}
}

func (h *TaxHandler) CalculateTax(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	incomeStr := params["income"]
	income, err := strconv.ParseFloat(incomeStr, 64)
	if err != nil {
		http.Error(w, "Invalid salary parameter", http.StatusBadRequest)
		return
	}

	yearStr := params["year"]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received request for Calculate Tax, income: %f, year: %d\n", income, year)

	response, err := h.taxService.CalculateTaxes(income, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshaling JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
package handlers

import (
	"encoding/json"
	"fiscally-fit/app/services"
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

	salaryStr := params["income"]
	salary, err := strconv.ParseFloat(salaryStr, 64)
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

	taxes, err := h.taxService.CalculateTaxes(salary, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"taxes": taxes,
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
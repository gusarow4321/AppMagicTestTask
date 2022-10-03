package service

import (
	"AppMagicTestTask/internal/data"
	"encoding/json"
	"net/http"
)

type Service struct {
	data *data.Results
}

func NewService(data *data.Results) *Service {
	return &Service{data}
}

func (s *Service) GetAllResults(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(s.data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

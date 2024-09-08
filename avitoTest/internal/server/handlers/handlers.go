package handlers

import (
	"avitoTest/avitoTest/internal/errors"
	s "avitoTest/avitoTest/internal/storage"
	"encoding/json"
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, req *http.Request) {
	err := s.Storage.Ping(req.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(errors.ErrorResponse{
			Reason: err.Error(),
		})
		if err != nil {
			log.Println(err.Error())
		}
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}

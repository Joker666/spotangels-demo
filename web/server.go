package web

import (
	"net/http"

	"github.com/Joker666/spotangels-demo/web/controller"
	"github.com/gorilla/mux"
)

// NewServer adds routes to mux
func NewServer(router *mux.Router, regulationController *controller.RegulationController) {
	router.HandleFunc("/regulation/{segment_id}", regulationController.GetActiveRegulation).Methods(http.MethodGet)
}

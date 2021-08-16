package web

import (
	"net/http"

	"github.com/Joker666/spotangels-demo/web/controller"
	"github.com/Joker666/spotangels-demo/web/resp"
	"github.com/gorilla/mux"
)

// NewServer adds routes to mux
func NewServer(router *mux.Router, regulationController *controller.RegulationController) {
	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		resp.ServeData(w, r, http.StatusOK, "Ok")
	})
	router.HandleFunc("/regulation/{segment_id}", regulationController.GetActiveRegulation).Methods(http.MethodGet)
}

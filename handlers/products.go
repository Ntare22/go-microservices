package handlers

import (
	"log"
	"net/http"

	"github.com/Ntare22/go-microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l*log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p* Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	
}
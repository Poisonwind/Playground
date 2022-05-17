package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	log *log.Logger
}

func NewGoodbye(log *log.Logger) *Goodbye {
	return &Goodbye{log}
}

func (gb *Goodbye) GetJustGoodbye(rw http.ResponseWriter, r *http.Request) {

	rw.Write([]byte("Goodbye!!"))

}

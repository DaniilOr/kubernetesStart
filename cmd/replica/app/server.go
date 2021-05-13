package app

import (
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Server struct {
	id string
	mux *http.ServeMux
}

func NewServer(mux *http.ServeMux) *Server {
	return &Server{id: uuid.New().String(), mux: mux}
}

func (s *Server) Init() {
	s.mux.HandleFunc("/", s.getID)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) getID(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
	}
	ans := make([]int64, 12)
	var transactions []Transactions
	err = json.Unmarshal(data, &transactions)
	if err != nil{
		log.Println(err)
	}
	for _, t := range transactions{
		ans[time.Unix(t.Date, 0).Month()] +=  t.Amount
	}
	log.Println(ans)
	resp := Response{Uid: s.id, Stats: ans}
	data, err = json.Marshal(resp)
	if err != nil{
		log.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	// по умолчанию статус 200 Ok
	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
	}
}

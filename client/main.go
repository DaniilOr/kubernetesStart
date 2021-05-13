package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)
type Transactions struct {
	Date int64 `json:"date"`
	Amount int64 `json:"amount"`
}
type Response struct {
	Uid string `json:"uid"`
	Stats []int64 `json:"stats"`
}
func send() {


}
func main(){
	services := make(map[string]int64)
	var wg sync.WaitGroup
	for i:=0; i <=1000; i++{
		wg.Add(1)
		go func(s map[string]int64){
			defer wg.Done()
			buf := []Transactions{{Date: 1620919954, Amount:  10000}, {Date: 1620970000, Amount:  10000}}
			data, err := json.Marshal(&buf)
			if err != nil{
				log.Fatalln(err)
			}
			resp, err := http.Post("http://10.110.182.76:9999/", "application/json",  bytes.NewBuffer(data))
			if err != nil{
				log.Fatalln(err)
			}
			var r Response
			d, err := ioutil.ReadAll(resp.Body)
			if err != nil{
				log.Fatalln(err)
			}
			err = json.Unmarshal(d, &r)
			if err != nil{
				log.Fatalln(err)
			}
			s[r.Uid] += 1
		}(services)
		wg.Wait()
	}
	log.Println(services)
}



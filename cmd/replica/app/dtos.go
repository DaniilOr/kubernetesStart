package app

type Transactions struct {
	Date int64 `json:"date"`
	Amount int64 `json:"amount"`
}
type Response struct {
	Uid string `json:"uid"`
	Stats []int64 `json:"stats"`
}
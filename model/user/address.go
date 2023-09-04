package user

import "time"

type Address struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Phone   string    `json:"phone"`
	AddTime time.Time `json:"add_time"`
}

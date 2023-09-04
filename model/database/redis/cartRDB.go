package redisdao

import (
	"PSHOP/model/user"
	"encoding/json"
)

var (
	prefix = "cart"
)

func MarshalBinary(c *user.Cart) ([]byte, error) {
	return json.Marshal(c)
}

func UnMarshalBinary(c *user.Cart) ([]byte, error) {
	return json.Marshal(c)
}

func KeyCart(useridentity string) string {
	return prefix + "-" + useridentity
}

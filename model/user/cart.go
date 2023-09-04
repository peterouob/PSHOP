package user

type Cart struct {
	UserIdentity string `json:"useridentity"`
	GoodName     string `json:"good_name"`
	GoodID       string `json:"good_id"`
	Nums         int    `json:"num"`
}

func (c *Cart) Cart() string {
	return "cart"
}

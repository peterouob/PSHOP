package user

type GoodInfo struct {
	Id      int        `json:"id"`
	GoodId  string     `json:"good_id"`
	Name    string     `json:"name"`
	Price   int        `json:"price"`
	Info    string     `json:"info"`
	Desc    string     `json:"desc"`
	Pic     string     `json:"pic"`
	Comment []*Comment `json:"comment" gorm:"foreignKey:good_comment_id"`
}

func (g *GoodInfo) TableName() string {
	return "goodinfo"
}

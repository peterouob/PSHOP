package user

type Goods struct {
	Id     int    `json:"id" gorm:"type:int(11);autoIncrement;primaryKey;column:id;"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	GoodId string `json:"good_id"`
	//對應外鍵要查詢的,值不變
	GoodsClass string `json:"goodsClass" gorm:"type:varchar(191)"`
}

func (g *Goods) TableName() string {
	return "goods"
}

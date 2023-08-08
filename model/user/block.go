package user

// 一個分類對應多個物品->1對多
type Block struct {
	ClassName string   `json:"class_name" gorm:"type:varchar(191);primaryKey"`
	Class     []*Goods `json:"class" gorm:"foreignKey:GoodsClass"`
}

func (b *Block) TableName() string {
	return "block"
}

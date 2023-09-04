package cart

import (
	redisdao "PSHOP/model/database/redis"
	"PSHOP/model/user"
	"context"
	"fmt"
)

func Add(cartModel user.Cart) error {
	var cart *user.Cart
	//TODO:解決redis重複插入值被覆蓋問題
	if cart == nil {
		cart = &user.Cart{
			UserIdentity: cartModel.UserIdentity,
			GoodID:       cartModel.GoodID,
			Nums:         cartModel.Nums,
			GoodName:     cartModel.GoodName,
		}
	}
	cart.Nums += cartModel.Nums
	c, err := redisdao.MarshalBinary(cart)
	if err != nil {
		fmt.Println("MarshalBinary error", err.Error())
	}
	set := redisdao.Rdb.Set(context.Background(), redisdao.KeyCart(cartModel.UserIdentity), c, 0)
	fmt.Println(map[string]interface{}{
		"user_cart_set": set,
	})
	redisdao.Rdb.SAdd(context.Background(), redisdao.KeyCart(cartModel.UserIdentity), c)
	return nil
}

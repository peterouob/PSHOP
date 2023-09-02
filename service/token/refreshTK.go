package serviceToken

//func RefreshToken(c *gin.Context) {
//	mapToken := map[string]string{}
//	if err := c.ShouldBind(&mapToken); err != nil {
//		H.Fail(c, err.Error())
//	}
//	refreshToken := mapToken["refresh_token"]
//	ref := utils.Config.GetString("token.refreshval")
//	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(ref), nil
//	})
//	if err != nil {
//		H.Fail(c, "Refresh token expired")
//	}
//	// is token Valid ?
//	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
//		H.Fail(c, err.Error())
//	}
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if ok && token.Valid {
//		refersUuid, ok := claims["refresh_uuid"].(string)
//		if !ok {
//			H.Fail(c, "Refresh token expired")
//		}
//		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
//		if err != nil {
//			H.Fail(c, "Refresh token expired")
//		}
//		//	Del the pre token
//		del, delErr := redisdao.DeleteTokenAuth(refersUuid)
//		if del == 0 || delErr != nil {
//			H.Fail(c, "unauthorized")
//		}
//		//create new token
//		ts, createErr := utils.CreateToken(userId)
//		if createErr != nil {
//			H.Fail(c, createErr.Error())
//		}
//		saveErr := redisdao.SaveTokenAuth(userId, ts)
//		if saveErr != nil {
//			H.Fail(c, saveErr.Error())
//		}
//		tokens := map[string]string{"access_token": ts.AccessToken, "refresh_token": ts.RefreshToken}
//		H.OK(c, tokens)
//	} else {
//		H.Fail(c, "token error !"+err.Error())
//	}
//}

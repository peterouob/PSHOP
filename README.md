- [x] 使用者登入註冊
- [x] 手寫session
- [x] 完善session功能
- [ ] sudo權限
- [ ] 完成token部分
- [ ] 商品主頁面(依照種類分類)
- [ ] 黑名單系統
- [ ] 購物車系統
- [ ] 模擬交易過程

### session 問題
- 使用GET請求時可以正常操作(set & get)
- POST請求時只能set無法get -> null
- 解決方法:
  - 將session直接寫入RDB檢查GET RDB的KEY
  - 使用token取代session進行身份驗證
- 成功解決:
  - 將底層重構,使用自行編寫的setCookie取代原先的r.Cookie
- Redis中因為session的名稱和cookie名稱不同導致無法刪除
  - 使用chan來傳遞redis的name

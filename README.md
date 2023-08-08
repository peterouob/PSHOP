- [x] 使用者登入註冊
- [x] 手寫session
- [x] 完善session功能
- [ ] sudo權限
- [x] 完成token部分
- [x] 商品主頁面(依照種類分類)
- [ ] 黑名單系統
- [ ] 購物車系統
- [ ] 商品資訊和留言 
- [ ] 簡易秒殺系統
- [ ] 令牌桶流量限制
- [ ] 模擬交易過程

### session 問題
- 使用GET請求時可以正常操作(set & get)POST則無法
  - 解決:
    - 將底層重構,使用自行編寫的setCookie取代原先的r.Cookie

- Redis中因為session的名稱和cookie名稱不同導致無法刪除
  - 解決:
    - 使用chan來傳遞redis的name

### Mysql 問題
- 一開始將主鍵設置在會auto inc的id上導致外鍵類型錯誤
  - 解決:
    - 將主鍵設置的column重新設計

- 帶有外鍵的column無法設置
  - 解決:
    - 先將主鍵設值

- 使用preload查詢where查詢依舊查到不希望存在的資料
  - 解決:
    - 在preload前再where查詢一次 
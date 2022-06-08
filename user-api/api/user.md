### 1. "获取用户信息"

1. route definition

- Url: /user/info
- Method: POST
- Request: `UserInfoRequest`
- Response: `UserInfoResponse`

2. request definition



```golang
type UserInfoRequest struct {
	UserId int64 `json:"userId"`
}
```


3. response definition



```golang
type UserInfoResponse struct {
	UserId int64 `json:"userId"`
	Nickname string `json:"nickname"`
}
```


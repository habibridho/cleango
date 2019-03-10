package rest

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

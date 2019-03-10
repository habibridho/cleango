package rest

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Movie struct {
	ID       uint64 `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
}

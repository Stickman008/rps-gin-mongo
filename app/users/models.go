package users

type Game struct {
	Enemyname string `json:"username" bson:"username"`
}

type Challenger struct {
	Enemyname string `json:"username" bson:"username"`
	rps       string `json:"rps"`
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	// Winning  int    `json:"nWinning"`
	// Challenge string `json:"username" bson:"cha"`
	// Status string `json:"status" bson:"status"`
}

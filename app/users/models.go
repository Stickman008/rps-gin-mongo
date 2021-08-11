package users

type Game struct {
	Enemyname string `json:"username" bson:"username"`
}

type Challenger struct {
	Enemyname string `json:"username" bson:"username"`
	Rps       string `json:"rps"`
}

type BattleLog struct {
	Win  string `json:"win"`
	Lose string `json:"lose"`
	Draw string `json:"draw"`
}

type User struct {
	Id       string    `json:"id"`
	Username string    `json:"username"`
	History  BattleLog `json:"history`
	// Challenge string `json:"username" bson:"cha"`
	// Status string `json:"status" bson:"status"`
}

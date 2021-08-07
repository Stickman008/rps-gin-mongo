package users

type User struct {
	Id     string `json:"id" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age    int    `json:"age" bson:"age"`
}

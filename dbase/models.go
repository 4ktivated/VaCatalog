package dbase

import "time"


type Vac struct {
	ID      string `bson:"_id"`
	Lang    string `bson:"lang"`
	Title   string `bson:"Title"`
	Company string `bson:"Company"`
	URL     string `bson:"URL"` //change on url
	Salary  string `bson:"Salary"`
	Info    string `bson:"Info"`
	ExpireAt *time.Time `bson:"expireAt,omitempty"`
}

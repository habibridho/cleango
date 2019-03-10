package data

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/habibridho/cleango/entities"
)

type UserDataImpl struct {
	MgSession mgo.Session
}

type User struct {
	ID       int32  `bson:"id"`
	Username string `bson:"username"`
	Fullname string `bson:"fullname"`
	Password string `bson:"password"`
}

func (impl UserDataImpl) GetUser(username string) (user entities.User, err error) {
	query := bson.M{
		"username": username,
	}

	result := User{}
	if err = impl.MgSession.DB("app").C("user").Find(query).One(&result); err != nil {
		return
	}

	user = entities.User{
		ID:       uint64(result.ID),
		Username: result.Username,
		Fullname: result.Fullname,
		Password: result.Password,
	}

	return
}

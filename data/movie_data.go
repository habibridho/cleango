package data

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/habibridho/cleango/entities"
	"github.com/kataras/go-errors"
)

type MovieDataImpl struct {
	MgSession mgo.Session
}

type Movie struct {
	ID       int32  `bson:"id"`
	Title    string `bson:"title"`
	Director string `bson:"director"`
}

func (impl MovieDataImpl) GetMovieList() (data []entities.Movie, err error) {
	movies := []Movie{}
	if err = impl.MgSession.DB("app").C("movie").Find(bson.M{}).All(&movies); err != nil {
		return
	}

	data = make([]entities.Movie, len(movies))
	for i, m := range movies {
		data[i] = entities.Movie{
			ID:       uint64(m.ID),
			Title:    m.Title,
			Director: m.Director,
		}
	}

	return
}

func (impl MovieDataImpl) LikeDislike(userID uint64, movieID uint64, status bool) (err error) {
	return errors.New("please implement")
}

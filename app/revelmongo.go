package revelmongo

import (
	"errors"
	"fmt"
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	Session *mgo.Session
	Url     string
)

func Init() {

	var found bool
	if Url, found = revel.Config.String("mongo.url"); !found {
		revel.ERROR.Fatal("No mongo.url found.")
	}

	var err error
	if Session, err = mgo.Dial(Url); err != nil {
		revel.ERROR.Panic(err)
	}

}

const (
	New   = 0
	Copy  = 1
	Clone = 1
)

type Session struct {
	*revel.Controller
	Session *mgo.Session
	Method  int
}

func SetMethod(method int) {

}

func (c *Session) Begin() {

}

func (c *Session) End() {

}

func init() {
	revel.InterceptMethod((*Session).Begin, revel.BEFORE)
	revel.InterceptMethod((*Session).End, revel.FINALLY)
}

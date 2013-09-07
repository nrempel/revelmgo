package revelmgo

import (
	"fmt"
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
)

var (
	Session *mgo.Session
	Url     string
	Method  string
)

func Init() {
	var found bool
	if Url, found = revel.Config.String("mongo.url"); !found {
		revel.ERROR.Fatal("No mongo.url found.")
	}

	if Method, found = revel.Config.String("mongo.method"); !found {
		revel.Error.Fatal("No mongo.method found")
	}

	var err error
	if Session, err = mgo.Dial(Url); err != nil {
		revel.ERROR.Panic(err)
	}
}

type Controller struct {
	*revel.Controller
	MongoSession *mgo.Session
}

func (c *Controller) Begin() {
	switch Method {
	case "new":
		MongoSession = Session.New()
	case "copy":
		MongoSession = Session.Copy()
	case "clone":
		MongoSession = Session.Clone()
	default:
		revel.Error.Fatal(fmt.Sprintf(
			"Invalid mongo.method: %s.\nUse new, copy, or clone.",
			Method))
	}
}

func (c *Controller) End() {
	Session.Close()
	c.MongoSession.Close()
}

func init() {
	revel.InterceptMethod((*Controller).Begin, revel.BEFORE)
	revel.InterceptMethod((*Controller).End, revel.FINALLY)
}

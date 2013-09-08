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
	if Url, found = revel.Config.String("mgo.url"); !found {
		revel.ERROR.Fatal("No mgo.url found")
	}

	if Method, found = revel.Config.String("mgo.method"); !found {
		revel.ERROR.Fatal("No mgo.method found")
	}

	var err error
	if Session, err = mgo.Dial(Url); err != nil {
		revel.ERROR.Panic(err)
	}
}

type MgoController struct {
	*revel.Controller
	MgoSession *mgo.Session
}

func (c *MgoController) Begin() revel.Result {
	switch Method {
	case "new":
		c.MgoSession = Session.New()
	case "copy":
		c.MgoSession = Session.Copy()
	case "clone":
		c.MgoSession = Session.Clone()
	default:
		revel.ERROR.Fatal(fmt.Sprintf(
			"Invalid mgo.method: %s.\nUse new, copy, or clone.",
			Method))
	}
	return nil
}

func (c *MgoController) End() revel.Result {
	Session.Close()
	c.MgoSession.Close()
	return nil
}

func init() {
	revel.InterceptMethod((*MgoController).Begin, revel.BEFORE)
	revel.InterceptMethod((*MgoController).End, revel.FINALLY)
}

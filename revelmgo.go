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

func (c *MgoController) New() {
	c.MgoSession = Session.New()
}

func (c *MgoController) Copy() {
	c.MgoSession = Session.Copy()
}

func (c *MgoController) Clone() {
	c.MgoSession = Session.Clone()
}

func (c *MgoController) End() revel.Result {
	Session.Close()
	c.MgoSession.Close()
	return nil
}

func init() {
	revel.InterceptMethod((*MgoController).End, revel.FINALLY)
}

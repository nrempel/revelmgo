package revelmgo

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
)

var (
	Session *mgo.Session
	Url     string
)

func Init() {
	var found bool
	if Url, found = revel.Config.String("mgo.url"); !found {
		revel.ERROR.Fatal("No mgo.url found")
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

func New() {
	revel.InterceptMethod((*MgoController).new, revel.BEFORE)
}

func Copy() {
	revel.InterceptMethod((*MgoController).copy, revel.BEFORE)
}

func Clone() {
	revel.InterceptMethod((*MgoController).clone, revel.BEFORE)
}

func (c *MgoController) new() revel.Result {
	c.MgoSession = Session.New()
	return nil
}

func (c *MgoController) copy() revel.Result {
	c.MgoSession = Session.Copy()
	return nil
}

func (c *MgoController) clone() revel.Result {
	c.MgoSession = Session.Clone()
	return nil
}

func (c *MgoController) close() revel.Result {
	c.MgoSession.Close()
	return nil
}

func init() {
	revel.InterceptMethod((*MgoController).close, revel.FINALLY)
}

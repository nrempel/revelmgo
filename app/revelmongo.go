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
	Method  string
	Url     string
)

func Init() {
	var found bool

}

type Session struct {
	*revel.Controller
	Session *mgo.Session
}

func (c *Session) Begin() {

}

func (c *Session) End() {

}

func init() {
	revel.InterceptMethod((*Session).Begin, revel.BEFORE)
	revel.InterceptMethod((*Session).End, revel.FINALLY)
}

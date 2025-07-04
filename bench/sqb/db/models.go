package db

import (
	"github.com/blink-io/sq"
)

type MODELS struct {
	sq.TableStruct
	ID      sq.NumberField
	NAME    sq.StringField
	TITLE   sq.StringField
	FAX     sq.StringField
	WEB     sq.StringField
	AGE     sq.NumberField
	RIGHT   sq.BooleanField
	COUNTER sq.NumberField
}

var ModelsTable = sq.New[MODELS]("mm")

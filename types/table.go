package types

import "gopkg.in/mgo.v2/bson"

type Tables struct {
	Hash map[string]interface{}
}

func (tab Tables) ToBSON() ([]byte, error) {
	return bson.Marshal(tab)
}

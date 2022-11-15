package lib

import (
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

func InitBase(name string) (*base.Base, error) {
	db, err := deta.New()
	if err != nil {
		return nil, err
	}

	return base.New(db, name)
}

func GetCollection(base *base.Base, id string) (col *CollectionProps, err error) {
	err = base.Get(id, col)
	return
}

package global

import (
	"titd/modules/datastore"
	"titd/settings"
)

var (
	DataStore *datastore.DataStore
)

func Init() {

	var err error

	if DataStore, err = datastore.NewStore(settings.DbPath); err != nil {
		panic(err)
	}

}

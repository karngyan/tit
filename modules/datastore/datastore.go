package datastore

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/dgraph-io/badger/v3"
	"github.com/dgraph-io/badger/v3/options"
)

var (
	BucketSeparator = "/"
)

type DataStore struct {
	path   string
	lastGC int64
	DB     *badger.DB
}

func NewStore(path string) (*DataStore, error) {
	var err error
	s := &DataStore{
		path: path,
	}
	badgerOptions := badger.DefaultOptions(path)
	badgerOptions = badgerOptions.WithSyncWrites(false) // default is true
	badgerOptions = badgerOptions.WithCompression(options.ZSTD)
	if s.DB, err = badger.Open(badgerOptions); err == nil {
		return s, nil
	}
	logs.Error(err)
	return s, err
}

func (s *DataStore) Close() error {
	return s.DB.Close()
}

func (s *DataStore) Put(bucket string, key string, val []byte) error {

	var err error
	key = bucket + BucketSeparator + key

	return s.DB.Update(func(tx *badger.Txn) error {
		e := badger.NewEntry([]byte(key), val)
		if err = tx.SetEntry(e); err != nil {
			logs.Error("error while inserting in bucket")
			logs.Error(err)
			return err
		}
		return nil
	})
}

func (s *DataStore) Delete(bucket string, key string) error {

	var err error
	key = bucket + BucketSeparator + key

	return s.DB.Update(func(tx *badger.Txn) error {
		e := []byte(key)
		if err = tx.Delete(e); err != nil {
			logs.Error("error while deleting from bucket")
			logs.Error(err)
			return err
		}
		return nil
	})
}

func (s *DataStore) GetAll(bucket string, vals *[][]byte) error {
	var err error

	err = s.DB.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		prefix := []byte(bucket)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				*vals = append(*vals, val)
				return nil
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		logs.Error(err)
	}

	return err
}

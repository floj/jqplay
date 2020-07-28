package boltstore

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/jingweno/jqplay/jq"
)

const BucketName = "snippets"

type BoltDbStore struct {
	db     *bolt.DB
	bucket []byte
}

func NewBoltDbStore(db *bolt.DB) (*BoltDbStore, error) {
	s := &BoltDbStore{db: db, bucket: []byte(BucketName)}
	err := s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(s.bucket)
		return err
	})
	return s, err
}

func (s *BoltDbStore) Get(slug string) (*jq.JQ, error) {
	d := &jq.JQ{}
	err := s.read(func(b *bolt.Bucket) error {
		v := b.Get([]byte(slug))
		if v == nil {
			return nil
		}
		return json.Unmarshal(v, d)
	})
	return d, err
}

func (s *BoltDbStore) Put(snip *jq.JQ) (string, error) {
	v, err := json.Marshal(snip)
	if err != nil {
		return "", err
	}
	id := slug(v)
	err = s.write(func(b *bolt.Bucket) error {
		return b.Put([]byte(id), v)
	})
	return id, err
}

func (s *BoltDbStore) read(f func(b *bolt.Bucket) error) error {
	return s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)
		return f(b)
	})
}

func (s *BoltDbStore) write(f func(b *bolt.Bucket) error) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)
		return f(b)
	})
}

func slug(s []byte) string {
	h := sha1.New()
	sum := h.Sum(s)
	b := make([]byte, base64.URLEncoding.EncodedLen(len(sum)))
	base64.URLEncoding.Encode(b, sum)
	return string(b)[:10]
}

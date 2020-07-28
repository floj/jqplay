package storage

import (
	"github.com/jingweno/jqplay/jq"
)

type Store interface {
	Get(slug string) (*jq.JQ, error)
	Put(snip *jq.JQ) (string, error)
}

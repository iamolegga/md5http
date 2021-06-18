package md5hasher

import (
	"crypto/md5"
	"fmt"
)

type Hasher struct{}

func New() *Hasher {
	return &Hasher{}
}

func (h *Hasher) Hash(bytes []byte) string {
	return fmt.Sprintf("%x", md5.Sum(bytes))
}

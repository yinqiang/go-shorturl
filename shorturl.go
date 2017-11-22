package shorturl

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
)

var (
	Keys = []byte("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func Short(url string) []string {
	hash := md5.Sum([]byte(url))
	var i0, i1, i2, i3 uint32
	b := bytes.NewBuffer(hash[0:16])
	binary.Read(b, binary.BigEndian, &i0)
	binary.Read(b, binary.BigEndian, &i1)
	binary.Read(b, binary.BigEndian, &i2)
	binary.Read(b, binary.BigEndian, &i3)
	i0 &= 0x3FFFFFFF
	i1 &= 0x3FFFFFFF
	i2 &= 0x3FFFFFFF
	i3 &= 0x3FFFFFFF
	k0, k1, k2, k3 := toKey(i0), toKey(i1), toKey(i2), toKey(i3)
	return []string{k0, k1, k2, k3}
}

func toKey(n uint32) string {
	b := make([]byte, 6)
	var i uint
	for i = 0; i < 6; i++ {
		j := n & 0x0000003D
		b[i] = Keys[j]
		n = n >> 5
	}
	return string(b)
}

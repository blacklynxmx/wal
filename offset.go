package wal

import (
	"encoding/binary"
)

// Offset records an offset in the WAL
type Offset []byte

func newOffset(fileSequence int64, position int64) Offset {
	o := make(Offset, 16)
	binary.BigEndian.PutUint64(o, uint64(fileSequence))
	binary.BigEndian.PutUint64(o[8:], uint64(position))
	return o
}

func (o Offset) FileSequence() int64 {
	return int64(binary.BigEndian.Uint64(o))
}

func (o Offset) Position() int64 {
	return int64(binary.BigEndian.Uint64(o[8:]))
}

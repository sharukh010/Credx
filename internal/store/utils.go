package store

import "time"

type hasID interface {
	SetID(int64)
}

type hasCreatedAt interface{
	SetCreatedAt(time.Time)
}

type hasUpdatedAt interface {
	SetUpdatedAt(time.Time)
}

func setID(data hasID, ID *int64) {
	*ID += 1
	data.SetID(*ID)
}

func setCreatedAt(data hasCreatedAt) {
	data.SetCreatedAt(time.Now())
}
func setUpdatedAt(data hasUpdatedAt) {
	data.SetUpdatedAt(time.Now())
}
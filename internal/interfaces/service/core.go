package service

import "bytes"

type Core interface {
	Check(filename string, image *bytes.Buffer, parkingZoneID int64) (err error)
}

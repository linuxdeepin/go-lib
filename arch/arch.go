package arch

import (
	"runtime"
)

type ArchFamilyType int

const (
	Unknown ArchFamilyType = iota
	AMD64
	Sunway
)

func Get() ArchFamilyType {
	switch runtime.GOARCH {
	case "sw_64":
		return Sunway
	case "amd64":
		return AMD64
	default:
		return Unknown
	}
}

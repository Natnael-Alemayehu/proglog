package log

type Config struct {
	Segment struct {
		MaxStoreByte  uint64
		MaxIndexByte  uint64
		InitialOffder uint64
	}
}

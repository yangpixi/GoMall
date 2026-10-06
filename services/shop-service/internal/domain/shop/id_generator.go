package shop

type IDGenerator interface {
	NextID() int64
}

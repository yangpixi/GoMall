package shared

type IDGenerator interface {
	NextID() int64
}

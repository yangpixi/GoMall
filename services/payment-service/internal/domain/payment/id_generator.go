package payment

type IDGenerator interface {
	NextID() int64
}

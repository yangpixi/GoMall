package address

// IDGenerator generate snowflake id for entity
type IDGenerator interface {
	NextID() int64
}

package shared

type IDGenerator interface {
	// NextID generate a new snowflake id
	NextID() int64
}

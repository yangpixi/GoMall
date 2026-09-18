package permission

type Permission struct {
	ID       int64
	code     string
	resource string
	action   string
	name     string
}

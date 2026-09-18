package role

type Role struct {
	id            int64
	code          string
	name          string
	permissionIDs []int64
}

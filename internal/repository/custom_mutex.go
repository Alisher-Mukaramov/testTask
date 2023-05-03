package repository

func (r *repository) Lock() {
	r.locked = true
}
func (r *repository) UnLock() {
	r.locked = false
}
func (r *repository) IsLocked() bool {
	return r.locked
}

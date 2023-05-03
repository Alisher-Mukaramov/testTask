package service

import "fmt"

func (s *service) State() (int, error) {
	if s.repo.IsLocked() {
		return 0, fmt.Errorf("%w", ErrorProcess)
	}
	return s.repo.State()
}

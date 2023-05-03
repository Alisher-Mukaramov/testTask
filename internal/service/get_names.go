package service

import "interviewTask/internal/models"

func (s *service) GetNames(names []string, _type string) ([]models.Name, error) {
	return s.repo.GetNames(names, _type)
}

package service

import (
	"encoding/json"
	"interviewTask/internal/models"
	"interviewTask/internal/pkg/utils"
)

const sdnType = "Individual"

func (s *service) Update(list models.SdnList) (err error) {
	se := utils.SliceFilter(list.SdnEntry, func(se models.SdnEntry) bool {
		if se.SdnType != sdnType {
			return false
		}
		return true
	})

	list.SdnEntry = se

	bytes, err := json.Marshal(list)
	if err != nil {
		return
	}

	return s.repo.Update(string(bytes))
}

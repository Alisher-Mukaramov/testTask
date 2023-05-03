package repository

import (
	"fmt"
	"interviewTask/internal/models"
)

const (
	type1 = "strong"
	type2 = "weak"
)

func (r *repository) GetNames(n []string, t string) (names []models.Name, err error) {
	query := `select uid, trim(first_name) as first_name, trim(last_name) as last_name 
			  from public.list where 1=1`

	switch {
	case len(n) > 1 && t == type1:
		query += fmt.Sprintf(` and first_name = '%s' and last_name = '%s'`, n[0], n[1])
	case len(n) > 1 && (t == type2 || t == ""):
		query += fmt.Sprintf(` and (first_name like '%%%s%%' or first_name like '%%%s%%' or last_name like '%%%s%%' or last_name like '%%%s%%')`, n[0], n[1], n[0], n[1])
	case len(n) == 1 && t == type1:
		query += fmt.Sprintf(` and (first_name = '%s' or last_name = '%s')`, n[0], n[0])
	case len(n) == 1 && t == type2:
		query += fmt.Sprintf(` and (first_name like '%%%s%%' or last_name like '%%%s%%')`, n[0], n[0])
	}

	err = r.pg.GetInstance().Select(&names, query)

	return
}

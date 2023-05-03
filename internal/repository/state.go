package repository

const stateFunc = "select public.state_f()"

func (r *repository) State() (count int, err error) {
	rows, err := r.pg.GetInstance().Queryx(stateFunc)
	if err != nil {
		return
	}

	for rows.Next() {
		var c int
		err = rows.Scan(&c)
		if err != nil {
			return
		}
		count += c
	}
	return
}

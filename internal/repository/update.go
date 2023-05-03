package repository

const updateFunc = "select public.update_f($1)"

func (r *repository) Update(XMLString string) error {
	r.Lock()
	defer r.UnLock()
	_, err := r.pg.GetInstance().Queryx(updateFunc, XMLString)
	return err
}

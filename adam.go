package adam

type Adam struct {
	cases map[string]*Cas
}

func (a *Adam) AddCases(name string, c *Cas) error {
	a.cases[name] = c
	return nil
}

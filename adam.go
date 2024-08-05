package adam

type Adam struct {
	cases map[string]*Cas
}

func (a *Adam) AddCase(name string, c *Cas) error {
	a.cases[name] = c
	return nil
}

func (a *Adam) DeleteCase(name string) error {
	return nil
}

func (a *Adam) UpdateCase(name string, c *Cas) error {
	return nil
}

func (a *Adam) GetCase(name string) (*Cas, error) {
	return nil, nil
}

package models

type ResourceA struct {
	Kind      string
	Name      string
	Namespace string
}

func (r ResourceA) GetName() string {
	return r.Name
}

type ResourceB struct {
	Kind      string
	Name      string
	Namespace string
}

func (r ResourceB) GetName() string {
	return r.Name
}

type ResourceC struct {
	Kind      string
	Name      string
	Namespace string
}

func (r ResourceC) GetName() string {
	return r.Name
}

type CN2Resource interface {
	ResourceA | ResourceB | ResourceC
	GetName() string
}

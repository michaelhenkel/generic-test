package ext

import (
	"github.com/michaelhenkel/generic-test/converter"
	"github.com/michaelhenkel/generic-test/models"
)

type ExtReconciler[R models.CN2Resource] struct {
	cn2Converter converter.ConvertToCN2
}

func New() *ExtReconciler[models.ResourceA] {
	return &ExtReconciler[models.ResourceA]{}
}

func (e *ExtReconciler[R]) AddCN2Converter(cn2Converter converter.ConvertToCN2) {
	e.cn2Converter = cn2Converter
}

type ExtResourceA struct {
	Bar string
	Foo string
}

func (e *ExtReconciler[R]) Convert(resource R) {
	//resource.GetName()
}

package converter

import (
	"github.com/michaelhenkel/generic-test/models"
)

type ConvertToExt[R models.CN2Resource] interface {
	Convert(R)
}

type ConvertToCN2 interface {
	Add()
	Get()
	Delete()
	Convert()
}

func ToExt[C any](c C) {

}

func ToCN2() {

}

func To[C models.CN2Resource]() {
}

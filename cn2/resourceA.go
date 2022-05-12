package cn2

import (
	"fmt"

	"github.com/michaelhenkel/generic-test/converter"
	"github.com/michaelhenkel/generic-test/models"
	contrail "ssd-git.juniper.net/contrail/cn2/contrail/pkg/apis/core/v1alpha1"
)

type ResourceAReconciler[R models.CN2Resource] struct {
	externalConverter converter.ConvertToExt[R]
}

func NewResourceAReconciler() *ResourceAReconciler[models.ResourceA] {
	vn := contrail.VirtualNetwork{}
	fmt.Println(vn)
	var reconciler ResourceAReconciler[models.ResourceA]
	return &reconciler
}

func (r *ResourceAReconciler[R]) AddExternalConverter(externalConverter converter.ConvertToExt[R]) {
	r.externalConverter = externalConverter
}

func (r *ResourceAReconciler[R]) Convert() {

}

func (r *ResourceAReconciler[R]) Add() {

}

func (r *ResourceAReconciler[R]) Get() {

}

func (r *ResourceAReconciler[R]) Delete() {

}

func (r *ResourceAReconciler[R]) Reconcile(resource R) {

	r.externalConverter.Convert(resource)

}

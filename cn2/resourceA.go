package cn2

import (
	"github.com/michaelhenkel/generic-test/converter"
	"github.com/michaelhenkel/generic-test/models"
	contrail "ssd-git.juniper.net/contrail/cn2"
)

type ResourceAReconciler[R models.CN2Resource] struct {
	externalConverter converter.ConvertToExt[R]
}

func NewResourceAReconciler() *ResourceAReconciler[models.ResourceA] {
	contrail.VirtualNetwork{}
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

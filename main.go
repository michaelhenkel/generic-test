package main

import (
	"github.com/michaelhenkel/generic-test/cn2"
	"github.com/michaelhenkel/generic-test/ext"
	"github.com/michaelhenkel/generic-test/models"
)

func main() {
	cn2Reconciler := cn2.NewResourceAReconciler()
	extReconciler := ext.New()

	cn2Reconciler.AddExternalConverter(extReconciler)
	extReconciler.AddCN2Converter(cn2Reconciler)

	cn2Reconciler.Reconcile(models.ResourceA{
		Kind: "ResourceA",
		Name: "FOOBAR",
	},
	)

}

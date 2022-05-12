module github.com/michaelhenkel/generic-test

go 1.18

replace (
	github.com/google/gnostic => github.com/googleapis/gnostic v0.5.5
	github.com/googleapis/gnostic => github.com/google/gnostic v0.5.7-v3refs
	google.golang.org/grpc => google.golang.org/grpc v1.32.0
	ssd-git.juniper.net/contrail/cn2 => ../../../ssd-git.juniper.net/contrail/cn2
)

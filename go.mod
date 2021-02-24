module github.com/maximilien/kn-source-pkg

go 1.15

require (
	github.com/maxbrunsfeld/counterfeiter/v6 v6.3.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/api v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/client v0.21.0
	knative.dev/pkg v0.0.0-20210216013737-584933f8280b
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3

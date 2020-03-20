module github.com/maximilien/kn-source-pkg

require (
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	knative.dev/client v0.13.0
	knative.dev/test-infra v0.0.0-20200229011351-4dac123b9a3d
)

// Temporary pinning certain libraries. Please check periodically, whether these are still needed
// ----------------------------------------------------------------------------------------------

// Fix for `[` in help messages and shell completion code
// See https://github.com/spf13/cobra/pull/899
replace github.com/spf13/cobra => github.com/chmouel/cobra v0.0.0-20191021105835-a78788917390

go 1.13

package action

import (
	"github.com/danusarf/glide/godep"
	"github.com/danusarf/glide/msg"
)

// ImportGodep imports a Godep file.
func ImportGodep(dest string) {
	base := "."
	config := EnsureConfig()
	if !godep.Has(base) {
		msg.Die("No Godep data found.")
	}
	deps, err := godep.Parse(base)
	if err != nil {
		msg.Die("Failed to extract Godeps file: %s", err)
	}
	appendImports(deps, config)
	writeConfigToFileOrStdout(config, dest)
}

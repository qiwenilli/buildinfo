package gomod

import (
	"fmt"
	"runtime/debug"
)

func BuildInfos() []string {

	var infos []string
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			infos = append(infos, recursionReplace(dep))
		}
	}
	return infos
}

func recursionReplace(mod *debug.Module) string {
	if mod.Replace == nil {
		return fmt.Sprintf("%s@%s", mod.Path, mod.Version)
	}
	return fmt.Sprintf("%s@%s => %s", mod.Path, mod.Version, recursionReplace(mod.Replace))
}

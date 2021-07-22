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
	ver := mod.Version
	if ver == "" {
		ver = "Local"
	}
	if mod.Replace == nil {
		return fmt.Sprintf("%s@%s", mod.Path, ver)
	}
	return fmt.Sprintf("%s@%s => %s", mod.Path, ver, recursionReplace(mod.Replace))
}

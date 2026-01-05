package systeminfo

import "runtime"

type Info struct {
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	CPUs      int    `json:"cpus"`
	GOVersion string `json:"go-version"`
}

func Collect() Info {
	return Info{
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		CPUs:      runtime.NumCPU(),
		GOVersion: runtime.Version(),
	}
}

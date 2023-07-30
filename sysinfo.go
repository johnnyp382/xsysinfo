package sysinfo

import "sync"

type cpuInfo struct {
	once sync.Once
	name string
}

var CPU cpuInfo

func (cpu *cpuInfo) Name() string {
	cpu.once.Do(func() {
		// Try to get the information from internal/cpu.
		// if name := internalcpu.Name(); name != "" {
		// 	cpu.name = name
		// 	return
		// }

		// ^^ The above is temporarily disabled while proof of concepting this
		// as a module (since it calls an internal go package), but would still
		// exist in the stdlib version of this.

		// Try to get the information from local system info.
		// TODO(martisch): use /proc/cpuinfo and /sys/devices/system/cpu/ on Linux as fallback.
		if name := sysinfoName(); name != "" {
			cpu.name = name
			return
		}
	})
	return cpu.name
}

// func sysinfoName() string is declared in platform specific files

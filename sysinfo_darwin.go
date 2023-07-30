package sysinfo

import "syscall"

// sysinfoName on Darwin utilizes a built-in macOS syscall
func sysinfoName() string {
	name, err := syscall.Sysctl("machdep.cpu.brand_string")
	if err != nil {
		return ""
	}
	return name
}

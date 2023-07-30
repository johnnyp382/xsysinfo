//go:build !darwin && !linux
// +build !darwin,!linux

package sysinfo

// sysinfoName returns zero value on unsupported systems
func sysinfoName() string {
	return ""
}

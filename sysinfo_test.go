package sysinfo

import (
	"testing"
)

func Test_cpuInfo_Name(t *testing.T) {
	t.Logf("CPU.Name() = %v", CPU.Name())
}

func Test_sysinfoName(t *testing.T) {
	t.Logf("sysinfoName() = %v", sysinfoName())
}

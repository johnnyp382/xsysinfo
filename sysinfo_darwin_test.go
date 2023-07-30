package sysinfo

import "testing"

// no known situations where darwin should fail to return brand string from
// syscall, so make it a failing test if so
func Test_sysinfoName_Darwin(t *testing.T) {
	name := sysinfoName()
	if name == "" {
		t.Fatal("got empty sysinfoName on darwin")
	}
}

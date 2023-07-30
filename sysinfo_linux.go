package sysinfo

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// sysinfoName on Linux attempts to parse /proc/cpuinfo
// TODO(martisch): use /proc/cpuinfo and /sys/devices/system/cpu/ on Linux as fallback.
func sysinfoName() string {
	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return ""
	}
	defer f.Close()

	name, _ := parseProcfsCpuinfo(f)
	return name
}

func parseProcfsCpuinfo(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		if bytes.HasPrefix(line, []byte("model name")) {
			_, v, ok := bytes.Cut(line, []byte(": "))
			if ok {
				return string(v), nil
			}
		}
	}

	return "", scanner.Err()
}

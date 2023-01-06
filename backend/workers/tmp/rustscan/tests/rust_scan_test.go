package tests

import (
	"rustscan/handler"
	"testing"
)

func TestRustScanFile(t *testing.T) {
	// handler.HandlerRustResult()
}

func TestRustScan(t *testing.T) {
	ip := "124.221.215.20"
	handler.RustScan(ip)
}

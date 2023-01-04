package main

import (
	"gobuster/handler"
	"testing"
)

func TestGobusterFile(t *testing.T) {
	handler.Handler_gobuster("/Users/liyang/tools/asm/ASM/backend/workers/gobuster/io.txt", "https://buffered.io")
}

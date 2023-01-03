package main

import (
	"gobuster/handler"
	"testing"
)

func TestGobusterFile(t *testing.T) {
	handler.Handler_gobuster("D:\\code\\asm-demo\\backend\\workers\\gobuster\\b8ff797e-3bfb-4bae-af48-31533304d038.txt", "https://buffered.io")
}

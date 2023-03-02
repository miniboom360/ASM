package activitys

import "testing"

func TestNaabu(t *testing.T) {
	targets := make([]string, 0)
	targets = append(targets, "106.75.13.27", "remote.cloudpnr.com")
	NaabuScan(targets, "top-1000")
}

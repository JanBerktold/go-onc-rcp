package rpc

import "testing"

func TestOverflow(t *testing.T) {
	client, err := DialTCP("localhost:111", Program{10000, 2})
	t.Log(client, err)
	for i := 0; i < 30; i++ {
		x, _ := client.obtainXId()
		t.Log(x)
	}
}

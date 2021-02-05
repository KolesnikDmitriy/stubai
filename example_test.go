package stubai

import "testing"

func TestExamplePass(t *testing.T) {
	t.Log("this test is passed")
}

func TestExampleFailed(t *testing.T) {
	t.Error("this test is passed")
}

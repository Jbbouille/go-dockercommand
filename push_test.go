package dockercommand

import "testing"

func TestDockerPush(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	docker, err := NewDocker("")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	err = docker.Push(&PushOptions{
		Image: "alpine",
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

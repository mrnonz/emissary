package kubeapply_test

import (
	"os"
	"testing"

	"github.com/datawire/dlib/dexec"
	"github.com/datawire/dlib/dlog"
	"github.com/emissary-ingress/emissary/v3/pkg/dtest"
	"github.com/emissary-ingress/emissary/v3/pkg/kubeapply"
)

func TestDocker(t *testing.T) {
	ctx := dlog.NewTestContext(t, false)

	if _, err := dexec.LookPath("docker"); err != nil {
		t.Skip(err)
	}

	if os.Getenv("DOCKER_REGISTRY") == "" {
		os.Setenv("DOCKER_REGISTRY", dtest.DockerRegistry(ctx))
	}

	_, err := kubeapply.ExpandResource(ctx, "docker.yaml")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

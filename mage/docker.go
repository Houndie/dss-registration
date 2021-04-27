package mage

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/moby/buildkit/frontend/dockerfile/dockerignore"
	"github.com/moby/moby/pkg/archive"
)

var HerokuProject = map[string]string{
	"testing": "dayton-swing-smackdown-testing",
}

func DockerBuild(ctx context.Context, client *client.Client, project, tag string) error {
	f, err := os.Open("dynamic/.dockerignore")
	if err != nil {
		return fmt.Errorf("error opening dockeringore file: %w", err)
	}

	excludes, err := dockerignore.ReadAll(f)
	if err != nil {
		return fmt.Errorf("error reading dockerignore file: %w", err)
	}

	buildCtx, err := archive.TarWithOptions("dynamic", &archive.TarOptions{
		ExcludePatterns: excludes,
	})
	if err != nil {
		return fmt.Errorf("error creating build context: %w", err)
	}

	_, err = client.ImageBuild(ctx, buildCtx, types.ImageBuildOptions{
		Tags:       []string{imageName(project, tag)},
		Dockerfile: "docker/Dockerfile.deploy",
	})
	if err != nil {
		return fmt.Errorf("error building docker image: %w", err)
	}

	return nil
}

func imageName(project, tag string) string {
	return fmt.Sprintf("registry.heroku.com/%s/%s", project, tag)
}

func DockerPush(ctx context.Context, client *client.Client, apiKey, project, tag string) error {
	_, err := client.ImagePush(ctx, imageName(project, tag), types.ImagePushOptions{
		RegistryAuth: base64.StdEncoding.EncodeToString([]byte("_:" + apiKey)),
	})
	if err != nil {
		return fmt.Errorf("error pushing docker image: %w", err)
	}

	return nil
}

package mage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/magefile/mage/mg"
	"github.com/moby/moby/builder/dockerignore"
	"github.com/moby/moby/pkg/archive"
	"github.com/moby/moby/pkg/jsonmessage"
	"github.com/moby/term"
)

func DockerBuild(ctx context.Context, root, dockerfile, name, tag string) error {
	mg.Deps(InitDockerClient)

	ignoreFile := filepath.Join(root, ".dockerignore")

	_, err := os.Stat(ignoreFile)
	var excludes []string
	if !os.IsNotExist(err) {
		f, err := os.Open(filepath.Join(root, ".dockerignore"))
		if err != nil {
			return fmt.Errorf("error opening dockeringore file: %w", err)
		}

		excludes, err = dockerignore.ReadAll(f)
		if err != nil {
			return fmt.Errorf("error reading dockerignore file: %w", err)
		}
	}

	buildCtx, err := archive.TarWithOptions(root, &archive.TarOptions{
		ExcludePatterns: excludes,
	})
	if err != nil {
		return fmt.Errorf("error creating build context: %w", err)
	}

	res, err := DockerClient().ImageBuild(ctx, buildCtx, types.ImageBuildOptions{
		Tags:       []string{DockerImageName(name, tag)},
		Dockerfile: dockerfile,
	})
	if err != nil {
		return fmt.Errorf("error building docker image: %w", err)
	}
	defer res.Body.Close()

	if err := ProcessDockerResponse(res.Body); err != nil {
		return err
	}

	return nil
}

func DockerImageName(project, tag string) string {
	return fmt.Sprintf("registry.heroku.com/%s/web:%s", project, tag)
}

func ProcessDockerResponse(body io.Reader) error {
	output := io.Discard
	if mg.Verbose() {
		output = os.Stderr
	}

	ptr, isTerm := term.GetFdInfo(output)
	if err := jsonmessage.DisplayJSONMessagesStream(body, output, ptr, isTerm, nil); err != nil {
		return fmt.Errorf("error in response from docker daemon: %w", err)
	}

	return nil
}

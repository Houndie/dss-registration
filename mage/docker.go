package mage

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/magefile/mage/mg"
	"github.com/moby/buildkit/frontend/dockerfile/dockerignore"
	"github.com/moby/moby/client"
	"github.com/moby/moby/pkg/archive"
	"github.com/moby/moby/pkg/jsonmessage"
	"github.com/moby/term"
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

	res, err := client.ImageBuild(ctx, buildCtx, types.ImageBuildOptions{
		Tags:       []string{imageName(project, tag)},
		Dockerfile: "docker/Dockerfile.deploy",
	})
	if err != nil {
		return fmt.Errorf("error building docker image: %w", err)
	}
	defer res.Body.Close()

	if err := processDockerResponse(res.Body); err != nil {
		return err
	}

	return nil
}

func processDockerResponse(body io.Reader) error {
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

func imageName(project, tag string) string {
	return fmt.Sprintf("registry.heroku.com/%s/web:%s", project, tag)
}

func DockerPush(ctx context.Context, client *client.Client, apiKey, project, tag string) error {
	buf, err := json.Marshal(&types.AuthConfig{
		Username: "_",
		Password: apiKey,
	})
	if err != nil {
		return fmt.Errorf("error encoding auth config: %w", err)
	}

	res, err := client.ImagePush(ctx, imageName(project, tag), types.ImagePushOptions{
		RegistryAuth: base64.URLEncoding.EncodeToString(buf),
	})
	if err != nil {
		return fmt.Errorf("error pushing docker image: %w", err)
	}
	defer res.Close()

	if err := processDockerResponse(res); err != nil {
		return err
	}

	return nil
}

func DockerSave(ctx context.Context, client *client.Client, project, tag, savefile string) error {
	f, err := os.OpenFile(savefile, os.O_WRONLY|os.O_CREATE, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("error opening %s for saving: %w", savefile, err)
	}

	res, err := client.ImageSave(ctx, []string{imageName(project, tag)})
	if err != nil {
		return fmt.Errorf("error saving docker image: %w", err)
	}
	defer res.Close()

	if _, err := io.Copy(f, res); err != nil {
		return fmt.Errorf("error copying from docker to output file: %w", err)
	}

	return nil
}

func DockerLoad(ctx context.Context, client *client.Client, savefile string) error {
	f, err := os.Open(savefile)
	if err != nil {
		return fmt.Errorf("error opening %s for loading: %w", savefile, err)
	}

	res, err := client.ImageLoad(ctx, f, true)
	if err != nil {
		return fmt.Errorf("error loading docker image: %w", err)
	}
	defer res.Body.Close()

	if err := processDockerResponse(res.Body); err != nil {
		return err
	}

	return nil
}

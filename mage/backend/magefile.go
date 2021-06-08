package backend

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/Houndie/dss-registration/mage"
	"github.com/docker/docker/api/types"
	"github.com/magefile/mage/mg"
	"github.com/moby/buildkit/frontend/dockerfile/dockerignore"
	"github.com/moby/moby/pkg/archive"
	"github.com/moby/moby/pkg/jsonmessage"
	"github.com/moby/term"
)

var HerokuProject = map[mage.WorkspaceType]string{
	"testing": "dayton-swing-smackdown-testing",
}

func Build(ctx context.Context) error {
	mg.Deps(mage.InitDeployVersion, mage.InitWorkspace, mage.InitDockerClient)

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

	res, err := mage.DockerClient().ImageBuild(ctx, buildCtx, types.ImageBuildOptions{
		Tags:       []string{imageName(HerokuProject[mage.Workspace()], mage.DeployVersion())},
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

func Deploy(ctx context.Context) error {
	mg.Deps(mage.InitHerokuAPIKey, mage.InitDeployVersion, mage.InitWorkspace, mage.InitDockerClient)

	buf, err := json.Marshal(&types.AuthConfig{
		Username: "_",
		Password: mage.HerokuAPIKey(),
	})
	if err != nil {
		return fmt.Errorf("error encoding auth config: %w", err)
	}

	res, err := mage.DockerClient().ImagePush(ctx, imageName(HerokuProject[mage.Workspace()], mage.DeployVersion()), types.ImagePushOptions{
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

func Save(ctx context.Context) error {
	mg.Deps(mage.InitDockerClient, mage.InitDockerCache, mage.InitWorkspace, mage.InitDeployVersion)

	f, err := os.OpenFile(mage.DockerCache(), os.O_WRONLY|os.O_CREATE, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("error opening %s for saving: %w", mage.DockerCache(), err)
	}

	res, err := mage.DockerClient().ImageSave(ctx, []string{imageName(HerokuProject[mage.Workspace()], mage.DeployVersion())})
	if err != nil {
		return fmt.Errorf("error saving docker image: %w", err)
	}
	defer res.Close()

	if _, err := io.Copy(f, res); err != nil {
		return fmt.Errorf("error copying from docker to output file: %w", err)
	}

	return nil
}

func Load(ctx context.Context) error {
	mg.Deps(mage.InitDockerCache, mage.InitDockerClient)

	f, err := os.Open(mage.DockerCache())
	if err != nil {
		return fmt.Errorf("error opening %s for loading: %w", mage.DockerCache(), err)
	}

	res, err := mage.DockerClient().ImageLoad(ctx, f, true)
	if err != nil {
		return fmt.Errorf("error loading docker image: %w", err)
	}
	defer res.Body.Close()

	if err := processDockerResponse(res.Body); err != nil {
		return err
	}

	return nil
}

func HealthCheck(ctx context.Context) error {
	mg.Deps(mage.InitBackendAddr)

	u, err := url.Parse(mage.BackendAddr())
	if err != nil {
		return fmt.Errorf("error parsing backend address \"%s\": %w", mage.BackendAddr(), err)
	}
	u.Path = path.Join(u.Path, "/twirp/dss.Info/Health")

	return mage.HealthCheck(ctx, http.MethodPost, u.String())
}

func VersionCheck(ctx context.Context) error {
	mg.Deps(mage.InitBackendAddr)

	u, err := url.Parse(mage.BackendAddr())
	if err != nil {
		return fmt.Errorf("error parsing backend address \"%s\": %w", mage.BackendAddr(), err)
	}
	u.Path = path.Join(u.Path, "/twirp/dss.Info/Version")

	return mage.VersionCheck(ctx, http.MethodPost, u.String())
}

func WaitForDeploy(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	for {
		if mg.Verbose() {
			fmt.Fprintln(os.Stderr, "performing health check")
		}

		err := HealthCheck(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return fmt.Errorf("backend not detected before timeout")
			}

			if mg.Verbose() {
				fmt.Fprintf(os.Stderr, "health check response: %s\n", err.Error())
			}

			time.Sleep(5 * time.Second)
			continue
		}

		if mg.Verbose() {
			fmt.Fprintln(os.Stderr, "performing version check")
		}

		err = VersionCheck(ctx)
		if err == nil {
			break
		}

		if ctx.Err() != nil {
			return fmt.Errorf("backend not detected before timeout")
		}

		if mg.Verbose() {
			fmt.Fprintf(os.Stderr, "version check response: %s\n", err.Error())
		}

		time.Sleep(5 * time.Second)
	}

	return nil
}

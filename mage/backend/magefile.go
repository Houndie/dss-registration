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
)

var HerokuProject = map[mage.WorkspaceType]string{
	"testing":    "dayton-swing-smackdown-testing",
	"production": "dayton-swing-smackdown-prod",
}

func Build(ctx context.Context) error {
	values := make([]string, 0, len(HerokuProject))
	for _, value := range HerokuProject {
		values = append(values, value)
	}

	return mage.DockerBuild(ctx, "dynamic", "docker/Dockerfile.deploy", values, mage.DeployVersion())
}

func Deploy(ctx context.Context) error {
	buf, err := json.Marshal(&types.AuthConfig{
		Username: "_",
		Password: mage.HerokuAPIKey(),
	})
	if err != nil {
		return fmt.Errorf("error encoding auth config: %w", err)
	}

	res, err := mage.DockerClient().ImagePush(ctx, mage.DockerImageName(HerokuProject[mage.Workspace()], mage.DeployVersion()), types.ImagePushOptions{
		RegistryAuth: base64.URLEncoding.EncodeToString(buf),
	})
	if err != nil {
		return fmt.Errorf("error pushing docker image: %w", err)
	}
	defer res.Close()

	if err := mage.ProcessDockerResponse(res); err != nil {
		return err
	}

	return nil
}

func Save(ctx context.Context) error {
	f, err := os.OpenFile(mage.DockerCache(), os.O_WRONLY|os.O_CREATE, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("error opening %s for saving: %w", mage.DockerCache(), err)
	}

	tags := make([]string, 0, len(HerokuProject))
	for _, project := range HerokuProject {
		tags = append(tags, mage.DockerImageName(project, mage.DeployVersion()))
	}

	res, err := mage.DockerClient().ImageSave(ctx, tags)
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
	f, err := os.Open(mage.DockerCache())
	if err != nil {
		return fmt.Errorf("error opening %s for loading: %w", mage.DockerCache(), err)
	}

	res, err := mage.DockerClient().ImageLoad(ctx, f, true)
	if err != nil {
		return fmt.Errorf("error loading docker image: %w", err)
	}
	defer res.Body.Close()

	if err := mage.ProcessDockerResponse(res.Body); err != nil {
		return err
	}

	return nil
}

func HealthCheck(ctx context.Context) error {
	terraformOutputs := mage.TerraformOutputs()

	u, err := url.Parse(terraformOutputs.BackendAddr)
	if err != nil {
		return fmt.Errorf("error parsing backend address \"%s\": %w", terraformOutputs.BackendAddr, err)
	}
	u.Path = path.Join(u.Path, "/twirp/dss.Info/Health")

	return mage.HealthCheck(ctx, http.MethodPost, u.String())
}

func VersionCheck(ctx context.Context) error {
	terraformOutputs := mage.TerraformOutputs()

	u, err := url.Parse(terraformOutputs.BackendAddr)
	if err != nil {
		return fmt.Errorf("error parsing backend address \"%s\": %w", terraformOutputs.BackendAddr, err)
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

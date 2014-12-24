package docker

import (
	"errors"
	"log"
	"os"
	"testing"

	dockerclient "github.com/koding/klient/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
	"github.com/koding/klient/Godeps/_workspace/src/github.com/koding/kite"
)

var (
	d                 *kite.Kite
	remote            *kite.Client
	TestContainerName = "dockertest"
	ErrNotFound       = errors.New("not found")
)

func init() {
	d = kite.New("docker", "0.0.1")
	d.Config.DisableAuthentication = true
	d.Config.Port = 3636
	d.Config.Username = "dockertest"

	dockerHost := os.Getenv("DOCKER_HOST")
	if dockerHost == "" {
		dockerHost = "tcp://192.168.59.103:2376" // darwin, boot2docker
	}

	dockerCertPath := os.Getenv("DOCKER_CERT_PATH")
	if dockerCertPath == "" {
		panic("please set DOCKER_CERT_PATH")
	}

	certFile := dockerCertPath + "/cert.pem"
	keyFile := dockerCertPath + "/key.pem"
	caFile := dockerCertPath + "/ca.pem"

	client, _ := dockerclient.NewTLSClient(dockerHost, certFile, keyFile, caFile)
	dock := &Docker{
		client: client,
	}

	d.HandleFunc("list", dock.List)
	d.HandleFunc("create", dock.Create)
	d.HandleFunc("destroy", dock.Destroy)
	d.HandleFunc("removeContainer", dock.RemoveContainer)

	go d.Run()
	<-d.ServerReadyNotify()

	remote = d.NewClient("http://127.0.0.1:3636/kite")
	err := remote.Dial()
	if err != nil {
		log.Fatal("err")
	}
}

func TestCreate(t *testing.T) {
	resp, err := remote.Tell("create", struct {
		Name  string
		Image string
	}{
		Name:  TestContainerName,
		Image: "ubuntu",
	})
	if err != nil {
		t.Fatal(err)
	}

	if TestContainerName != resp.MustString() {
		t.Errorf("container name is wrong, have '%s', want '%s'",
			resp.MustString(), TestContainerName)
	}
}

func TestList(t *testing.T) {
	container, err := getContainer(TestContainerName)
	if err != nil {
		t.Errorf("No image found with name '%s': %s\n", TestContainerName, err)
	}

	if container.Names[0] != "/"+TestContainerName {
		t.Errorf("container name is wrong, have '%s', want '%s'", container.Names[0], TestContainerName)
	}
}

func TestRemoveContainer(t *testing.T) {
	container, err := getContainer(TestContainerName)
	if err != nil {
		t.Errorf("No image found with name '%s': %s\n", TestContainerName, err)
	}

	if container.ID == "" {
		t.Error("container Id is empty, can't remove anything")
		return
	}

	_, err = remote.Tell("removeContainer", struct {
		ID string
	}{
		ID: container.ID,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = getContainer(TestContainerName)
	if err != nil && err != ErrNotFound {
		t.Errorf("No image found with name '%s': %s\n", TestContainerName, err)
	}
}

func getContainer(containerName string) (*dockerclient.APIContainers, error) {
	resp, err := remote.Tell("list")
	if err != nil {
		return nil, err
	}

	var containers []dockerclient.APIContainers

	err = resp.Unmarshal(&containers)
	if err != nil {
		return nil, err
	}

	for _, container := range containers {
		name := container.Names[0]
		// there is a slash in front of the names, so include it
		if name == "/"+containerName {
			return &container, nil
		}
	}

	return nil, ErrNotFound
}

// func TestDestroy(t *testing.T) {
// 	resp, err := remote.Tell("destroy", struct {
// 		Name string
// 	}{
// 		Name: TestContainerName,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	fmt.Printf("resp.MustBool() %+v\n", resp.MustBool())
// }

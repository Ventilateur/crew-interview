package mongodb

import (
	"context"
	"fmt"
	"github.com/Ventilateur/crew-interview/domain/talent"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"os"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

const (
	mongoImage    = "mongo:5.0.9"
	mongoHost     = "127.0.0.1"
	mongoPort     = "27017"
	mongoUsername = "root"
	mongoPassword = "root"
)

type TalentTestSuite struct {
	suite.Suite
	dockerClient *client.Client
	containerId  string
}

func (s *TalentTestSuite) SetupSuite() {
	// Create docker client
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	s.Nil(err)
	s.dockerClient = dockerClient

	// Pull mongo image
	reader, err := dockerClient.ImagePull(ctx, mongoImage, types.ImagePullOptions{})
	s.Nil(err)
	defer reader.Close()
	_, err = io.Copy(os.Stdout, reader)
	s.Nil(err)

	// Create mongo container
	resp, err := dockerClient.ContainerCreate(ctx,
		&container.Config{
			Image: mongoImage,
			ExposedPorts: nat.PortSet{
				mongoPort: {},
			},
			Env: []string{
				fmt.Sprintf("MONGO_INITDB_ROOT_USERNAME=%s", mongoUsername),
				fmt.Sprintf("MONGO_INITDB_ROOT_PASSWORD=%s", mongoPassword),
			},
		},
		&container.HostConfig{
			PortBindings: nat.PortMap{
				mongoPort: {{HostIP: mongoHost, HostPort: mongoPort}},
			},
		},
		nil,
		nil,
		"",
	)
	s.Nil(err)
	s.containerId = resp.ID

	// Start mongo container
	err = dockerClient.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	s.Nil(err)
}

func (s *TalentTestSuite) TearDownSuite() {
	// Stop mongo container
	timeout := 1 * time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err := s.dockerClient.ContainerStop(ctx, s.containerId, &timeout)
	s.Nil(err)
	err = s.dockerClient.ContainerRemove(ctx, s.containerId, types.ContainerRemoveOptions{RemoveVolumes: true})
	s.Nil(err)
}

func (s *TalentTestSuite) TestListAdd() {
	ctx := context.Background()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoUsername, mongoPassword, mongoHost, mongoPort)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	s.Nil(err)

	repo := NewTalentRepo(mongoClient.Database("test").Collection("test"))

	nbTalents := 5
	var talents []talent.Talent
	for i := 0; i < nbTalents; i++ {
		talents = append(talents, talent.Talent{
			Id:        fmt.Sprintf("id_%d", i),
			FirstName: fmt.Sprintf("fn_%d", i),
			LastName:  fmt.Sprintf("ln_%d", i),
			Picture:   fmt.Sprintf("pi_%d", i),
			Job:       fmt.Sprintf("jo_%d", i),
			Location:  fmt.Sprintf("lo_%d", i),
			LinkedIn:  fmt.Sprintf("li_%d", i),
			Github:    fmt.Sprintf("gh_%d", i),
			Twitter:   fmt.Sprintf("tw_%d", i),
			Tags: []string{
				fmt.Sprintf("tag_a_%d", i),
				fmt.Sprintf("tag_b_%d", i),
			},
			Stage: fmt.Sprintf("stg_%d", i),
		})

		err = repo.AddTalent(ctx, talents[i])
		s.Nil(err)
	}

	talentList, err := repo.ListTalents(ctx, "", 0)
	s.Nil(err)
	s.Equal(talents, talentList)
}

func TestTalentTestSuite(t *testing.T) {
	suite.Run(t, new(TalentTestSuite))
}

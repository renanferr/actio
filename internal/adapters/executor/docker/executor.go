package docker

type DockerExecutor struct{}

func NewDockerExecutor() *DockerExecutor {
	return &DockerExecutor{}
}

func (e *DockerExecutor) Execute() error {
	return nil
}

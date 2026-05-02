package kubernetes

type KubernetesExecutor struct{}

func NewKubernetesExecutor() *KubernetesExecutor {
	return &KubernetesExecutor{}
}

func (e *KubernetesExecutor) Execute() error {
	return nil
}

package kube

import (
	"fmt"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type RESTConfigFactory interface {
	New(kubeconfigPath string) (*rest.Config, error)
}

func NewDefaultRESTConfigFactory() *DefaultRESTConfigFactory {
	return &DefaultRESTConfigFactory{}
}

type DefaultRESTConfigFactory struct{}

func (r *DefaultRESTConfigFactory) New(kubeconfigPath string) (*rest.Config, error) {
	var errExt error

	if kubeconfigPath != "" {
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err == nil {
			return config, nil
		}

		errExt = err
	}

	config, errInt := rest.InClusterConfig()
	if errInt != nil {
		return nil, fmt.Errorf("cannot build config from flags nor create it in cluster: %w, %w", errExt, errInt)
	}

	return config, nil
}

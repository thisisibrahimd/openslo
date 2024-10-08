package v2alpha1

import "github.com/OpenSLO/OpenSLO/pkg/openslo"

var _ = openslo.Object(Service{})

type Service struct {
	APIVersion openslo.Version `yaml:"apiVersion"`
	Kind       openslo.Kind    `yaml:"kind"`
	Metadata   Metadata        `yaml:"metadata"`
	Spec       ServiceSpec     `yaml:"spec"`
}

func (s Service) GetVersion() openslo.Version {
	return APIVersion
}

func (s Service) GetKind() openslo.Kind {
	return openslo.KindService
}

func (s Service) GetName() string {
	return s.Metadata.Name
}

func (s Service) Validate() error {
	return nil
}

type ServiceSpec struct {
	Description string `yaml:"description,omitempty"`
}

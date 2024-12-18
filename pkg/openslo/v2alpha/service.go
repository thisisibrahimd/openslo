package v2alpha

import "github.com/OpenSLO/OpenSLO/pkg/openslo"

var _ = openslo.Object(Service{})

type Service struct {
	APIVersion openslo.Version `json:"apiVersion"`
	Kind       openslo.Kind    `json:"kind"`
	Metadata   Metadata        `json:"metadata"`
	Spec       ServiceSpec     `json:"spec"`
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
	Description string `json:"description,omitempty"`
}

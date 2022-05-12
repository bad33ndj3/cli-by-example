package secret

type Secret struct {
	Name       string
	secretName string
}

func (s *Secret) SetName(name string) {
	s.Name = name
}

func (s *Secret) GetName() string {
	return s.Name
}

func (s *Secret) GetSecretName() string {
	return s.secretName
}

type Safe interface {
	SetName(name string)
	GetSecretName() string
	GetName() string
}

func New(secret string) Safe {
	return &Secret{
		secretName: secret,
	}
}

package multiapikey

type Service interface { // For External usage, outside of current package only need to know MultiApiKey has below function
	ValidateApiKey(apikey string) (Client, bool)
}

type Client struct {
	Name    string   `yaml:"name"`
	Secrets []string `yaml:"secrets"`
}

type service struct { // For internal package usage
	ClientKey map[string]Client
}

/*
config file template

name: "default_client"
secrets:
  - api_key01
  - api_key02
*/
func New(cfgclient []Client) (Service, error) {
	clients := make(map[string]Client)
	for _, client := range cfgclient {
		for _, key := range client.Secrets {
			clients[key] = Client{
				Name:    client.Name,
				Secrets: client.Secrets,
			}
		}
	}
	return &service{
		ClientKey: clients,
	}, nil
}

func (s *service) ValidateApiKey(apikey string) (Client, bool) {
	client, ok := s.ClientKey[apikey]
	return client, ok
}

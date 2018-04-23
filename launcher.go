package stan

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for stan service with config file namespaces.
type Namespaces struct {
	NatsStream services.Namespace
}

// Launcher represents a stan launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	service *Service
	m       sync.Mutex
}

// NewLauncher returns a new stan Launcher.
func (Service *Service) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		service: Service,
		ns:      ns,
	}
}

// Up starts the stan service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	cfg := Config{}
	if err := cfg.Dial(configs[l.ns.NatsStream]); err != nil {
		// Add namespace key when returning error with logrus
		return err
	}
	return l.service.Dial(cfg)
}

// Down stops the stan service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.service.Close()
}

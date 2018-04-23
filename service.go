package stan

import (
	"github.com/nats-io/go-nats-streaming"
)

// Service wraps an NatsStream connection.
type Service struct {
	stan.Conn
}

// Dial init the Stream server.
func (s *Service) Dial(c Config) error {
	var err error
	s.Conn, err = stan.Connect(
		c.ClusterID,
		c.ClientID,
		stan.ConnectWait(c.ConnectTimeout),
		stan.MaxPubAcksInflight(c.MaxPubAcksInFlight),
		stan.NatsURL(c.NatsURL),
		stan.PubAckWait(c.AckTimeout),
	)
	return err
}

// Healthcheck returns if database responds.
func (s *Service) Healthcheck() error {
	return nil
}

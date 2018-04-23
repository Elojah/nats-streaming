package stan

import (
	"errors"
	"time"
)

// Config is a Stream server config.
type Config struct {
	ClusterID          string
	ClientID           string
	NatsURL            string
	ConnectTimeout     time.Duration
	AckTimeout         time.Duration
	MaxPubAcksInFlight int
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return c == rhs
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	var err error
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cClusterID, ok := fconf["cluster_id"]
	if !ok {
		return errors.New("missing key cluster_id")
	}
	if c.ClusterID, ok = cClusterID.(string); !ok {
		return errors.New("key cluster_id invalid. must be string")
	}

	cClientID, ok := fconf["client_id"]
	if !ok {
		return errors.New("missing key client_id")
	}
	if c.ClientID, ok = cClientID.(string); !ok {
		return errors.New("key client_id invalid. must be string")
	}

	cNatsURL, ok := fconf["nats_url"]
	if !ok {
		return errors.New("missing key nats_url")
	}
	if c.NatsURL, ok = cNatsURL.(string); !ok {
		return errors.New("key nats_url invalid. must be string")
	}

	cConnectTimeout, ok := fconf["connect_timeout"]
	if !ok {
		return errors.New("missing key connect_timeout")
	}
	cConnectTimeoutString, ok := cConnectTimeout.(string)
	if !ok {
		return errors.New("key connect_timeout invalid. must be string")
	}
	if c.ConnectTimeout, err = time.ParseDuration(cConnectTimeoutString); err != nil {
		return err
	}

	cAckTimeout, ok := fconf["ack_timeout"]
	if !ok {
		return errors.New("missing key ack_timeout")
	}
	cAckTimeoutString, ok := cAckTimeout.(string)
	if !ok {
		return errors.New("key ack_timeout invalid. must be string")
	}
	if c.AckTimeout, err = time.ParseDuration(cAckTimeoutString); err != nil {
		return err
	}

	cMaxPubAcksInFlight, ok := fconf["max_pub_acks_in_flight"]
	if !ok {
		return errors.New("missing key max_pub_acks_in_flight")
	}
	if c.MaxPubAcksInFlight, ok = cMaxPubAcksInFlight.(int); !ok {
		return errors.New("key max_pub_acks_in_flight invalid. must be int")
	}

	return nil
}

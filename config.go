/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package extfs

// Config epresents the configurable options for a filesystem.
type Config struct {
	// User specifies which HDFS user the client will act as. HDFS only
	User string

	// Addresses specifies the namenode(s) to connect to. HDFS only
	Addresses []string

	// UseDatanodeHostname specifies whether the client should connect to the
	// datanodes via hostname (which is useful in multi-homed setups) or IP
	// address
	UseDatanodeHostname bool
}

// ClientOption func for each Config argument
type ClientOption func(cfg *Config) error

// WithAddresses option to configure hdfs namenode addresses.
func WithAddresses(addrs []string) ClientOption {
	return func(cfg *Config) error {
		cfg.Addresses = addrs[:]
		return nil
	}
}

// WithUser option to configure hdfs user
func WithUser(user string) ClientOption {
	return func(cfg *Config) error {
		cfg.User = user
		return nil
	}
}

// WithUseDatanodeHostname option to configure use datanode hostname
func WithUseDatanodeHostname(use bool) ClientOption {
	return func(cfg *Config) error {
		cfg.UseDatanodeHostname = use
		return nil
	}
}

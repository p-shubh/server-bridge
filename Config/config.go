package config

import (
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

// SSHConfig holds the configuration for the SSH connection
type SSHConfig struct {
	User    string
	KeyPath string
	Server  string
}

// Connect establishes an SSH connection using a private key
func Connect(user, keyPath, server string) (*ssh.Client, error) {
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	client, err := ssh.Dial("tcp", server, config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

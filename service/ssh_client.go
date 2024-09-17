package services

import (
	"fmt"
	"io"
	"os"
	config "server-bridge/Config"

	"golang.org/x/crypto/ssh"
)

// SSHClient represents an SSH client
type SSHClient struct {
	Client *ssh.Client
}

// NewSSHClient initializes a new SSH client
func NewSSHClient(user, password, server string) (*SSHClient, error) {
	client, err := config.Connect(user, password, server)
	if err != nil {
		return nil, err
	}
	return &SSHClient{Client: client}, nil
}

// UploadFile uploads a file to the remote server
func (c *SSHClient) UploadFile(remotePath, localPath string) error {
	session, err := c.Client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	targetFile, err := session.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to create target file pipe: %w", err)
	}
	defer targetFile.Close()

	go func() {
		file, _ := os.Open(localPath)
		defer file.Close()
		io.Copy(targetFile, file)
	}()

	err = session.Run(fmt.Sprintf("cat > %s", remotePath))
	if err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}

	return nil
}

// DownloadFile downloads a file from the remote server
func (c *SSHClient) DownloadFile(remotePath, localPath string) error {
	session, err := c.Client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	targetFile, err := os.Create(localPath)
	if err != nil {
		return fmt.Errorf("failed to create target file: %w", err)
	}
	defer targetFile.Close()

	output, err := session.Output(fmt.Sprintf("cat %s", remotePath))
	if err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}

	targetFile.Write(output)
	return nil
}

// ListDirectory lists the details of the directory from the root folder
func (c *SSHClient) ListDirectory() (string, error) {
	session, err := c.Client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	output, err := session.Output("ls -la /")
	if err != nil {
		return "", fmt.Errorf("failed to run command: %w", err)
	}

	return string(output), nil
}

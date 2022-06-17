package sftp

import (
	"errors"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

// Dir an implementation of http.FileSystem for sftp
type Dir struct {
	root   string
	client *sftpClient
}

// NewDir returns a http.FileSystem instance for sftp
func NewDir(root string, address string, userName string, password string) (http.FileSystem, error) {
	root = strings.TrimSpace(root)
	if len(root) == 0 {
		root = "."
	}
	userName = strings.TrimSpace(userName)
	if len(userName) == 0 {
		return nil, errors.New("invalid username for sftp")
	}
	password = strings.TrimSpace(password)
	if len(password) == 0 {
		return nil, errors.New("invalid password for sftp")
	}
	client := newSFTPClient(address, userName, password, true)
	return &Dir{
		client: client,
		root:   root,
	}, client.Connect()
}

// Open opens the named file for reading
func (d *Dir) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	fullName := filepath.ToSlash(filepath.Join(d.root, filepath.FromSlash(path.Clean("/"+name))))
	return d.client.Open(fullName)
}

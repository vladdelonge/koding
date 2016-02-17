package req

import "time"

// MountFolder is the request struct for remote.mountFolder method.
type MountFolder struct {
	Name           string `json:"name"`
	LocalPath      string `json:"localPath"`
	RemotePath     string `json:"remotePath"`
	NoIgnore       bool   `json:"noIgnore"`
	NoPrefetchMeta bool   `json:"noPrefetchMeta"`
	PrefetchAll    bool   `json:"prefetchAll"`
	NoWatch        bool   `json:"noWatch"`
	CachePath      string `json:"cachePath"`
}

// UnmountFolder is the request struct for remote.UnmountFolder method.
type UnmountFolder struct {
	Name      string `json:"name"`
	LocalPath string `json:"localPath"`
}

// Exec is the request struct for remote.exec method.
type Exec struct {
	// TODO: Standardize this field name among all remote* methods.
	Machine string
	Command string
	Path    string
}

// SSHAuthSock is the request struct for remote.sshKeysAdd method.
type SSHKeyAdd struct {
	Name string
	Key  []byte
}

// Cache is the request struct for remote.cache method.
type Cache struct {
	Name       string        `json:"name"`
	LocalPath  string        `json:"localPath"`
	RemotePath string        `json:"remotePath"`
	Interval   time.Duration `json:interval`

	// Implementation details required by rsync currently. Not great to expose the
	// underlying implementation, but required currently.

	// The username of ssh user to connect to.
	Username string `json:"username"`

	// The SSH_AUTH_SOCK to set as an Environment variable when calling RSync.
	// Because RSync is being run from Klient, Klient may not have the SSH_AUTH_SOCK
	// var set for the calling user. This results in SSH failing.
	SSHAuthSock string `json:"sshAuthSock"`

	// The keypath that SSH will use for rsync.
	SSHPrivateKeyPath string `json:"sshPrivateKeyPath"`
}

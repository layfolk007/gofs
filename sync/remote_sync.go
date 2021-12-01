package sync

import (
	"github.com/no-src/gofs/contract"
	"github.com/no-src/gofs/core"
)

// NewRemoteSync auto create an instance of remoteServerSync or remoteClientSync according to src and target
func NewRemoteSync(src, target core.VFS, bufSize int, users []*contract.User) (Sync, error) {
	if src.Server() {
		return NewRemoteServerSync(src, target, bufSize)
	}
	return NewRemoteClientSync(src, target, bufSize, users)
}

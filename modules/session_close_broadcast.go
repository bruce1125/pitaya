package modules

import (
	"github.com/topfreegames/pitaya/v2/cluster"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/session"
)

// UniqueSession module watches for sessions using the same UID and kicks them
type SessionCloseBroadcast struct {
	Base
	rpcClient   cluster.RPCClient
	sessionPool session.SessionPool
}

// NewSessionCloseBroadcast creates a new session-close broadcast module
func NewSessionCloseBroadcast(rpcClient cluster.RPCClient, sessionPool session.SessionPool) *SessionCloseBroadcast {
	return &SessionCloseBroadcast{
		rpcClient:   rpcClient,
		sessionPool: sessionPool,
	}
}

// Init initializes the module
func (u *SessionCloseBroadcast) Init() error {
	u.sessionPool.OnSessionClose(func(s session.Session) {
		if s.UID() != "" {
			err := u.rpcClient.BroadcastSessionClosed(s.UID())
			if err != nil {
				logger.Log.Errorf("Failed to broadcast session-close event,UID=%s,err=%v", s.UID(), err)
			}
		}

	})
	return nil
}

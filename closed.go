package pitaya

import (
	"github.com/topfreegames/pitaya/v2/cluster"
	"github.com/topfreegames/pitaya/v2/constants"
)

// RegisterSessionClosedCallback register a listener to handle session closed event,only should be called in backend servers
func (app *App) RegisterSessionClosedCallback(listener cluster.RemoteClosedListener) error {
	if app.server.Frontend {
		return constants.ErrFrontendCantRegister
	}

	app.remoteService.AddRemoteClosedListener(listener)
	return nil
}

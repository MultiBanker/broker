package market

import (
	"net/http"

	_ "github.com/MultiBanker/broker/src/servers/adminhttp/dto"
)

// @Summary      VIN Connection
// @Description  VIN Connection
// @Tags         Markets
// @Accept       json
// @Produce      json
// @Security     ApiTokenAuth
// @Param        connect  body      dto.ConnectAuto  true  "auto to connect"
// @Success      204
// @Failure      400      {object}  httperrors.Response
// @Failure      401      {object}  httperrors.Response
// @Failure      404      {object}  httperrors.Response
// @Failure      422      {object}  httperrors.Response
// @Failure      500      {object}  httperrors.Response
// @Router       /api/v1/markets/auto/connect [put]
func (res resource) autoConnect(w http.ResponseWriter, r *http.Request) {

}

package api

import (
	"net/http"
)

// LastBlock godoc
// @Summary      Get last sepolia block number
// @Tags         get
// @Produce      json
// @Success      200  {object}  types.BlockNumber
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /eth/last_block [get]
func (api *API) lastBlock(w http.ResponseWriter, r *http.Request) {
	response, err := api.ethClient.GetLastBlockNumber()
	if err != nil {
		api.log.Err(err).Send()
		api.WriteError(w, r, err)
		return
	}

	api.WriteJSON(w, r, response)
}

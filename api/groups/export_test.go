package groups

import chainApiShared "github.com/TerraDharitri/drt-go-chain/api/shared"

// HandleErrorAndReturn -
func HandleHTTPError(err string) (int, chainApiShared.ReturnCode) {
	return handleHTTPError(err)
}

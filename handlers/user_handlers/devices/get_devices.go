package devices

import (
	"AltTube-Go/database"
	"AltTube-Go/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// GetDevices godoc
// @Summary Get devices
// @Description Get devices
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Device "GetDevices"
// @Security AccessToken
// @Router /user/devices [get]
func GetDevices(ctx *gin.Context) {
	// Get authorization header
	tokenString := ctx.GetHeader("Authorization")
	// Assuming currentRefreshToken is provided as 'Bearer <currentRefreshToken>'
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	authUserIDInterface, exists := ctx.Get("uuid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No UUID found in currentRefreshToken"})
		return
	}

	authUserID, ok := authUserIDInterface.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - UUID format invalid"})
		return
	}

	refreshTokens, err := database.GetRefreshTokenByUserID(authUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting devices for user"})
		return
	}
	if len(refreshTokens) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No devices found"})
		return
	}
	// Filter only the necessary fields
	var devices []models.Device
	for i := range refreshTokens {
		devices = append(devices, models.Device{
			ID:         refreshTokens[i].ID,
			LastActive: refreshTokens[i].CreatedAt,
			UserAgent:  refreshTokens[i].UserAgent,
			IPAddress:  refreshTokens[i].IPAddress,
		})
	}

	currentRefreshToken, err := database.GetRefreshTokenByAccessToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting current device"})
		return
	}

	var deviceList models.DeviceList
	deviceList.CurrentDeviceID = currentRefreshToken.ID
	deviceList.Devices = devices

	ctx.JSON(http.StatusOK, deviceList)
}

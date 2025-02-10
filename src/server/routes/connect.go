package routes

import (
	"image/png"
	"net/http"
	"whatsappbot/src/wpp"

	"github.com/boombuler/barcode"
	"github.com/gin-gonic/gin"
)

func ConnectQrCodeRoute(ctx *gin.Context) {
	client := wpp.ConnClient

	// Check if already authenticated
	if client.Store.ID != nil {
		ctx.String(http.StatusOK, "Já autenticado")
		return
	}

	// Generate or retrieve the QR code (assuming wpp.CodQrCode contains the QR code)
	qrCode := wpp.CodQrCode

	// Resize the QR code to increase the image size
	// Example: Resize to 300x300 pixels
	scaledQRCode, err := barcode.Scale(qrCode, 300, 300)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao redimensionar QR Code"})
		return
	}

	// Set content type for PNG
	ctx.Header("Content-Type", "image/png")

	// Encode the resized QR code to PNG and write it to the response
	err = png.Encode(ctx.Writer, scaledQRCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao codificar QR Code em PNG"})
		return
	}
}

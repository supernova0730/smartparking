package controllers

import (
	"github.com/gofiber/fiber/v2"
	"path/filepath"
	"smartparking/config"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/transport/response"
	"smartparking/pkg/tools"
)

type ImageController struct {
	m manager.Manager
}

type ImageResponse struct {
	Image string `json:"image"`
}

func NewImageController(m manager.Manager) *ImageController {
	return &ImageController{m: m}
}

// Upload godoc
// @Description upload image
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "image"
// @Success 200 {object} ImageResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /upload-image [post]
func (ctl *ImageController) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	newFilename := tools.FilenameWithCurrentTime(file.Filename)
	savePath := filepath.Join(config.GlobalConfig.Web.FileStorage, newFilename)

	err = c.SaveFile(file, savePath)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, ImageResponse{
		Image: newFilename,
	})
}

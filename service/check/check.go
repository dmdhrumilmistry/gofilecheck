package check

import (
	"github.com/gabriel-vasile/mimetype"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type CheckHandler struct{}

func NewCheckHandler() *CheckHandler {
	return &CheckHandler{}
}

func (h *CheckHandler) SetupRoutes(router fiber.Router) {
	router.Post("/check", h.Check)
}

func (h *CheckHandler) Check(c *fiber.Ctx) error {
	// accept form file
	file, err := c.FormFile("file")
	if err != nil {
		log.Error("Error in getting file from form: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "unable to get file from form",
		})
	}

	log.Info("File received: ", file.Filename)
	log.Info("File size: ", file.Size)
	log.Info("File type: ", file.Header.Get("Content-Type"))

	fileMultiPart, err := file.Open()
	if err != nil {
		log.Error("Error in opening file: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "unable to open file",
		})
	}

	defer fileMultiPart.Close()
	mtype, err := mimetype.DetectReader(fileMultiPart)
	if err != nil {
		log.Error("Error in detecting file type: ", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "unable to detect file type",
		})
	}

	log.Info("File detected type: ", mtype.String())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"filename":  file.Filename,
		"size":      file.Size,
		"type":      mtype.String(),
		"extension": mtype.Extension(),
	})
}

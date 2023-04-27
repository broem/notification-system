package handler

import (
	"net/http"

	"github.com/broem/notification-system/backend/internal/model"
	"github.com/broem/notification-system/backend/internal/repository"
	"github.com/labstack/echo/v4"
)

type NotificationHandler struct {
	repo repository.NotificationRepository
}

func NewNotificationHandler(repo repository.NotificationRepository) *NotificationHandler {
	return &NotificationHandler{repo: repo}
}

func (h *NotificationHandler) GetAllNotifications(c echo.Context) error {
	notifications, err := h.repo.GetAllNotifications()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch notifications"})
	}
	return c.JSON(http.StatusOK, notifications)
}

func (h *NotificationHandler) CreateNotification(c echo.Context) error {
	notification := new(model.Notification)
	if err := c.Bind(notification); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	resp, err := h.repo.CreateNotification(notification)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create notification"})
	}

	return c.JSON(http.StatusCreated, resp)
}

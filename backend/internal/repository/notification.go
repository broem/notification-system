package repository

import (
	"github.com/broem/notification-system/backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type NotificationRepository interface {
	GetAllNotifications() ([]*model.Notification, error)
}

type notificationRepository struct {
	db *sqlx.DB
}

func CreateNotificationRepository(db *sqlx.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) GetAllNotifications() ([]*model.Notification, error) {
	notifications := []*model.Notification{}
	err := r.db.Select(&notifications, "SELECT * FROM notifications")
	return notifications, err
}
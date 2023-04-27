package repository

import (
	"time"

	"github.com/broem/notification-system/backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type NotificationRepository interface {
	GetAllNotifications() ([]*model.Notification, error)
	CreateNotification(notification *model.Notification) (*model.Notification, error)
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

func (r *notificationRepository) CreateNotification(n *model.Notification) (*model.Notification, error) {
	n.CreatedAt = time.Now().UTC()
	query := "INSERT INTO notifications (title, message, created_at) VALUES ($1, $2, $3) RETURNING id;"
	err := r.db.QueryRow(query, n.Title, n.Message, n.CreatedAt).Scan(&n.ID)

	if err != nil {
		return nil, err
	}

	return n, nil
}
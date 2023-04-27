import React from 'react';
import styles from './NotificationListItem.module.css';
import { Notification } from '../../types/notification';

interface NotificationListItemProps {
  notification: Notification;
}

const NotificationListItem: React.FC<NotificationListItemProps> = ({ notification }) => {
  return (
    <div className={styles.notificationListItem}>
      <h3 className={styles.title}>{notification.title}</h3>
      <p className={styles.message}>{notification.message}</p>
    </div>
  );
};

export default NotificationListItem;
import React from 'react';
import NotificationListItem from './NotificationListItem';
import styles from './NotificationList.module.css';
import { Notification } from '../../types/notification';

interface NotificationListProps {
  notifications: Notification[];
}

const NotificationList: React.FC<NotificationListProps> = ({ notifications }) => {
  return (
    <div className={styles.notificationList}>
      {notifications.map((notification) => (
        <NotificationListItem key={notification.id} notification={notification} />
      ))}
    </div>
  );
};

export default NotificationList;
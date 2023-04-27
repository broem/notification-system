import React, { useState, useEffect } from 'react';
import Header from './components/Header/Header';
import NotificationList from './components/NotificationList/NotificationList';
import { fetchNotifications, createNotification } from './services/api';
import { Notification } from './types/notification';
import AddNotificationButton from './components/AddNotificationButton/AddNotificationButton';
import AddNotificationModal from './components/AddNotificationModal/AddNotificationModal';

const App: React.FC = () => {
  const [notifications, setNotifications] = useState<Notification[]>([]);
  const [modalOpen, setModalOpen] = useState(false);

  const handleModalClose = () => {
    setModalOpen(false);
  };

  useEffect(() => {
    const fetchData = async () => {
      const data = await fetchNotifications();
      setNotifications(data);
    };
    fetchData();
  }, []);

  const handleCreateNotification = async (title: string, message: string) => {
    try {
      console.log('Creating notification...');
      console.log(`Title: ${title}`);
      console.log(`Message: ${message}`);
      const newNotification = await createNotification(title, message);
      setNotifications([newNotification, ...notifications]);
    } catch (error: any) {
      console.error(`Error creating notification: ${error.message}`);
    }
  };

  return (
    <div>
      <Header />
      <NotificationList notifications={notifications} />
      <AddNotificationButton onClick={() => setModalOpen(true)} />
      <AddNotificationModal
        isOpen={modalOpen}
        onClose={handleModalClose}
        onSubmit={handleCreateNotification}
      />
    </div>
  );
};

export default App;
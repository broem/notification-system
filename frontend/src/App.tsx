import React, { useState, useEffect } from 'react';
import Header from './components/Header/Header';
import NotificationList from './components/NotificationList/NotificationList';
import { fetchNotifications } from './services/api';
import { Notification } from './types/notification';

const App: React.FC = () => {
  const [notifications, setNotifications] = useState<Notification[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      const data = await fetchNotifications();
      setNotifications(data);
    };
    fetchData();
  }, []);

  return (
    <div>
      <Header />
      <NotificationList notifications={notifications} />
    </div>
  );
};

export default App;
import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api/v1';

export const fetchNotifications = async () => {
  const response = await axios.get(`${API_URL}/notifications`);
  return response.data;
};
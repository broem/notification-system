import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api/v1';

export const fetchNotifications = async () => {
  const response = await axios.get(`${API_URL}/notifications`);
  return response.data;
};

export async function createNotification(title: string, message: string) {
  try {
    const response = await axios.post(`${API_URL}/notifications`, {
      title: title, 
      message: message,
    });

    if (response.status === 201) {
      return response.data;
    } else {
      throw new Error(`Error creating notification: ${response.status} ${response.statusText}`);
    }
  } catch (error: any) {
    throw new Error(`Error sending request: ${error.message}`);
  }
}
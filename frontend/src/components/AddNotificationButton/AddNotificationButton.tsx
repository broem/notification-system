import React from 'react';
import styles from './AddNotificationButton.module.css';

interface AddNotificationButtonProps {
  onClick: () => void;
}

const AddNotificationButton: React.FC<AddNotificationButtonProps> = ({ onClick }) => {
  return (
    <button className={styles.addNotificationButton} onClick={onClick}>
      <i className="fas fa-plus"></i>
    </button>
  );
};

export default AddNotificationButton;
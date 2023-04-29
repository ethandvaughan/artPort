import { useState } from 'react';

export default function useId() {
  const getId = () => {
    const idString = localStorage.getItem('user_id');
    const userId = JSON.parse(idString);
    return userId;
  };

  const [user_id, setId] = useState(getId());

  const saveId = (userId) => {
    localStorage.setItem('user_id', JSON.stringify(userId));
    setId(userId.user_id);
  };

  return {
    setId: saveId,
    user_id,
  };
}

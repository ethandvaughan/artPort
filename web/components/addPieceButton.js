import { useState } from 'react';
import Add from './addPiece';
import styles from './addPieceButton.module.css';

export default function PopupButton() {
  const [showPopup, setShowPopup] = useState(false);

  const handleClick = () => {
    setShowPopup(!showPopup);
  };

  return (
    <div className='flex'>
      <button className={styles.addButton} onClick={handleClick}>
        <span className={styles.plus}>+</span>
      </button>
      {showPopup && <Add setShowPopup={setShowPopup} />}
    </div>
  );
}

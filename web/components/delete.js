import React, { useState } from 'react';
import styles from './addPiece.module.css';

const Delete = (props) => {
  const handleConfirm = (event) => {
    fetch(`http://localhost:8080/piece/${props.id}`, {
      method: 'DELETE',
    })
      .then((response) => {
        console.log(response);
      })
      .catch((error) => {
        console.error(error);
      });
    window.location.reload();
  };
  return (
    <div className={`${styles.popup} drop-shadow-lg z-10`}>
      <div className={styles.popupContent}>
        <div className='mb-4'>
          <label htmlFor='confirm-delete' className='block font-medium text-gray-700'>
            Are you sure you would like to delete {props.title}?
          </label>
        </div>
        <div className='relative'>
          <button className='' type='button' onClick={() => props.setShowDelete(false)}>
            Cancel
          </button>
          <button
            className='absolute bottom-0 right-0 px-4 py-2 bg-blue-500 text-white rounded'
            type='submit'
            onClick={handleConfirm}
          >
            Delete
          </button>
        </div>
      </div>
    </div>
  );
};

export default Delete;

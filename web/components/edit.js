import { useState } from 'react';
import styles from './edit.module.css';

const Edit = () => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [showPopup, setShowPopup] = useState(false);

  const handlePopup = () => {
    setShowPopup(!showPopup);
  };

  const handleTitleChange = (event) => {
    setTitle(event.target.value);
  };

  const handleDescriptionChange = (event) => {
    setDescription(event.target.value);
  };

  const handleSaveChanges = () => {
    // Handle saving changes here
  };

  return (
    <>
      <button className='float-left' onClick={handlePopup}>
        <img src='/edit.png' style={{ height: '17px', width: '17px' }} />
      </button>
      {showPopup && (
        <div className={`${styles.edit} z-10`}>
          <div className={`${styles.editContainer}`}>
            <h2 className='font-bold mb-4'>Edit Artwork</h2>
            <div className='mb-4'>
              <label htmlFor='title' className='block font-bold mb-2'>
                Title
              </label>
              <input
                type='text'
                id='title'
                value={title}
                onChange={handleTitleChange}
                className='w-full border border-gray-400 p-2 rounded'
              />
            </div>
            <div className='mb-4'>
              <label htmlFor='description' className='block font-bold mb-2'>
                Description
              </label>
              <textarea
                id='description'
                value={description}
                onChange={handleDescriptionChange}
                className='w-full border border-gray-400 p-2 rounded'
              />
            </div>
            <div className='flex justify-end'>
              <button
                className='bg-gray-400 text-white px-4 py-2 rounded mr-2'
                onClick={handlePopup}
              >
                Close
              </button>
              <button
                className='bg-blue-500 text-white px-4 py-2 rounded'
                onClick={handleSaveChanges}
              >
                Save Changes
              </button>
            </div>
          </div>
        </div>
      )}
    </>
  );
};

export default Edit;

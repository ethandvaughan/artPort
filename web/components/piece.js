import { useState } from 'react';
import Delete from './delete';
import Edit from './edit';
import ImageCarousel from './imageCarousel';
import styles from './piece.module.css';

const Piece = (props) => {
  const [showDelete, setShowDelete] = useState(false);
  const [showEdit, setShowEdit] = useState(false);
  const handleShowDelete = () => {
    setShowDelete(!showDelete);
  };

  const handleShowEdit = () => {
    setShowEdit(!showEdit);
  };

  return (
    <div className={`${styles.artwork}`}>
      <div className={styles.details}>
        <ImageCarousel images={props.images} />
        <h2 className='font-bold italic'>{props.title}</h2>
        <h2>{new Date(props.date).getFullYear()}</h2>
        <h3>By: {props.artist}</h3>
        <p>Medium: {props.category}</p>
        <button className='float-left' onClick={handleShowEdit}>
          <img src='/edit.png' style={{ height: '18px', width: '18px' }} />
        </button>
        {showEdit && (
          <Edit
            setShowEdit={setShowEdit}
            id={props.artwork.id}
            title={props.title}
            artist={props.artist}
            category={props.category}
            size={props.artwork.size}
            description={props.artwork.description}
            date={props.date}
            clay={props.artwork.clay}
            bisque={props.artwork.bisque_cone}
            glaze={props.artwork.glaze_cone}
            glazeDescription={props.artwork.glaze_description}
          />
        )}
        <button className='float-right' onClick={handleShowDelete}>
          <img src='/trash.webp' style={{ height: '20px', width: '20px' }} />
        </button>
        {showDelete && <Delete setShowDelete={setShowDelete} id={props.id} title={props.title} />}
      </div>
    </div>
  );
};

export default Piece;

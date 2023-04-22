import Edit from './edit';
import ImageCarousel from './imageCarousel';
import styles from './piece.module.css';

const Piece = (props) => {
  const handleDelete = (event) => {
    event.preventDefault();
    fetch(`http://localhost:8080/pieces/${props.id}`, {
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
    <div className={`${styles.artwork}`}>
      <div className={styles.details}>
        <ImageCarousel images={props.images} />
        <h2 className='font-bold italic'>{props.title}</h2>
        <h2>{new Date(props.date).getFullYear()}</h2>
        <h3>By: {props.artist}</h3>
        <p>Medium: {props.category}</p>
        <Edit />
        <button className='float-right' onClick={handleDelete}>
          <img src='/trash.webp' style={{ height: '20px', width: '20px' }} />
        </button>
      </div>
    </div>
  );
};

export default Piece;

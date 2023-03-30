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
        <h2>{props.title}</h2>
        <h3>by {props.artist}</h3>
        <p>Medium: {props.category}</p>
        <button className={styles.deleteButton} onClick={handleDelete}>
          Delete
        </button>
      </div>
    </div>
  );
};

export default Piece;

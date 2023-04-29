import { useState, useEffect } from 'react';
import PopupButton from './addPieceButton';
import Piece from './piece';
import useId from 'components/useId';

const Pieces = () => {
  const [data, setData] = useState(null);
  const { user_id, setId } = useId();

  useEffect(() => {
    fetch(`http://localhost:8080/pieces/${user_id}`)
      .then((response) => response.json())
      .then((data) => setData(data))
      .catch((error) => console.error(error));
  }, []);

  return (
    <div>
      <PopupButton />
      <div className='grid gap-4 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 mx-8 my-4'>
        {data ? (
          data.map((artwork) => (
            <Piece
              key={artwork.id}
              id={artwork.id}
              title={artwork.title}
              date={artwork.date}
              artist={artwork.artist}
              category={artwork.category}
              images={artwork.images}
              artwork={artwork}
            />
          ))
        ) : (
          <div>No artwork to display</div>
        )}
      </div>
    </div>
  );
};

export default Pieces;

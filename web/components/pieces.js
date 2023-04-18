import { useState, useEffect } from 'react';
import PopupButton from './addPieceButton';
import Piece from './piece';

const Pieces = () => {
  const [data, setData] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080/pieces')
      .then((response) => response.json())
      .then((data) => setData(data))
      .catch((error) => console.error(error));
  }, []);

  return (
    <div>
      <PopupButton />
      <div className='grid gap-4 grid-cols-4 mx-8 my-4'>
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
            />
          ))
        ) : (
          <div>Loading...</div>
        )}
      </div>
    </div>
  );
};

export default Pieces;

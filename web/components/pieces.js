import { useState, useEffect } from 'react';
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
      {data ? (
        data.map((artwork) => (
          <Piece
            key={artwork.id}
            id={artwork.id}
            title={artwork.title}
            artist={artwork.artist}
            category={artwork.category}
          />
        ))
      ) : (
        <div>Loading...</div>
      )}
    </div>
  );
};

export default Pieces;

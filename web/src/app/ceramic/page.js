'use client';
import PopupButton from 'components/addPieceButton';
import Piece from 'components/piece';
import useToken from 'components/useToken';
import { useEffect, useState } from 'react';

const Ceramic = () => {
  const [data, setData] = useState(null);
  const { token, setToken } = useToken();

  useEffect(() => {
    fetch('http://localhost:8080/ceramic')
      .then((response) => response.json())
      .then((data) => setData(data))
      .catch((error) => console.error(error));
  }, []);

  if (!token) {
    window.location.href = '/login';
  }

  return (
    <>
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
          <div>Loading...</div>
        )}
      </div>
    </>
  );
};

export default Ceramic;

'use client';
import PopupButton from 'components/addPieceButton';
import Footer from 'components/footer';
import Header from 'components/header';
import Piece from 'components/piece';
import { useEffect, useState } from 'react';

const Ceramic = () => {
  const [data, setData] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080/ceramic')
      .then((response) => response.json())
      .then((data) => setData(data))
      .catch((error) => console.error(error));
  }, []);

  return (
    <>
      <Header title='Arfol' />
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
            />
          ))
        ) : (
          <div>Loading...</div>
        )}
      </div>
      <Footer />
    </>
  );
};

export default Ceramic;

'use client';
import Header from 'components/header';
import Piece from 'components/piece';
import Remove from 'components/removePiece';
import Footer from 'components/footer';
import PopupButton from 'components/addPieceButton';

export default function Home() {
  return (
    <>
      <Header title='Arfol' />
      <PopupButton />
      <Piece />
      <Remove />
      <Footer />
    </>
  );
}

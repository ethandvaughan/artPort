'use client';
import Header from 'components/header';
import Piece from 'components/piece';
import Add from 'components/addPiece';
import Remove from 'components/removePiece';
import Footer from 'components/footer';

export default function Home() {
  return (
    <>
      <Header title='Arfol' />
      <Piece />
      <Add />
      <Remove />
      <Footer />
    </>
  );
}

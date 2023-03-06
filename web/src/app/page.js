'use client';
import Header from 'components/header';
import Pieces from 'components/pieces';
import Footer from 'components/footer';
import PopupButton from 'components/addPieceButton';

export default function Home() {
  return (
    <>
      <Header title='Arfol' />
      <PopupButton />
      <Pieces />
      <Footer />
    </>
  );
}

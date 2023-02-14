import React from 'react';
import Header from '../components/header';
import Footer from '../components/footer';
import Piece from '../components/piece';
import Add from '../components/addPiece';
import Remove from '../components/removePiece';

const Home = () => {
  return (
    <div>
      <Header />
      <Piece />
      <Add />
      <Remove />
      <Footer />
    </div>
  );
}

export default Home;
'use client';
import Header from 'components/header';
import Pieces from 'components/pieces';
import Footer from 'components/footer';
import Login from 'components/login';
import useToken from 'components/useToken';

export default function Home() {
  const { token, setToken } = useToken();

  if (!token) {
    return (
      <>
        <Header title='Arfol' />
        <Login setToken={setToken} />
        <Footer />
      </>
    );
  }

  return (
    <>
      <Header title='Arfol' />
      <Pieces />
      <Footer />
    </>
  );
}

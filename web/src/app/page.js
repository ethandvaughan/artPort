'use client';
import Pieces from 'components/pieces';
import useToken from 'components/useToken';
import Welcome from 'components/welcome';

export default function Home() {
  const { token, setToken } = useToken();

  if (!token) {
    return (
      <>
        <Welcome />
      </>
    );
  }

  return (
    <>
      <Pieces />
    </>
  );
}

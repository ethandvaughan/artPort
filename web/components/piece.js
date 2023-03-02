import { useState, useEffect } from 'react';
import styles from './piece.module.css';

const Piece = () => {
  const [data, setData] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080/pieces')
      .then((response) => response.json())
      .then((data) => setData(data))
      .catch((error) => console.error(error));
  }, []);

  console.log(data);

  return <div>{data ? <pre>{JSON.stringify(data, null, 2)}</pre> : <div>Loading...</div>}</div>;
};

export default Piece;

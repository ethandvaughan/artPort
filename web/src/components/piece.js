import React, { useState, useEffect } from "react";

const Piece = () => {
  const [data, setData] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080/pieces')
      .then(response => response.json())
      .then(data => setData(data))
      .catch(error => console.error(error));
  }, []);

  return data ? (
    <ul>
      {data.map(item => (
        <li key={item.id}>{item.name}</li>
      ))}
    </ul>
  ) : (
    <p>Loading...</p>
  );
};

export default Piece;





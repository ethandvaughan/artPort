import React, { useState } from 'react';

const Add = () => {
  const [response, setResponse] = useState(null);
  const [idInput, setId] = useState("");
  const [titleInput, setTitle] = useState("");
  const [artistInput, setArtist] = useState("");
  const [categoryInput, setCategory] = useState("");

  const handleSubmit = async (event) => {
    event.preventDefault();

    const response = await fetch('http://localhost:8080/pieces', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        id: idInput,
        title: titleInput,
        artist: artistInput,
        category: categoryInput 
      })
    });

    const json = await response.json();
    setResponse(json);

    window.location.reload();
  };

  return (
    <form onSubmit={handleSubmit}>
      <ol>
      <li>ID: <input type="text" value={idInput} onChange={(event) => setId(parseInt(event.target.value))} /></li>
      <li>Title: <input type="text" value={titleInput} onChange={(event) => setTitle(event.target.value)} /></li>
      <li>Artist: <input type="text" value={artistInput} onChange={(event) => setArtist(event.target.value)} /></li>
      <li>Category: <input type="text" value={categoryInput} onChange={(event) => setCategory(event.target.value)} /></li>
      </ol>
      <button type="submit">Submit</button>
      {response ? (
        <pre>
          {JSON.stringify(response, null, 2)}
        </pre>
      ) : null}
    </form>
  );

}

export default Add;
import React, { useState } from "react";

const Remove = () => {
  const [idInput, setId] = useState("");

  const handleDelete = (event) => {
    event.preventDefault();
    fetch(`http://localhost:8080/pieces/${idInput}`, {
      method: 'DELETE'
    })
    .then(response => {
      console.log(response);
    })
    .catch(error => {
      console.error(error);
    });
  };

  return (
    <form>
      <p>ID of piece to delete: </p>
      <input type="text" value={idInput} onChange={event => setId(event.target.value)} />
      <br />
      <button type="submit" onClick={handleDelete}>Delete</button>
    </form>
  );
}

export default Remove;
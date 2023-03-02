import { useState, useEffect } from 'react';
import styles from './addPiece.module.css';

const Add = (props) => {
  const [response, setResponse] = useState(null);
  const [titleInput, setTitle] = useState('');
  const [artistInput, setArtist] = useState('');
  const [categoryInput, setCategory] = useState('');
  const [sizeInput, setSize] = useState('');
  const [descriptionInput, setDescription] = useState('');
  const [dateInput, setDate] = useState('');
  const [clayTypeInput, setClayType] = useState('');
  const [bisqueConeInput, setBisqueCone] = useState('');
  const [glazeConeInput, setGlazeCone] = useState('');
  const [glazeDescriptionInput, setGlazeDescription] = useState('');

  const [clays, setClays] = useState([]);
  const [cones, setCones] = useState([]);

  const handleClose = () => {
    props.setShowPopup(false);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    const response = await fetch('http://localhost:8080/pieces', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        title: titleInput,
        artist: artistInput,
        glaze_description: {
          String: glazeDescriptionInput,
          Valid: true,
        },
        clay: {
          String: clayTypeInput,
          Valid: true,
        },
        bisque_cone: {
          String: bisqueConeInput,
          Valid: true,
        },
        glaze_cone: {
          String: glazeConeInput,
          Valid: true,
        },
        date: dateInput + 'T00:00:00Z',
        category: categoryInput,
        description: {
          String: descriptionInput,
          Valid: true,
        },
        size: {
          String: sizeInput,
          Valid: true,
        },
      }),
    });

    const json = await response.json();
    setResponse(json);

    window.location.reload();
  };

  const categories = [
    'Acrylic',
    'Ceramic',
    'Charcoal',
    'Digital Art',
    'Fabric',
    'Gouache',
    'Graphic Design',
    'Graphite',
    'Ink',
    'Mixed Media',
    'Oil',
    'Photography',
    'Print Making',
    'Watercolor',
    'Other',
  ];

  useEffect(() => {
    fetch('http://localhost:8080/clays')
      .then((response) => response.json())
      .then((data) => setClays(data))
      .catch((error) => console.error(error));
  }, []);

  useEffect(() => {
    fetch('http://localhost:8080/cones')
      .then((response) => response.json())
      .then((data) => setCones(data))
      .catch((error) => console.error(error));
  }, []);

  return (
    <div className={styles.popup}>
      <div className={styles.popupContent}>
        <button className={styles.closeButton} onClick={handleClose}>
          <span>X</span>
        </button>
        <form onSubmit={handleSubmit}>
          <ol>
            <li>
              Title:{' '}
              <input
                type='text'
                value={titleInput}
                onChange={(event) => setTitle(event.target.value)}
              />
            </li>
            <li>
              Artist:{' '}
              <input
                type='text'
                value={artistInput}
                onChange={(event) => setArtist(event.target.value)}
              />
            </li>
            <li>
              Medium:{' '}
              <select value={categoryInput} onChange={(event) => setCategory(event.target.value)}>
                <option value=''>--Select medium--</option>
                {categories.map((category, index) => (
                  <option key={index} value={category}>
                    {category}
                  </option>
                ))}
              </select>
            </li>

            {categoryInput === 'Ceramic' && (
              <>
                <li>
                  Clay Type:{' '}
                  <select
                    value={clayTypeInput}
                    onChange={(event) => setClayType(event.target.value)}
                  >
                    <option value=''>--Select Category--</option>
                    {clays.map((clay, index) => (
                      <option key={index} value={clay}>
                        {clay}
                      </option>
                    ))}
                  </select>
                </li>
                <li>
                  Bisque Cone:{' '}
                  <select
                    value={bisqueConeInput}
                    onChange={(event) => setBisqueCone(event.target.value)}
                  >
                    <option value=''>--Select Bisque Cone--</option>
                    {cones.map((cone, index) => (
                      <option key={index} value={cone}>
                        {cone}
                      </option>
                    ))}
                  </select>
                </li>
                <li>
                  Glaze Cone:{' '}
                  <select
                    value={glazeConeInput}
                    onChange={(event) => setGlazeCone(event.target.value)}
                  >
                    <option value=''>--Select Glaze Cone--</option>
                    {cones.map((cone, index) => (
                      <option key={index} value={cone}>
                        {cone}
                      </option>
                    ))}
                  </select>
                </li>
                <li>
                  Glaze Description:{' '}
                  <input
                    type='text'
                    value={glazeDescriptionInput}
                    onChange={(event) => setGlazeDescription(event.target.value)}
                  />
                </li>
              </>
            )}
            <li>
              Size:{' '}
              <input
                type='text'
                value={sizeInput}
                onChange={(event) => setSize(event.target.value)}
              />
            </li>
            <li>
              Date:{' '}
              <input
                type='date'
                value={dateInput}
                onChange={(event) => setDate(event.target.value)}
              />
            </li>
            <li>
              Description:{' '}
              <input
                type='text'
                value={descriptionInput}
                onChange={(event) => setDescription(event.target.value)}
              />
            </li>
          </ol>
          <button type='submit'>Submit</button>
          {response ? <pre>{JSON.stringify(response, null, 2)}</pre> : null}
        </form>
      </div>
    </div>
  );
};

export default Add;

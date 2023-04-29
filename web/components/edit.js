import { useEffect, useState } from 'react';
import styles from './addPiece.module.css';

const Edit = (props) => {
  const [titleInput, setTitle] = useState(props.title);
  const [artistInput, setArtist] = useState(props.artist);
  const [categoryInput, setCategory] = useState(props.category);
  const [sizeInput, setSize] = useState(props.size.String);
  const [descriptionInput, setDescription] = useState(props.description.String);
  const [dateInput, setDate] = useState(props.date.slice(0, 10));
  const [clayTypeInput, setClayType] = useState(props.clay.String);
  const [bisqueConeInput, setBisqueCone] = useState(props.bisque.String);
  const [glazeConeInput, setGlazeCone] = useState(props.glaze.String);
  const [glazeDescriptionInput, setGlazeDescription] = useState(props.glazeDescription.String);

  const [clays, setClays] = useState([]);
  const [cones, setCones] = useState([]);

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

  const handleSubmit = async (event) => {
    event.preventDefault();

    const response = await fetch(`http://localhost:8080/piece/${props.id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        id: props.id,
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

    window.location.reload();
    console.log(response);
  };

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
    <div className={`${styles.popup} drop-shadow-lg z-10`}>
      <div className={styles.popupContent}>
        <button
          className={styles.closeButton}
          onClick={() => {
            props.setShowEdit(false);
          }}
        >
          <span>X</span>
        </button>
        <form onSubmit={handleSubmit}>
          <ol>
            <li>
              Title:{' '}
              <input
                required
                className='block bg-white w-full border border-slate-300 rounded-md'
                type='text'
                value={titleInput}
                onChange={(event) => setTitle(event.target.value)}
              />
            </li>
            <li>
              Artist:{' '}
              <input
                required
                className='block bg-white w-full border border-slate-300 rounded-md'
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
                    className='block bg-white w-full border border-slate-300 rounded-md'
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
                className='block bg-white w-full border border-slate-300 rounded-md'
                type='text'
                value={sizeInput}
                onChange={(event) => setSize(event.target.value)}
              />
            </li>
            <li>
              Date:{' '}
              <input
                required
                type='date'
                value={dateInput}
                onChange={(event) => setDate(event.target.value)}
              />
            </li>
            <li>
              Description:{' '}
              <input
                className='block bg-white w-full border border-slate-300 rounded-md'
                type='text'
                value={descriptionInput}
                onChange={(event) => setDescription(event.target.value)}
              />
            </li>
          </ol>
          <button
            className='mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded'
            type='submit'
          >
            Submit
          </button>

          {/*response ? <pre>{JSON.stringify(response, null, 2)}</pre> : null*/}
        </form>
      </div>
    </div>
  );
};

export default Edit;

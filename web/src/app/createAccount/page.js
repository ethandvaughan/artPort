'use client';
import { useState } from 'react';
import { createHash } from 'crypto';

async function createUser(credentials) {
  console.log(JSON.stringify(credentials));

  return fetch('http://localhost:8080/user', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(credentials),
  });
}

const CreateAccount = () => {
  const [first_name, setFName] = useState('');
  const [last_name, setLName] = useState('');
  const [username, setUsername] = useState('');
  const [prePassword, setPassword] = useState('');
  const [errorMessage, setErrorMessage] = useState('');

  const handleCreateAccount = async (event) => {
    event.preventDefault();
    // Validate email and password

    const hash = createHash('sha256');
    hash.update(prePassword);
    var password = hash.digest('hex');

    // Create account with Go api
    const user = await createUser({
      username,
      password,
      first_name,
      last_name,
    });
    // window.location.href = '/login';
  };

  return (
    <div className='flex-col items-center justify-center'>
      <h1 className='text-4xl font-bold mb-4'>Create Account</h1>
      <form onSubmit={handleCreateAccount} className='flex-col'>
        <label htmlFor='first_name' className='font-bold'>
          First Name
        </label>
        <input
          type='first_name'
          id='first_name'
          value={first_name}
          onChange={(event) => setFName(event.target.value)}
          required
          className='border border-gray-300 p-2'
        />
        <label htmlFor='last_name' className='font-bold'>
          Last Name
        </label>
        <input
          type='last_name'
          id='last_name'
          value={last_name}
          onChange={(event) => setLName(event.target.value)}
          required
          className='border border-gray-300 p-2'
        />
        <label htmlFor='username' className='font-bold'>
          Username
        </label>
        <input
          type='username'
          id='username'
          value={username}
          onChange={(event) => setUsername(event.target.value)}
          required
          className='border border-gray-300 p-2'
        />
        <label htmlFor='password' className='font-bold'>
          Password
        </label>
        <input
          type='password'
          id='password'
          value={prePassword}
          onChange={(event) => setPassword(event.target.value)}
          required
          className='border border-gray-300 p-2'
        />
        {errorMessage && <p className='text-red-500'>{errorMessage}</p>}
        <button type='submit' className='bg-blue-500 text-white py-2 px-4 rounded'>
          Create Account
        </button>
      </form>
    </div>
  );
};

export default CreateAccount;

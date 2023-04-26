'use client';
import { useState } from 'react';
import { createHash } from 'crypto';
import useToken from 'components/useToken';
import Link from 'next/link';

async function loginUser(credentials) {
  console.log(JSON.stringify(credentials));
  return fetch('http://localhost:8080/auth', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(credentials),
  }).then((data) => data.json());
}

export default function Login() {
  const [username, setUsername] = useState('');
  const [prePassword, setPassword] = useState('');
  const [error, setError] = useState('');
  const { token, setToken } = useToken();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);
    const hash = createHash('sha256');
    hash.update(prePassword);
    var password = hash.digest('hex');
    const token = await loginUser({
      username,
      password,
    });
    console.log(token.token);
    if (token.token == 'invalid password') {
      setError('Invalid password');
    } else if (token.token == 'no user') {
      setError('Username not found');
    } else {
      setToken(token);
    }
  };

  if (token) {
    window.location.href = '/';
  }

  return (
    <>
      <div>
        <h2 className='mt-6 text-center text-3xl font-extrabold text-gray-900'>
          Log in to your account
        </h2>
      </div>
      {error && (
        <div className='text-red-700 px-4 py-3 rounded relative' role='alert'>
          <strong className='font-bold'>Error: </strong>
          <div className='sm:inline'>{error}</div>
        </div>
      )}
      <form onSubmit={handleSubmit} className='mt-8 space-y-6'>
        <input type='hidden' name='remember' value='true' />
        <div>
          <label htmlFor='username' className='sr-only'>
            Username:
          </label>
          <input
            id='username'
            name='username'
            type='username'
            autoComplete='username'
            required
            placeholder='Username'
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div>
          <label htmlFor='password' className='sr-only'>
            Password:
          </label>
          <input
            id='password'
            name='password'
            type='password'
            autoComplete='current-password'
            required
            placeholder='Password'
            value={prePassword}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>

        <div className='text-sm'>
          <a href='#' className='font-medium text-primary-600 hover:text-primary-500'>
            Forgot your password?
          </a>
        </div>

        <Link href='/createAccount'>
          <p>Create Account</p>
        </Link>

        <div>
          <button type='submit'>Sign in</button>
        </div>
      </form>
    </>
  );
}

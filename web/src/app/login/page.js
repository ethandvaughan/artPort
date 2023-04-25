'use client';
import Footer from 'components/footer';
import Header from 'components/header';
import sha256 from 'sha256';
import { useState } from 'react';

export default function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    const hashedPassword = sha256(password);

    const response = await fetch('/api/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, hashedPassword }),
    });

    if (response.ok) {
      // Redirect to home page
      window.location.href = '/';
    } else {
      // Show error message
      const { message } = await response.json();
      alert(message);
    }
  };

  return (
    <>
      <div>
        <h2 className='mt-6 text-center text-3xl font-extrabold text-gray-900'>
          Log in to your account
        </h2>
      </div>
      <form onSubmit={handleSubmit} className='mt-8 space-y-6'>
        <input type='hidden' name='remember' value='true' />
        <div>
          <label htmlFor='email-address' className='sr-only'>
            Email address:
          </label>
          <input
            id='email-address'
            name='email'
            type='email'
            autoComplete='email'
            required
            placeholder='Email address'
            value={email}
            onChange={(e) => setEmail(e.target.value)}
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
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>

        <label htmlFor='remember-me' className='ml-2 block text-sm text-gray-900'>
          Remember me
        </label>
        <input id='remember-me' name='remember-me' type='checkbox' />

        <div className='text-sm'>
          <a href='#' className='font-medium text-primary-600 hover:text-primary-500'>
            Forgot your password?
          </a>
        </div>

        <div>
          <button type='submit'>Sign in</button>
        </div>
      </form>
    </>
  );
}

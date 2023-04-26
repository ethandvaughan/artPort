'use client';
import Link from 'next/link';
import styles from './header.module.css';
import useToken from 'components/useToken';
import HeadLogin from './headLogin';
import LogoutButton from './logoutButton';

const Header = ({ title }) => {
  const { token, setToken } = useToken();
  return (
    <header className={styles.header}>
      <div className='flex w-full flex-wrap items-center justify-between px-6'>
        <div className='flex items-center'>
          <Link href='/'>
            <h1 className={`${styles.title} justify-start`}>{title}</h1>
          </Link>
          <Link href='/ceramic'>
            <img
              src='./ceramic.png'
              className='ml-4'
              style={{ height: '30px', width: '30px' }}
              alt='Navigation: Ceramic Bowl'
              loading='lazy'
            />
          </Link>
        </div>
        <div className='relative'>
          {token ? (
            <>
              <Link href='/account'>
                <img
                  src='/profile.jpg'
                  className='rounded-full'
                  style={{ height: '30px', width: '30px' }}
                  alt=''
                  loading='lazy'
                />
              </Link>
              <LogoutButton />
            </>
          ) : (
            <HeadLogin />
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;

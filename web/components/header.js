import Link from 'next/link';
import styles from './Header.module.css';

const Header = ({ title }) => {
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
          <Link href='/account'>
            <img
              src='/profile.jpg'
              className='rounded-full'
              style={{ height: '30px', width: '30px' }}
              alt=''
              loading='lazy'
            />
          </Link>
        </div>
      </div>
    </header>
  );
};

export default Header;

import Link from 'next/link';
import styles from './Header.module.css';

const Header = ({ title }) => {
  return (
    <header className={styles.header}>
      <div className='flex w-full flex-wrap items-center justify-between px-6'>
        <Link href='/'>
          <h1 className={`${styles.title} justify-start`}>{title}</h1>
        </Link>
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

import styles from './Header.module.css';

const Header = ({ title }) => {
  return (
    <header className={styles.header}>
      <div className='flex w-full flex-wrap items-center justify-between px-6'>
        <h1 className={`${styles.title} justify-start`}>{title}</h1>
        <div className='relative'>
          <img
            src='/profile.jpg'
            className='rounded-full'
            style={{ height: '30px', width: '30px' }}
            alt=''
            loading='lazy'
          />
        </div>
      </div>
    </header>
  );
};

export default Header;

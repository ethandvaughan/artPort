import styles from './Header.module.css';

const Header = ({ title }) => {
  return (
    <header className={styles.header}>
      <h1 className={`${styles.title} justify-start`}>{title}</h1>
      <div className='relative'>
        <img
          src='../public/profile.jpg'
          class='rounded-full'
          style={{ height: '25px', width: '25px' }}
          alt=''
          loading='lazy'
        />
      </div>
    </header>
  );
};

export default Header;

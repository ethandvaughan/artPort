import styles from './Header.module.css';

const Header = ({ title }) => {
  return (
    <header className={styles.header}>
      <nav className={styles.nav}>{/* Navigation links */}</nav>
      <h1 className={styles.title}>{title}</h1>
    </header>
  );
};

export default Header;

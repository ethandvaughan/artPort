import Link from 'next/link';
import styles from './Footer.module.css';

const Footer = () => {
  return (
    <footer className={`${styles.footer} absolute inset-x-0 bottom-0`}>
      <nav>
        <ul>
          <li>
            <Link href='/about'>
              <p className={styles.footer__link}>About</p>
            </Link>
          </li>
          <li>
            <Link href='/contact'>
              <p className={styles.footer__link}>Contact</p>
            </Link>
          </li>
        </ul>
      </nav>
      <div className={styles.footer__text}>© 2023 Arfol</div>
    </footer>
  );
};

export default Footer;

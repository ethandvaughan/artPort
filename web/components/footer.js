import Link from 'next/link';
import styles from './Footer.module.css';

const Footer = () => {
  return (
    <footer className={styles.footer}>
      <nav>
        <ul>
          <li>
            <Link href="/about">
              <p className="footer__link">About</p>
            </Link>
          </li>
          <li>
            <Link href="/contact">
              <p className="footer__link">Contact</p>
            </Link>
          </li>
        </ul>
      </nav>
      <div className='footer__text'>© 2023 Arfol</div>
    </footer>
  );
};

export default Footer;
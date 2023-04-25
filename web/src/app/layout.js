import Header from 'components/header';
import './globals.css';
import Footer from 'components/footer';

export default function RootLayout({ children }) {
  return (
    <html lang='en'>
      {/*
        <head /> will contain the components returned by the nearest parent
        head.js. Find out more at https://beta.nextjs.org/docs/api-reference/file-conventions/head
      */}
      <head />
      <body>
        <div>
          <Header title='Arfol' />
          {children}
          <Footer />
        </div>
      </body>
    </html>
  );
}

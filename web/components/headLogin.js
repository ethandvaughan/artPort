import Link from 'next/link';

const HeadLogin = () => {
  return (
    <Link href='/login'>
      <p className='bold'>Login</p>
    </Link>
  );
};

export default HeadLogin;

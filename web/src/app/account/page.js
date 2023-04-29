'use client';
import AccountDetails from 'components/accountDetails';
import useToken from 'components/useToken';

export default function Account() {
  const { token, setToken } = useToken();
  if (!token) {
    window.location.href = '/login';
  }
  return (
    <>
      <AccountDetails />
    </>
  );
}

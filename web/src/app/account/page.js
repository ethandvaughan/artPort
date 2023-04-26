import AccountDetails from 'components/accountDetails';

export default function Account() {
  if (!token) {
    window.location.href = '/login';
  }
  return (
    <>
      <AccountDetails />
    </>
  );
}

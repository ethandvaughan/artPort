function handleLogout() {
  localStorage.removeItem('token');
  window.location.href = '/';
}

const LogoutButton = () => {
  return <button onClick={handleLogout}>Logout</button>;
};

export default LogoutButton;

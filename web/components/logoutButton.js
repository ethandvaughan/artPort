function handleLogout() {
  localStorage.removeItem('token');
  localStorage.removeItem('user_id');
  window.location.href = '/';
}

const LogoutButton = () => {
  return <button onClick={handleLogout}>Logout</button>;
};

export default LogoutButton;

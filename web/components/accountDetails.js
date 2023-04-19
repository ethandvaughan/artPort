const AccountDetails = () => {
  return (
    <div className='bg-gray-100 min-h-screen'>
      <div className='container mx-auto py-10'>
        <h1 className='text-2xl font-bold mb-5'>My Account</h1>
        <div className='grid grid-cols-2 gap-4'>
          <div className='bg-white p-4 rounded-lg shadow-md'>
            <h2 className='text-lg font-medium mb-4'>Profile Information</h2>
            <form>
              <div className='mb-4'>
                <label className='block text-gray-700 font-bold mb-2' htmlFor='name'>
                  Name
                </label>
                <input
                  className='appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline'
                  id='name'
                  type='text'
                  placeholder='Your Name'
                />
              </div>
              <div className='mb-4'>
                <label className='block text-gray-700 font-bold mb-2' htmlFor='email'>
                  Email
                </label>
                <input
                  className='appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline'
                  id='email'
                  type='email'
                  placeholder='you@example.com'
                />
              </div>
              <button
                className='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline'
                type='button'
              >
                Update Profile
              </button>
            </form>
          </div>
          <div className='bg-white p-4 rounded-lg shadow-md'>
            <h2 className='text-lg font-medium mb-4'>Change Password</h2>
            <form>
              <div className='mb-4'>
                <label className='block text-gray-700 font-bold mb-2' htmlFor='current-password'>
                  Current Password
                </label>
                <input
                  className='appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline'
                  id='current-password'
                  type='password'
                  placeholder='**********'
                />
              </div>
              <div className='mb-4'>
                <label className='block text-gray-700 font-bold mb-2' htmlFor='new-password'>
                  New Password
                </label>
                <input
                  className='appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline'
                  id='new-password'
                  type='password'
                  placeholder='**********'
                />
              </div>
              <button
                className='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline'
                type='button'
              >
                Change Password
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default AccountDetails;

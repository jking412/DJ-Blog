import request from './requests/request';

const App = () => {

  const register = () => {
    request.register()
  }

  const login = () => {
    request.login()
  }

  const logout = () => {
    request.logout()
  }

  return (
    <div>
      <button onClick={register}>register</button>
      <div></div>
      <button onClick={login}>login</button>
      <div></div>
      <button onClick={logout}>logout</button>
    </div>
  );
}

export default App;

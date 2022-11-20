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

  const articleCreate = () => {
    request.articleCreate()
  }

  const articleDelete = () => {
    request.articleDelete()
  }

  const articleIndex = () => {
    request.articleIndex()
  }

  const articleUpdate = () => {
    request.articleUpdate()
  }

  const articleDetail = () => {
    request.articleDetail()
  }

  const articleOrderByTime = () => {
    request.articleOrderByTime()
  }

  const articleByTagId = () => {
    request.articleByTagId()
  }

  return (
    <div>
      <button onClick={register}>register</button>
      <div></div>
      <button onClick={login}>login</button>
      <div></div>
      <button onClick={logout}>logout</button>
      <div></div>
      <button onClick={articleCreate}>articleCreate</button>
      <div></div>
      <button onClick={articleDelete}>articleDelete</button>
      <div></div>
      <button onClick={articleIndex}>articleIndex</button>
      <div></div>
      <button onClick={articleUpdate}>articleUpdate</button>
      <div></div>
      <button onClick={articleDetail}>articleDetail</button>
      <div></div>
      <button onClick={articleOrderByTime}>articleOrderByTime</button>
      <div></div>
      <button onClick={articleByTagId}>articleByTagId</button>
    </div>
  );
}

export default App;

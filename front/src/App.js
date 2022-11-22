import { useState } from 'react';
import request from './requests/request';

const App = () => {

  const [url, setUrl] = useState('')

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

  const articleByCategoryId = () => {
    request.articleByCategoryId()
  }

  const showTags = () => {
    request.showTags()
  }

  const showCategories = () => {
    request.showCategories()
  }

  const search = () => {
    request.search()
  }

  const getAvatar = () => {
    request.getAvatar().then(resp => {
      console.log(resp)
    let array  = new  Uint8Array(resp.data)
    console.log(array)
    let str12 = arrayBufferToBase64(array);//转换字符串
    console.log(str12);
    setUrl('data:image/png;base64,'+str12)
    function arrayBufferToBase64( buffer ) {
        var binary = '';
        var bytes = new Uint8Array( buffer );
        var len = bytes.byteLength;
        for (var i = 0; i < len; i++) {
            binary += String.fromCharCode( bytes[ i ] );
        }
        return window.btoa( binary );
    }
  })
}


  const uploadAvatar = () => {
    const file = document.getElementById('file').files[0];
    console.log(file)
    request.uploadAvatar(file)
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
      <div></div>
      <button onClick={articleByCategoryId}>articleByCategoryId</button>
      <div></div>
      <button onClick={showTags}>showTags</button>
      <div></div>
      <button onClick={showCategories}>showCategories</button>
      <div></div>
      <button onClick={search}>search</button>
      <div></div>
      <button onClick={getAvatar}>getAvatar</button>
      <div>
        <input type="file" id="file"/>
        <img src={url}/>
      </div>
      <div></div>
      <button onClick={uploadAvatar}>uploadAvatar</button>
    </div>
  );
}

export default App;

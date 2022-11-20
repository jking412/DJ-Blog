import axios from 'axios'

const baseUrl = 'http://localhost:8000'
axios.defaults.withCredentials = true

const register = () => {
    console.log('register')
    axios.post(`${baseUrl}/user/register`, {
        username: 'test',
        password: '123456'
    }).then(resp => {
        console.log(resp)
    }).catch(err => {
        console.log(err)
    })

}

const login = () => {
    console.log('login')
    axios.post(`${baseUrl}/user/login`, {
        username: 'test',
        password: '123456'
    }).then(resp => {
        console.log(resp)
    })
}

const logout = () => {
    console.log('logout')
    axios.get(`${baseUrl}/user/logout`, {}).then(resp => {
        console.log(resp)
    })
}

const articleCreate = () => {
    axios.post(`${baseUrl}/article`, {
        title: 'test2',
        originContent: 'test',
        parseContent: 'test',
        tags: [3],
        Categories: [1,2]
    }).then(resp => {
        console.log(resp)
    })
}

const articleDelete = () => {
    axios.delete(`${baseUrl}/article/2`, {}).then(resp => {
        console.log(resp)
    }).then(resp => {
        console.log(resp)
    })
}

const articleIndex = () => {
    axios.get(`${baseUrl}/article`, {
        params: {
            pageNum: 1,
            pageSize: 10
        }
    }).then(resp => {
        console.log(resp)
    })
}

const articleUpdate = () => {
    axios.put(`${baseUrl}/article/4`, {
        title: 'test',
        originContent: 'test',
        parseContent: 'test',
        tags: [1,2],
        Categories: [1,2]
    }).then(resp => {
        console.log(resp)
    })
}

const articleDetail = () => {
    axios.get(`${baseUrl}/article/4`, {}).then(resp => {
        console.log(resp)
    })
}

const articleOrderByTime = () => {
    axios.get(`${baseUrl}/article/time`, {
        params: {
            pageNum: 1,
            pageSize: 10
        }
    }).then(resp => {
        console.log(resp)
    })
}

const articleByTagId = () => {
    axios.get(`${baseUrl}/article/tag/3`, {
        params: {
            pageNum: 1,
            pageSize: 10
        }
    }).then(resp => {
        console.log(resp)
    })

}

const articleByCategoryId = () => {
    axios.get(`${baseUrl}/article/category/1`, {
        params: {
            pageNum: 1,
            pageSize: 10
        }
    }).then(resp => {
        console.log(resp)
    })
}

const showTags = () => {
    axios.get(`${baseUrl}/tag`, {}).then(resp => {
        console.log(resp)
    })
}

const showCategories = () => {
    axios.get(`${baseUrl}/category`, {}).then(resp => {
        console.log(resp)
    })
}

const search = () => {
    axios.get(`${baseUrl}/article/search`, {
        params: {
            pageNum: 1,
            pageSize: 10,
            keyword: 'test2'
        }
    }).then(resp => {
        console.log(resp)
    })
}



export default {
    register,
    login,
    logout,
    articleCreate,
    articleDelete,
    articleIndex,
    articleUpdate,
    articleDetail,
    articleOrderByTime,
    articleByTagId,
    articleByCategoryId,
    showTags,
    showCategories,
    search
}
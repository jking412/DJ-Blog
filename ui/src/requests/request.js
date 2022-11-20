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


export default {
    register,
    login,
    logout
}
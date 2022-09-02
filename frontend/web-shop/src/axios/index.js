import axios from 'axios'

const axiosInstance = axios.create({
})
axiosInstance.defaults.baseURL = process.env.baseURL || process.env.apiUrl || 'http://localhost:8084';

axiosInstance.interceptors.request.use(request => {
        console.log('HERERER')
        // add auth header with jwt if account is logged in and request is to the api url
        if (!request.url.includes('login')){
            let token = localStorage.getItem('api-token')
            request.headers.common.Authorization = `Bearer ${token}`;
             console.log(token)
        }
       

       
        request.headers.common.ContentType = "application/json";

        

        return request;
    });

export default axiosInstance;
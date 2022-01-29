import ky from 'ky';

const baseApi = ky.create({ prefixUrl: 'http://0.0.0.0:9000/api' });

class Api {
    async getRequest(apiUrl, params) {
        const data = new URLSearchParams(params);
        const url = data ? `${apiUrl}?${data}` : `${apiUrl}`;
        try {
            return await baseApi.get(url);
        } catch (error) {
            if (error.name === 'TypeError') {
                return 'Failed fetch';
            } else {
                return error.response;
            }
        }
    }

    async postRequest(apiUrl, params) {
        try {
            return await baseApi.post(apiUrl, params);
        } catch (error) {
            return error.response
        }
    }
}

const api = new Api();
export default api;

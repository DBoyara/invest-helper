import ky from 'ky';

const baseApi = ky.create({ prefixUrl: 'http://0.0.0.0:9000/api' });
const baseUrl = 'http://0.0.0.0:9000/api';

class Api {
    async getRequest(apiUrl, params) {
        const data = new URLSearchParams(params);
        const url = params ? `${apiUrl}?${data}` : `${apiUrl}`;
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
            return await await fetch(`${baseUrl}/${apiUrl}`, {
                method: 'POST',
                body: JSON.stringify(params),
                headers: {
                    'content-type': 'application/json',
                },
            });
        } catch (error) {
            return error.response;
        }
    }

    async putRequest(apiUrl, params) {
        try {
            return await await fetch(`${baseUrl}/${apiUrl}`, {
                method: 'PUT',
                body: JSON.stringify(params),
                headers: {
                    'content-type': 'application/json',
                },
            });
        } catch (error) {
            return error.response;
        }
    }
}

const api = new Api();
export default api;

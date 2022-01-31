import api from '../api/Api';

const baseUrl = 'pass-gen';

class PassGenApi {
    async getPass(params) {
        return await api.getRequest(baseUrl, params);
    }
}

const passGenApi = new PassGenApi();
export default passGenApi;

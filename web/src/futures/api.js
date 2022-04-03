import api from '../api/Api';

const baseUrl = 'futures';

class FuturesApi {
    async createFuturesRecord(params) {
        return await api.postRequest(baseUrl, params);
    }

    async getFuturesList(params) {
        return await api.getRequest(baseUrl, params);
    }

    async getSummary(id) {
        const uri = `${baseUrl}/summary/${id}`;
        return await api.getRequest(uri);
    }

    async updateFutures(id, params) {
        const uri = `${baseUrl}/${id}`;
        return await api.putRequest(uri, params);
    }
}

const futuresApi = new FuturesApi();
export default futuresApi;

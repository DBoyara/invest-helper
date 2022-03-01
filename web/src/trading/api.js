import api from '../api/Api';

const baseUrl = 'trading';

class TradingLogApi {
    async createTradeRecord(params) {
        return await api.postRequest(baseUrl, params);
    }

    async getTradeRecords(params) {
        return await api.getRequest(baseUrl, params);
    }

    async getSummary(type) {
        const uri = `${baseUrl}/summary/${type}`;
        return await api.getRequest(uri);
    }

    async getCommissions() {
        const uri = `${baseUrl}/commissions`;
        return await api.getRequest(uri);
    }

    async markRecordsClosed(params) {
        return await api.putRequest(baseUrl, params);
    }
}

const tradingApi = new TradingLogApi();
export default tradingApi;

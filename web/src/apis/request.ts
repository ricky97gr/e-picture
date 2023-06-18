import TimeManager from "../util/time"

const apiRequest = {
    GetOperationLogRequest: {
        method: "GET",
        body: {},
        params: { page: 1, pageSize: 20, startTime: TimeManager.GetNowDate }
    },
    GetUserInfo: {
        method: "GET",
    }
}


export default apiRequest
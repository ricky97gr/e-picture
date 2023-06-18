const TimeManager = {
    GetNowDate: GetNowDate()
}

function GetNowDate() {
    const date = new Date()
    let month: string | number = date.getMonth() + 1
    let day: string | number = date.getDate()
    if (month <= 9) {
        month = "0" + month
    }
    if (day <= 9) {
        day = "0" + day
    }
    return date.getFullYear() + "-" + month + "-" + day + " " + date.getHours() + ":" + date.getMinutes() + ":" + date.getSeconds()
}

export default TimeManager
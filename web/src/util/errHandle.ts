import { ElMessage } from "element-plus"
import router from "../router"
import successHandle from "./successHandle"

function errorHandle(data: any, msg: string) {
    if (msg != "") {
        ElMessage.error(msg);
        return false
    }
    if (data.code != 200) {
        ElMessage.error(data.msg);
        console.log(data)
        if (data.code == 100003) {
            router.push("/login")
            return true
        }
        return true
    }
    if (data.code == 200) {
        if (msg != "") {
            successHandle(msg)
        }

        return false
    }
}

export default errorHandle
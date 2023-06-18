import { ElMessage } from "element-plus"

function successHandle(msg: any) {
    ElMessage({
        message: msg,
        type: 'success',
    })
}
export default successHandle
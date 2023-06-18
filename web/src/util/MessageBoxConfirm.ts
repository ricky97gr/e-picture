import { ElMessageBox } from 'element-plus'
const MessageBoxConfirm = function MessageBoxConfirm(msg: any, done?: any) {
    ElMessageBox.confirm(
        msg,
        {
            distinguishCancelAndClose: true,
            confirmButtonText: "确认",
            cancelButtonText: "取消",
        }

    )
        .then(() => {
            done()
        })
        .catch(() => {

        })
}

export default MessageBoxConfirm
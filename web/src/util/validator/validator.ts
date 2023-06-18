import { FormRules } from "element-plus"
import { reactive } from "vue"

const validator = {
    Login: reactive<FormRules>({
        "phone": [
            { required: true, message: "请输入用户", trigger: "blur" }
        ],
        "passwd": [
            { required: true, message: "请输入密码", trigger: "blur" }
        ]
    }),
    Bucket: reactive<FormRules>({
        "name": [
            { required: true, message: "请输入桶名称" }
        ],
        "status": [],
        "capacityMode": [],
        "totalNumber": [],
        "totalSize": []
    }),
    UpLoadPhoto: reactive<FormRules>({
        "bucketName": [
            { required: true, message: "请输入桶名称", trigger: "blur" }
        ],
    }),
    ChangePasswd: reactive<FormRules>({
        "oldPasswd": [
            { required: true, message: "请输入旧密码", trigger: "blur" }
        ],
        "newPasswd": [
            { required: true, message: "请输入新密码", trigger: "blur" }
        ],
        "newPasswdAgain": [
            { required: true, message: "请再次输入新密码", trigger: "blur" }
        ],
    })
}

export default validator
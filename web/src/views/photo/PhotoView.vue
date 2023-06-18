
<script lang="ts" setup>
import PhotoSingle from './PhotoSingle.vue';
import { ref, reactive } from 'vue'
import { Plus } from '@element-plus/icons-vue'

import type { UploadFile, UploadFiles, UploadInstance, UploadProps, UploadRequestOptions, UploadUserFile } from 'element-plus'

import MessageBoxConfirm from '../../util/MessageBoxConfirm'
import httprequest from '../../util/http';
import apiUrl from '../../apis/url';
import errorHandle from '../../util/errHandle';
import validator from '../../util/validator/validator';
import successHandle from '../../util/successHandle';



const fileList = ref<UploadUserFile[]>([
])


const photoInfo = reactive({
    bucketName: "",
    pictureName: ""
})

const dialogImageUrl = ref('')
const dialogVisible = ref(false)
const photoPreviewVisible = ref(false)
const upLoadRef = ref<UploadInstance>()

const handleRemove: UploadProps['onRemove'] = (uploadFile, uploadFiles) => {
    console.log(uploadFile, uploadFiles)
}

const handlePictureCardPreview: UploadProps['onPreview'] = (uploadFile) => {
    dialogImageUrl.value = uploadFile.url!
    photoPreviewVisible.value = true
}

const handleClose = (done: () => void) => {
    MessageBoxConfirm("确定取消上传图片吗？", done)
}

function upLoadPhoto() {
    dialogVisible.value = false
    upLoadRef.value!.submit()

}

function handleLimit(files: File[], uploadFiles: UploadUserFile[]) {
    errorHandle("", "最大上传文件数为1")
}

const uploadAction = (option: UploadRequestOptions) => {
    let params = new FormData()
    console.log(option.file)
    params.append("file", option.file)
    params.append("bucketName", photoInfo.bucketName)
    params.append("name", photoInfo.pictureName)
    return httprequest(apiUrl.PostUpload, {
        method: "POST",
        // headers: { "Content-Type": "multipart/form-data" },
        body: params
    }, true).then(data => {
        if (errorHandle(data, "")) {
            return
        }
        successHandle("上传成功")
    })

}


const photoList = ref();
photoList.value = []
function listPhoto() {
    httprequest("/photo/list", {
        method: "GET",
    }).then(data => {
        if (errorHandle(data, "")) {
            return
        }
        photoList.value = data.result

    })
}

function handleSuccess(response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) {
    fileList.value = <UploadUserFile[]>[]
    console.log("上传成功")
    listPhoto()
}

function handleError(error: Error, uploadFile: UploadFile, uploadFiles: UploadFiles) {
    console.log(error)
    console.log("上传失败")
    listPhoto()
}

listPhoto()
</script>

<template>
    <el-button @click="dialogVisible = true">上传图片</el-button>

    <el-dialog v-model="photoPreviewVisible">
        <el-image :src="dialogImageUrl" alt="Preview Image" />
    </el-dialog>

    <el-dialog v-model="dialogVisible" title="上传照片" width="40%" :before-close="handleClose">
        <el-form label-position="left" :rules=validator.UpLoadPhoto>
            <el-upload v-model:file-list="fileList" action="#" list-type="picture-card"
                :on-preview="handlePictureCardPreview" :http-request=uploadAction :on-remove="handleRemove"
                :auto-upload=false ref="upLoadRef" :on-success="handleSuccess" :on-error="handleError" :limit=1
                :on-exceed="handleLimit">
                <el-icon>
                    <Plus />
                </el-icon>
            </el-upload>
            <el-form-item label="所属桶" prop="bucketName">
                <el-input v-model="photoInfo.bucketName"></el-input>
            </el-form-item>
            <el-form-item label="照片名称">
                <el-input v-model="photoInfo.pictureName"></el-input>
            </el-form-item>

        </el-form>

        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click=upLoadPhoto>
                    上传
                </el-button>
            </span>
        </template>
    </el-dialog>
    <el-row :gutter="10">
        <el-col :span="4" v-for="photo in photoList">
            <PhotoSingle :name=photo.name :url=photo.url :uuid=photo.uuid :bucket-name=photo.bucketName></PhotoSingle>
        </el-col>
    </el-row>
</template>
  
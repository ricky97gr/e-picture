<template>
    <el-card shadow="always">


        <el-image :src=photo.url :preview-src-list=[photo.url]></el-image>
        <span style="text-align: center;">{{ photo.name }}</span>
        <br>
        <el-button>复制</el-button>
        <el-button @click=getPhotoInfo>详情</el-button>
        <el-button @click=deletePhoto>删除</el-button>
    </el-card>

    <el-dialog v-model="isShowInfoDialog">
        <el-descriptions title="照片详情" :column="1" size="large" :border=true>
            <el-descriptions-item label="照片名称">{{ photo.name }}</el-descriptions-item>
            <el-descriptions-item label="照片状态"></el-descriptions-item>
            <el-descriptions-item label="所属桶">{{ photo.bucketName }}</el-descriptions-item>
            <el-descriptions-item label="照片大小"> 30k </el-descriptions-item>
        </el-descriptions>
    </el-dialog>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import errorHandle from "../../util/errHandle";
import httprequest from "../../util/http"
import successHandle from "../../util/successHandle";

const photo = defineProps<{
    name: string
    url: string
    uuid: string
    bucketName: string
}>()

const isShowInfoDialog = ref(false)

function deletePhoto() {
    httprequest("/photo/delete", {
        method: "DELETE",
        body: {
            name: photo.name,
            uuid: photo.uuid,
            bucketName: photo.bucketName
        }
    }).then(data => {
        if (errorHandle(data, "")) {
            return
        }
        successHandle("删除照片成功")
    })
}

function getPhotoInfo() {
    isShowInfoDialog.value = true
    httprequest("/photo" + photo.uuid + "info", {
        method: "GET"
    }).then(data => {
        if (errorHandle(data, "")) {
            return
        }
    })
}
</script>
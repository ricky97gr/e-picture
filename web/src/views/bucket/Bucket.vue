<template>
    <el-button @click="isShowDrawer = true">新建桶</el-button>
    <el-table :data="bucketList" border style="width: 100%" :header-cell-style=styleString.tableHeaderStyle
        :row-style=styleString.tableRowStyle :cell-style=styleString.tableCellStyle>

        <el-table-column prop="name" label="桶名称" width="180" />
        <el-table-column prop="totalNumber" label="桶容量" width="180" />
        <el-table-column prop="ownerName" label="所有者" />
        <el-table-column label="操作" v-slot="scope">
            <el-button @click="getBucketInfo(scope.row)">详情</el-button>
            <el-button @click="beforeUpdateBucket(scope.row)">修改</el-button>
            <el-button @click="deleteBucket(scope.row)">删除</el-button>
        </el-table-column>
    </el-table>

    <el-drawer title="新建桶" v-model=isShowDrawer width="40%">
        <el-form :model="newBucket">
            <el-form-item label="桶名称" prop="name">
                <el-input v-model="newBucket.name"></el-input>
            </el-form-item>
            <el-form-item label="桶状态" prop="status">
                <el-switch v-model="newBucket.status" style="--el-switch-on-color: #13ce66;"></el-switch>
            </el-form-item>
            <el-form-item label="桶限制模式" prop="capacityMode">
                <el-radio-group v-model="newBucket.capacityMode" size="large">
                    <el-radio-button label="number">数量上限</el-radio-button>
                    <el-radio-button label="size">容量上限</el-radio-button>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="桶照片上限" prop="totalNumber">
                <el-input v-model="newBucket.totalNumber" placeholder=""></el-input>
            </el-form-item>
            <el-form-item label="桶容量上限" prop="totalSize">
                <el-input v-model="newBucket.totalSize" placeholder=""></el-input>
            </el-form-item>


        </el-form>
        <el-button @click="isShowDrawer = false">取消</el-button>
        <el-button @click=createBucket>创建</el-button>
    </el-drawer>

    <el-drawer title="更新桶" v-model=isShowUpdateDrawer width="40%">
        <el-form :model="newBucket">
            <el-form-item label="桶名称" prop="name">
                <el-input v-model="newBucket.name" :placeholder=newBucket.name :disabled="true"></el-input>
            </el-form-item>
            <el-form-item label="桶状态" prop=status>
                <el-switch v-model="newBucket.status" style="--el-switch-on-color: #13ce66;"></el-switch>
            </el-form-item>
            <el-form-item label="桶限制模式" prop="capacityMode">
                <el-radio-group v-model="newBucket.capacityMode" size="large">
                    <el-radio-button label="number">数量上限</el-radio-button>
                    <el-radio-button label="size">容量上限</el-radio-button>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="桶照片上限" prop="totalNumber">
                <el-input v-model=newBucket.totalNumber></el-input>
            </el-form-item>
            <el-form-item label="桶容量上限" prop="totalSize">
                <el-input v-model=newBucket.totalSize></el-input>
            </el-form-item>


        </el-form>
        <el-button @click="isShowUpdateDrawer = false">取消</el-button>
        <el-button>更新</el-button>
    </el-drawer>

    <el-dialog v-model="isShowInfoDialog">
        <el-descriptions title="桶详情" :column="1" size="large" border="true">
            <el-descriptions-item label="桶名称"></el-descriptions-item>
            <el-descriptions-item label="桶状态"></el-descriptions-item>
            <el-descriptions-item label="桶限制模式"></el-descriptions-item>
            <el-descriptions-item label="照片数量"></el-descriptions-item>
            <el-descriptions-item label="照片上限数量"></el-descriptions-item>
            <el-descriptions-item label="照片已使用存储"></el-descriptions-item>
        </el-descriptions>
    </el-dialog>
</template>
  
<script lang="ts" setup>
import httprequest from '../../util/http';
import { ref, reactive } from 'vue'
import styleString from "../../util/style";
import errorHandle from '../../util/errHandle';
import successHandle from '../../util/successHandle';
import validator from '../../util/validator/validator';
import type { Bucket } from '../../types/interfaces';
import BucketManager from './bucketManager';



const bucketList = ref();

const bucketManager = new (BucketManager)

// const newBucket = reactive({
//     name: "",
//     status: false,
//     totalSize: 0,
//     totalNumber: 0,
//     capacityMode: ""
// })

var newBucket = reactive<Bucket>({
    name: '',
    status: '',
    totalNumber: 0,
    totalSize: 0,
    capacityMode: ''
})

const isShowDrawer = ref(false)
const isShowUpdateDrawer = ref(false)
const isShowInfoDialog = ref(false)
bucketList.value = []

function listBucket() {
    httprequest("/bucket/list", {
        method: "GET",
    }).then(data => {
        if (errorHandle(data, "")) {
            return
        }
        bucketList.value = data.result
    })
}

function createBucket() {
    isShowDrawer.value = false
    bucketManager.createBucket(newBucket)
    return
    httprequest("/bucket/create", {
        method: "POST",
        body: {
            name: newBucket.name,
            status: "inuse",
            totalNumber: newBucket.totalNumber,
            totalSize: newBucket.totalSize,
            capacityMode: newBucket.capacityMode
        }
    }).then(data => {
        if (errorHandle(data, "")) {
            return
        }
        successHandle("创建成功")
        listBucket()
    })
}

function deleteBucket(row: any) {
    console.log(row)
    httprequest("/bucket/delete", {
        method: "DELETE",
        body: {
            name: row.name,
            id: row.id
        }
    }).then(data => {
        if (errorHandle(data, "")) {
            return
        }
        successHandle("删除成功")
        listBucket()
    })

}

function beforeUpdateBucket(row: any) {
    isShowUpdateDrawer.value = true
    newBucket.name = row.name
    newBucket.status = row.size
    newBucket.totalNumber = row.totalNumber
    newBucket.totalSize = row.totalSize
    newBucket.capacityMode = row.capacityMode

}

function getBucketInfo(row: any) {
    isShowInfoDialog.value = true
    httprequest("/bucket/" + row.name + "/info", {
        method: "GET",
    }).then(data => {
        if (errorHandle(data, "")) {
            return
        }
    })

}

listBucket()



</script>
  
<template>
    <el-card shadow="always">
        <el-tag>
            {{ bucket.type }}
        </el-tag>
        <el-tag>
            {{ bucket.isUsed }}
        </el-tag>
        <el-tag>
            图片数量:{{ bucket.usedNumber + "/" + bucket.totalNumber }}
        </el-tag>
        <el-tag>
            占用存储:{{ (bucket.usedSize / 1024 / 1024).toFixed(3) + "MB/" + bucket.totalSize }}
        </el-tag>
        <el-tag>
            {{ bucket.capacityMode }}
        </el-tag>
        <br>
        桶名称：{{ bucket.name }}
        <br>
        创建时间：{{ bucket.createTime }}
        <br>

        <el-button @click="showDialog = true">编辑</el-button>
        <el-button>启用</el-button>
        <el-button>删除</el-button>
    </el-card>

    <el-dialog v-model=showDialog title="修改桶配置">
        <el-form :model="newBucket">
            <el-form-item label="桶名称">
                <el-input v-model="newBucket.name" disabled=false :placeholder="bucket.name"></el-input>
            </el-form-item>
            <el-form-item label="桶状态">
                <el-switch v-model="bucket.status" style="--el-switch-on-color: #13ce66;" active-value="inused"
                    inactive-value="unused" change=changeSwitch></el-switch>
            </el-form-item>
            <el-form-item label="桶限制模式">
                <el-radio-group v-model="newBucket.capacityMode" size="large">
                    <el-radio-button label="number">数量上限</el-radio-button>
                    <el-radio-button label="size">容量上限</el-radio-button>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="桶照片上限">
                <el-input v-model="newBucket.totalNumber" :placeholder=bucket.totalNumber></el-input>
            </el-form-item>
            <el-form-item label="桶容量上限">
                <el-input v-model="newBucket.totalSize" :placeholder=bucket.totalSize></el-input>
            </el-form-item>


        </el-form>
        <el-button @click="showDialog = false">取消</el-button>
        <el-button @click="updateBucket">创建</el-button>
    </el-dialog>
</template>
<script lang="ts" setup>
import { ref, reactive } from 'vue';
import errorHandle from '../../util/errHandle';
import httprequest from '../../util/http';
import successHandle from '../../util/successHandle';


const bucket = defineProps<{
    type: String
    isUsed: Boolean
    capacityMode: String
    createTime: String
    name: String
    usedNumber: Number
    totalNumber: Number
    usedSize: Number
    totalSize: Number
    status: string
}>()

const newBucket = reactive({
    name: "",
    status: bucket.status,
    totalSize: "",
    totalNumber: "",
    capacityMode: ""
})

const showDialog = ref()
showDialog.value = false

function updateBucket() {
    showDialog.value = false
    httprequest("/bucket/update", {
        method: "POST",
        body: {
            name: newBucket.name,
            status: "inuse",
            totalNumber: parseInt(newBucket.totalNumber, 10),
            totalSize: parseInt(newBucket.totalSize, 10),
            capacityMode: newBucket.capacityMode
        }
    }).then(data => {
        if (errorHandle(data, "")) {
            return
        }
        successHandle("更新成功")
    })
}
function changeSwitch(val: any) {
    newBucket.status = val
}

</script>
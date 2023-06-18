<template>
    <el-row :gutter="5" style="height:19%">
        <el-col :span="6">
            <el-card shadow="always">
                <el-statistic :value="DashBoardInfo.photoCount">
                    <template #title>
                        <div style="display: inline-flex; align-items: center">
                            照片数量
                            <el-tooltip effect="dark" content="Number of users who logged into the product in one day"
                                placement="top">
                                <el-icon style="margin-left: 4px" :size="12">
                                    <Warning />
                                </el-icon>
                            </el-tooltip>
                        </div>

                    </template>
                </el-statistic>

                <div class="statistic-footer">
                    <div class="footer-item">
                        <span>比上周增加</span>
                        <span class="green">
                            24%
                            <el-icon>
                                <CaretTop />
                            </el-icon>
                        </span>
                    </div>
                </div>
            </el-card>
        </el-col>
        <el-col :span="6">
            <el-card shadow="always">
                <el-statistic :value="DashBoardInfo.bucketCount">
                    <template #title>
                        <div style="display: inline-flex; align-items: center">
                            存储桶数量
                            <el-tooltip effect="dark" content="Number of users who logged into the product in one day"
                                placement="top">
                                <el-icon style="margin-left: 4px" :size="12">
                                    <Warning />
                                </el-icon>
                            </el-tooltip>
                        </div>

                    </template>
                </el-statistic>

                <div class="statistic-footer">
                    <div class="footer-item">
                        <span>比上周增加</span>
                        <span class="green">
                            100%
                            <el-icon>
                                <CaretTop />
                            </el-icon>
                        </span>
                    </div>
                </div>
            </el-card>
        </el-col>
        <el-col :span="6">
            <el-card shadow="always">
                <template #header>
                    <div class="card-header">
                        <span style="text-align:center">占用总存储量(M)</span>
                    </div>
                </template>
                <h3 style="text-align:center">{{ (DashBoardInfo.bucketUsedSize / 1024 / 1024).toFixed(3) + "MB" }}</h3>
            </el-card>
        </el-col>
        <el-col :span="6">
            <el-card shadow="always">
                待定
            </el-card>
        </el-col>
    </el-row>
    <el-row :gutter="5" style="height:30%;margin-top:10px">
        <el-col :span="24">
            <el-card shadow="always">
                系统使用率
                <!-- <useRate></useRate> -->
            </el-card>
        </el-col>
    </el-row>
    <el-row :gutter="5" style="height:48%;margin-top:10px">
        <el-col :span="8">
            <el-card shadow="always">
                存储桶情况
            </el-card>
        </el-col>
        <el-col :span="8">
            <el-card shadow="always">
                <template #header>
                    <div class="card-header">
                        <span>相片情况</span>
                    </div>
                </template>
            </el-card>
        </el-col>
        <el-col :span="8">
            <el-card shadow="always">
                <template #header>
                    <div class="card-header">
                        <span>近期动态</span>
                    </div>
                </template>
                <h5 v-for="data in DashBoardInfo.operationLogInfo.slice(0, 10)" :key="data.uuid" style="margin: 5px;">{{
                    data.userName + ' ' +
                    data.createTime + ' ' +
                    data.message
                }}</h5>
            </el-card>
        </el-col>
    </el-row>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue';
import apiRequest from '../../apis/request';
import apiUrl from '../../apis/url';
import errorHandle from '../../util/errHandle';
import httprequest from '../../util/http';

const tableData = ref();
const DashBoardInfo = reactive({
    photoCount: 0,
    bucketCount: 0,
    operationLogInfo: [],
    bucketUsedSize: 0

})
tableData.value = []
function getOperationLog() {
    httprequest(apiUrl.GetOperationLog, apiRequest.GetOperationLogRequest).then(data => {
        if (errorHandle(data, "")) {
            return
        }
        DashBoardInfo.operationLogInfo = data.result

    })


}

function getBucketCount() {
    httprequest("/bucket/count", {
        method: "GET"
    }).then(data => {
        if (errorHandle(data, "")) {
            return
        }
        DashBoardInfo.bucketCount = data.total
    })
}

function getPhotoCount() {
    httprequest("/photo/count", {
        method: "GET"
    }).then(data => {
        if (errorHandle(data, "")) {
            return
        }
        DashBoardInfo.photoCount = data.total
    })
}
function getBucketUsedSize() {
    httprequest("/bucket/usedSize", {
        method: "GET"
    }).then(data => {
        if (errorHandle(data, "")) {
            return
        }
        DashBoardInfo.bucketUsedSize = data.total
    })
}

getBucketCount()
getPhotoCount()
getOperationLog()
getBucketUsedSize()

</script>

<style scoped>
.statistic-footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 12px;
    color: var(--el-text-color-regular);
    margin-top: 16px;
}

.green {
    color: var(--el-color-success);
}
</style>
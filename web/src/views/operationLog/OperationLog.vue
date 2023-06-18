<template>
    <el-table
        v-bind:data="tableData.slice((pageInfo.currentPage - 1) * pageInfo.pageSize, pageInfo.currentPage * pageInfo.pageSize)"
        border style="width:100%" :header-cell-style=styleString.tableHeaderStyle :row-style=styleString.tableRowStyle
        :cell-style=styleString.tableCellStyle>
        <el-table-column prop="createTime" label="操作时间" />
        <el-table-column prop="userName" label="用户" width="180" />
        <el-table-column prop="module" label="模块" width="180" />
        <el-table-column prop="message" label="操作内容" />
        <el-table-column prop="ip" label="IP" />

        <el-table-column prop="handle" label="按钮">
            <el-button>删除</el-button>
        </el-table-column>
    </el-table>
    <Pagination :page-size=pageInfo.pageSize :current-page=pageInfo.currentPage :data-length=pageInfo.dataLength
        v-bind:handle-current-change="handleCurrentChange" v-bind:handle-size-change="handleSizeChange"></Pagination>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue"
import httprequest from '../../util/http';
import Pagination from '../../common/Pagination.vue'
import apiUrl from '../../apis/url'
import apiRequest from '../../apis/request'
import errorHandle from "../../util/errHandle";
import styleString from "../../util/style";

const pageInfo = reactive({
    pageSize: 20,
    dataLength: 0,
    currentPage: 1
})

const tableData = ref();
tableData.value = []

function getOperationLog() {
    httprequest(apiUrl.GetOperationLog, apiRequest.GetOperationLogRequest).then(data => {
        if (errorHandle(data, "")) {
            return
        }
        tableData.value = data.result
        pageInfo.dataLength = data.total
    })

}

function handleSizeChange(val: any) {
    pageInfo.pageSize = val
}
function handleCurrentChange(val: any) {
    pageInfo.currentPage = val
}

function deleteOperation(row: any) {

}

getOperationLog()
</script>
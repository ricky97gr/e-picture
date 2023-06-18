<template>
    <el-pagination @size-change="page.handleSizeChange" @current-change="page.handleCurrentChange"
        layout="sizes, prev, pager, next" :page-sizes="[5, 10, 20, 50]" :total=page.dataLength :page-size=page.pageSize
        :current-page=page.currentPage class="manager-pagination">
    </el-pagination>
</template>

<script lang="ts" setup>


const page = defineProps({
    pageSize: Number,
    pageTotal: Number,
    currentPage: Number,
    data: Array,
    dataLength: Number,
    handleCurrentChange: function(){},
    handleSizeChange:function(){},
})

</script>

<style scoped>
.manager-pagination {
    margin-top: 15px;
    float: right;
    margin-bottom: 10px;
    padding-bottom: 30px;
}
</style>
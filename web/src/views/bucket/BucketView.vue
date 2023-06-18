
<script lang="ts" setup>
import BucketSingle from './BucketSingle.vue';
import httprequest from '../../util/http';
import { ref, reactive } from 'vue'
import errorHandle from '../../util/errHandle';

const bucketList = ref();
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

listBucket()
</script>


<template >
    <el-row :gutter="10">

        <el-col :span="6" v-for="bucket in bucketList" :key="bucket.id">
            <BucketSingle  :createTime="bucket.createTime"  :isUsed=true
                type="本地存储" :name="bucket.name" :capacityMode="bucket.capacityMode" :usedNumber="bucket.usedNumber" :totalNumber="bucket.totalNumber" :totalSize="bucket.totalSize" :usedSize="bucket.usedSize">
            </BucketSingle>
        </el-col>


    </el-row>
</template>

  
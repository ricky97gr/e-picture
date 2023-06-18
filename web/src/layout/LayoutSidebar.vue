<template>
    <el-row class="tac">
        <el-col :span="24">
            <el-menu default-active="2" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose" :router=true
                background-color="#666666" text-color="white">

                <!-- <el-menu-item index="test" route="/test">
                    <el-icon><icon-menu /></el-icon>
                    <span>测试</span>
                </el-menu-item> -->


                <el-menu-item index="home" route="/home">
                    <el-icon><icon-menu /></el-icon>
                    <span>首页</span>
                </el-menu-item>
                <el-menu-item index="photoView" route="/photoView">
                    <el-icon>
                        <document />
                    </el-icon>
                    <span>照片墙</span>
                </el-menu-item>
                <el-menu-item index="bucketView" route="/bucketView">
                    <el-icon>
                        <document />
                    </el-icon>
                    <span>存储桶</span>
                </el-menu-item>
                <el-menu-item index="operationLog" route="/operationLog">
                    <el-icon>
                        <setting />
                    </el-icon>
                    <span>操作日志</span>
                </el-menu-item>

                <el-sub-menu index="1">
                    <template #title>
                        <el-icon>
                            <location />
                        </el-icon>
                        <span>系统</span>
                    </template>

                    <el-menu-item index="user" disabled>用户管理</el-menu-item>
                    <el-menu-item index="setting" route="/setting">个人信息</el-menu-item>

                    <el-menu-item index="1-3" disabled>系统设置</el-menu-item>
                    <el-menu-item index="bucket" route="/bucket">桶管理</el-menu-item>
                    <el-menu-item index="uplog" route="/uplog">更新日志</el-menu-item>

                </el-sub-menu>
            </el-menu>
        </el-col>

    </el-row>
</template>
  
<script lang="ts" setup>
import {
    Document,
    Menu as IconMenu,
    Location,
    Setting,
} from '@element-plus/icons-vue'
const handleOpen = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
</script>
  
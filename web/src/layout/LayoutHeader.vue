<template>
    <el-dropdown trigger="click" class="header-avator" size="large" @command="handleCommand">
        <el-avatar :size="40" fit="cover" class="avator" src="http://minio.test.com:9000/test/login.jpg">
        </el-avatar>
        <template #dropdown>
            <el-dropdown-menu>
                <el-dropdown-item command="person">个人中心</el-dropdown-item>
                <!-- <el-dropdown-item>修改密码</el-dropdown-item> -->
                <el-dropdown-item command="logout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
        </template>
    </el-dropdown>
</template>

<style scoped>
.header-avator {
    float: right;
    margin-top: 10px;
    margin-bottom: 10px;
}
</style>

<script lang="ts" setup>
import router from '../router';
import httprequest from '../util/http';

const handleCommand = (command: string) => {
    switch (command) {
        case "person":
            router.push("/setting")
            return
        case "logout":
            httprequest("/user/logout", {
                method: "GET"
            })
            sessionStorage.removeItem("token")
            router.push("/login")
            return
    }
}
</script>
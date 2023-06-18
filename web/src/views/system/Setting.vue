<script lang="ts" setup>
import apiRequest from '../../apis/request';
import apiUrl from '../../apis/url';
import httprequest from '../../util/http';
import { reactive } from 'vue'
import errorHandle from '../../util/errHandle';
import getResult from '../../util/getResult';
import validator from '../../util/validator/validator'

const userInfo = reactive({
  userName: String,
  email: String,
  phone: String,
  uuid: String,
  avatar: String,
})

function getUserInfo() {
  httprequest(apiUrl.GetUserInfo, apiRequest.GetUserInfo).then(data => {
    if (errorHandle(data, "")) {
      return
    }
    userInfo.email = getResult(data).email
    userInfo.phone = getResult(data).phone
    userInfo.userName = getResult(data).userName
    userInfo.avatar = getResult(data).avatar
  })
}

getUserInfo()

</script>

<template>
  <el-row :gutter="20">
    <el-col :span="8">
      <el-card>
        <template #header>
          <div class="card-header">
            <!-- <span>用户详情</span> -->
            <el-image :src="userInfo.avatar"></el-image>
          </div>
        </template>
        <el-descriptions title="用户详情" :column="1" size="large" :border="true">
          <el-descriptions-item label="用户名称">{{ userInfo.userName }}</el-descriptions-item>
          <el-descriptions-item label="用户手机">{{ userInfo.phone }}</el-descriptions-item>
          <el-descriptions-item label="用户邮箱">{{ userInfo.email }}</el-descriptions-item>
          <el-descriptions-item label="用户角色"></el-descriptions-item>
        </el-descriptions>
      </el-card>
    </el-col>

    <el-col :span="16">
      <el-card>
        <el-tabs type="border-card">
          <el-tab-pane label="基础信息">
            <el-form>
              <el-form-item label="用户名称">
                <el-input :placeholder="userInfo.userName"></el-input>
              </el-form-item>
              <el-form-item label="用户手机">
                <el-input :placeholder="userInfo.phone"></el-input>
              </el-form-item>
              <el-form-item label="用户邮箱">
                <el-input :placeholder="userInfo.email"></el-input>
              </el-form-item>
              <el-button>保存</el-button>
            </el-form>
          </el-tab-pane>
          <el-tab-pane label="修改密码">
            <el-form :rules=validator.ChangePasswd>
              <el-form-item label="旧密码" prop="oldPasswd">
                <el-input></el-input>
              </el-form-item>
              <el-form-item label="新密码" prop="newPasswd">
                <el-input></el-input>
              </el-form-item>
              <el-form-item label="确认密码" prop="newPasswdAgain">
                <el-input></el-input>
              </el-form-item>
              <el-button>保存</el-button>
            </el-form>
          </el-tab-pane>
        </el-tabs>
      </el-card>
    </el-col>
  </el-row>
</template>
  
<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}
</style>
  

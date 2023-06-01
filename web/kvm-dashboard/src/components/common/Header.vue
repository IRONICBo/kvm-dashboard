<template>
    <el-row class="header">
      <el-col :span="6">
        <span class="span">KVM Dashboard</span>
      </el-col>
      <el-col :span="1" :offset="16">
        <el-dropdown v-model="method">
          <el-icon style="color: #337ecc;padding-top:2vh;"><Setting /></el-icon>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="jumpToConfig()"> Config </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-col>
      <el-col :span="1" >
        <el-dropdown v-model="method">
          <el-icon style="color: #337ecc;padding-top:2vh;"><User /></el-icon>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="logoutAndClean()"> Logout </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-col>
    </el-row>
</template>

<script>
import { ElNotification } from 'element-plus';

import { logout } from '@/api/user';
import { delCookie } from '@/utils/cookie';

export default {
  methods: {
    jumpToConfig() {
      this.$router.replace('/config').catch(err => {err})
    },
    logoutAndClean() {
      this.$router.replace('/login').catch(err => {err})

      // clean cookies
      logout().then(response => {
        delCookie();

        ElNotification({
          title: 'Success',
          message: 'Logout successfully',
          type: 'success'
        });
      }).catch(error => {
        console.log(error)
      })
    }
  }
}
</script>

<style>
header {
    text-align: left;
    line-height: 50px;
    background-color: #fff;

    border-bottom: 1px solid var(--el-border-color);;
}
.span {
  color: #337ecc;
  padding-left: 20px;
}
</style>

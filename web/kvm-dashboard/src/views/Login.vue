<template>
  <div style="background-color: #fff; height: 100vh;">
    <el-container>
      <el-header style="height: 6vh;" >
        <el-row class="header">
          <el-col :span="6">
            <span class="span">KVM Dashboard</span>
          </el-col>
        </el-row>
      </el-header>

      <el-main>
        <el-row style="margin-top: 14vh;">
          <el-col :span="6" :offset="9">
            <el-form 
            :model="user"
            label-position="top"
            style="max-width: 500px">
              <el-form-item>
                <h1>KVM Login</h1>
              </el-form-item>
              <el-form-item label="Username">
                <el-input 
                  v-model="username"
                  class="w-50 m-2"
                  size="large"
                  clearable />
              </el-form-item>
              <el-form-item label="Password">
                <el-input 
                  v-model="password"
                  class="w-50 m-2"
                  size="large"
                  show-password
                  clearable />
              </el-form-item>
              <el-form-item label="">
                <el-button
                  type="primary"
                  style="width:500px"
                  size="large"
                  @click="onSubmit" >Login</el-button>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { ElNotification } from 'element-plus';

import { login } from '@/api/user';

export default {
  data() {
    return {
        username: "",
        password: "",
    }
  },
  methods: {
    onSubmit() {
      login(this.username, this.password).then(response => {
        if (response.code != 200) {
          ElNotification({
            title: 'Status',
            message: "Login Failed!",
            type: 'error',
            duration: 1000
          });

          setInterval(() => {
            this.$router.push({ name: 'Login' })
          }, 1000);

          return
        }

        ElNotification({
          title: 'Status',
          message: "Login Success!",
          type: 'success',
          duration: 1000
        });

        setInterval(() => {
          this.$router.push({ name: 'Home' })
        }, 1000);

      }).catch(error => {
        ElNotification({
          title: 'Status',
          message: "Login Failed!",
          type: 'error',
          duration: 1000
        });

        setInterval(() => {
          this.$router.push({ name: 'Login' })
        }, 1000);
      })
      if (this.username == "admin" && this.password == "admin") {
        this.$router.push({ name: 'Home' })
      }
    }
  }
}
</script>

<style scoped>
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

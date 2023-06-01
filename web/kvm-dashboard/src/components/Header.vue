<template>
    <el-row class="header">
      <el-col :span="6">
        <span class="span">KVM Dashboard</span>
      </el-col>
      <el-col :span="1" :offset="15">
        <el-icon style="color: #337ecc;padding-top:2vh;"><Plus /></el-icon>  
      </el-col>
      <el-col :span="1" >
        <el-icon style="color: #337ecc;padding-top:2vh;"><Bell /></el-icon>  
      </el-col>
      <el-col :span="1" >
        <el-dropdown v-model="method">
          <el-icon style="color: #337ecc;padding-top:2vh;"><User /></el-icon>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="logout()"> Logout </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-col>
    </el-row>
</template>

<script>
export default {
  methods: {
    logout() {
      this.$router.push({ name: 'Login' })
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

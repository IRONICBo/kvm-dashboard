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
        <el-icon style="color: #337ecc;padding-top:2vh;"><User /></el-icon>  
      </el-col>
    </el-row>
</template>

<script>
export default {

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

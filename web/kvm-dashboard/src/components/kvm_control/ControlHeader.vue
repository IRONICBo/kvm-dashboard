<template>
  <el-row class="header">
    <el-col :span="6">
      <el-button
        type="text"
        style="background: white"
        @click="JumptoControl"
        class="back_button"
        ><span class="span">KVM Dashboard</span></el-button
      >
    </el-col>
  </el-row>
</template>

<script>
export default {
  methods: {
    JumptoControl() {
      this.$router.push("/");
    },
  },
};
</script>

<style>
header {
  text-align: left;
  line-height: 6.25vh;
  background-color: #fff;

  border-bottom: 1px solid var(--el-border-color);
}
.span {
  color: #337ecc;
  padding-left: 20px;
  font-size: large;
}
.back_button {
  margin-top: -2px;
  margin-left: -1px;
}
</style>

<template>
  <!-- todo: side bar -->
  <el-scrollbar class="scrollbar">
    <el-menu default-active="2" >
        <el-sub-menu index="1">
            <template #title>
                <el-icon><Monitor /></el-icon>
                <span><el-text size="large" tag="b">Host nodes</el-text></span>
            </template>
            <el-menu-item v-for="(item, idx) in fakeHostList" :key="item" :index="idx">{{ item }}</el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="2">
            <template #title>
                <el-icon><Money /></el-icon>
                <span><el-text size="large" tag="b">VM nodes</el-text></span>
            </template>
            <el-menu-item v-for="(item, idx) in fakeVmList" :key="item" :index="idx">{{ item }}</el-menu-item>
        </el-sub-menu>
    </el-menu>
  </el-scrollbar>
</template>

<script>
export default {
  data() {
    return {
      fakeHostList: [
        "Ubuntu-4080",  
        "Ubuntu-A100",  
      ],
      fakeVmList: [
        "vm1",
        "kylin10.0",
        "kylin10.0-2",
        "RemoteTest1",
        "RemoteTest2",
        "RemoteTest3",
        "gpu0",
        "gpu1",
        "gpu2",
        "gpu3",
        "centos-1",
        "centos-2",
        "centos-3",
        "centos-4",
        "centos-5",
        "node1",
        "node3",
        "node4",
        "node5",
      ]
    }
  },
  mounted() {
  }
}
</script>

<style>
.scrollbar {
    text-align: center;
    background-color: #fff;
}
</style>
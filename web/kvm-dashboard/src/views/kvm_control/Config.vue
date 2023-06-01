<template>
  <div class="common-layout">
    <el-container>
      <el-header style="height: 6vh">
        <controlHeader />
      </el-header>
      <el-container>
        <el-aside style="width: 250px" class="el-aside">
          <el-scrollbar class="scrollbar">
            <el-menu @select="menuClick" default-active="/host/manage">
              <el-menu-item index="/host/manage">
                <template #title>
                  <el-icon><Operation /></el-icon>
                  <span
                    ><el-text size="large" tag="b">宿主机管理</el-text></span
                  >
                </template>
              </el-menu-item>
              <el-menu-item index="/domain/manage">
                <template #title>
                  <el-icon><Operation /></el-icon>
                  <span
                    ><el-text size="large" tag="b">虚拟机管理</el-text></span
                  >
                </template>
              </el-menu-item>
              <el-menu-item index="/domain/create">
                <template #title>
                  <el-icon><CirclePlus /></el-icon>
                  <span
                    ><el-text size="large" tag="b">虚拟机创建</el-text></span
                  >
                </template>
              </el-menu-item>
            </el-menu>
          </el-scrollbar>
        </el-aside>

        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import { Operation } from "@element-plus/icons-vue";
import ControlHeader from "@/components/kvm_control/ControlHeader.vue";

export default {
  name: "HomeView",
  components: {
    Operation,
    ControlHeader,
  },
  methods: {
    menuClick(index) {
      this.$router.push(index);
    },
  },
};
</script>
<style scoped>
.scrollbar {
  text-align: center;
  background-color: #fff;
}
.el-aside {
  display: block;
  position: relative;
}
</style>
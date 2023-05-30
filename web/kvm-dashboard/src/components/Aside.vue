<template>
  <!-- todo: side bar -->
  <el-scrollbar class="scrollbar">
    <el-menu default-active="vm-0" >
        <el-sub-menu index="1">
            <template #title>
                <el-icon><Monitor /></el-icon>
                <span><el-text size="large" tag="b">Host nodes</el-text></span>
            </template>
            <el-menu-item 
              v-for="(item, idx) in HostList" 
              :key="item" 
              :index="'host-' + idx"
              @click="changeItem(item.uuid)" >{{ item.name }}</el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="2">
            <template #title>
                <el-icon><Money /></el-icon>
                <span><el-text size="large" tag="b">VM nodes</el-text></span>
            </template>
            <el-menu-item 
              v-for="(item, idx) in VmList" 
              :key="item" 
              :index="'vm-' + idx"
              @click="changeItem(item.uuid)" >{{ item.name }}</el-menu-item>
        </el-sub-menu>
    </el-menu>
  </el-scrollbar>
</template>

<script>
import { getMachineList } from '@/api/machine';

export default {
  emits: ['updateUUID'],
  data() {
    return {
      HostList: [],
      VmList: [],
    }
  },
  mounted() {
    this.initList();
  },
  methods: {
    initList() {
      getMachineList().then((response) => {
        // clear list
        this.HostList = [];
        this.VmList = [];
        let resp = response.data;

        resp["HostInfo"].forEach((item) => {
          this.HostList.push({
            "name": item.name,
            "uuid": item.uuid,
          })
        })

        resp["VmInfo"].forEach((item) => {
          this.VmList.push({
            "name": item.name,
            "uuid": item.uuid,
          })
        })

        // init selected uuid
        this.$emit('updateUUID', this.VmList[0].uuid);
      }).catch((error) => {
        console.log(error)
      })
    },
    changeItem(uuid) {
      this.$emit('updateUUID', uuid);
    }
  }
}
</script>

<style>
.scrollbar {
    text-align: center;
    background-color: #fff;
}
</style>
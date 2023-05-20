<template>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>System Info</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="18">
          <el-descriptions :column="6">
            <!-- Use for each -->
            <el-descriptions-item label="ID">{{ vmInfo.id }}</el-descriptions-item>
            <el-descriptions-item label="Name">{{ vmInfo.name }}</el-descriptions-item>
            <el-descriptions-item label="State">{{ vmInfo.state }}</el-descriptions-item>
            <el-descriptions-item label="CPU">{{ vmInfo.cpu }}</el-descriptions-item>
            <el-descriptions-item label="MaxMem">{{ vmInfo.max_mem }}</el-descriptions-item>
            <el-descriptions-item label="IPAddress">{{ vmInfo.ip_address }}</el-descriptions-item>
            <el-descriptions-item label="OSType">{{ vmInfo.os_type }}</el-descriptions-item>
            <el-descriptions-item label="IsPersistent">{{ vmInfo.is_persistent }}</el-descriptions-item>
            <el-descriptions-item label="AutoStart">{{ vmInfo.auto_start }}</el-descriptions-item>
            <el-descriptions-item label="UUID">{{ vmInfo.uuid }}</el-descriptions-item>
          </el-descriptions>
        </el-col>
        <el-col :span="6">
          <el-button type="primary"><el-icon><VmStart /></el-icon></el-button>
          <el-button type="primary"><el-icon><VmStop /></el-icon></el-button>
          <el-button type="primary"><el-icon><VmPause /></el-icon></el-button>
          <el-button type="primary"><el-icon><VmRestart /></el-icon></el-button>
        </el-col>
      </el-row>
    </el-card>
</template>
  
<script>
import VmPause from '@/assets/icons/VmPause';
import VmRestart from '@/assets/icons/VmRestart';
import VmStart from '@/assets/icons/VmStart';
import VmStop from '@/assets/icons/VmStop';

import {getVMInfo} from '@/api/info';
import {TEMPINFO} from '@/constant/constant';

export default {
    data() {
        return {
            hostInfo: {},
            vmInfo: {
              "id": "",
              "name": "",
              "uuid": "",
              "os_type": "",
              "state": "",
              "cpu": "",
              "max_mem": "",
              "is_persistent": "",
              "auto_start": "",
              "ip_address": "",
            },
        }
    },
    props: {
        UUID: "",
    },
    mounted() {
      this.fetchVMInfo();
    },
    methods: {
      fetchVMInfo() {
        getVMInfo(TEMPINFO.uuid)
        .then(response => {
            this.vmInfo = response.data;
            // render
        })
        .catch(error => {
            console.log(error);
        });
      }
    },
    components: {
        VmStart,
        VmStop,
        VmPause,
        VmRestart
    }
}
</script>
  
<style scoped>
.card-header {
    display: flex;
    justify-content: space-between;
    align-items: left;
}

.text {
    font-size: 14px;
}

.item {
    margin-bottom: 18px;
}

.box-card {
    margin-bottom: 14px;
}
.el-row {
    margin-bottom: 20px;
}
.el-row:last-child {
    margin-bottom: 0;
}
.el-col {
    border-radius: 4px;
}

.grid-content {
    border-radius: 4px;
    min-height: 36px;
}
</style>

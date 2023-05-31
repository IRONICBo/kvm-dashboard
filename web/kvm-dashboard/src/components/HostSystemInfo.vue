<template>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>Host System Info</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="18">
          <el-descriptions :column="3">
            <!-- Use for each -->
            <el-descriptions-item label="UUID">{{ hostInfo.uuid }}</el-descriptions-item>
            <el-descriptions-item label="Name">{{ hostInfo.name }}</el-descriptions-item>
            <el-descriptions-item label="CPU">{{ hostInfo.cpu }}</el-descriptions-item>
            <el-descriptions-item label="OsType">{{ hostInfo.os_type }}</el-descriptions-item>
            <el-descriptions-item label="IPAddress">{{ hostInfo.ip }}</el-descriptions-item>
            <el-descriptions-item label="MaxMem">{{ hostInfo.max_mem }}</el-descriptions-item>
            <el-descriptions-item label="CPUName">{{ hostInfo.cpu_name }}</el-descriptions-item>
            <el-descriptions-item label="Hostname">{{ hostInfo.hostname }}</el-descriptions-item>
            <el-descriptions-item label="SSH">{{ hostInfo.ssh_port }}</el-descriptions-item>
            <el-descriptions-item label="Libvirt">{{ hostInfo.libvirt_url }}</el-descriptions-item>
          </el-descriptions>
        </el-col>
        <el-col :span="6">
          <el-row justify="space-between">
            <el-col :span="12"><el-button @click="hostStart" type="primary"><el-icon><VmStart /></el-icon>&nbsp;Start Monitor</el-button></el-col>
            <el-col :span="12"><el-button @click="hostStop" type="primary"><el-icon><VmStop /></el-icon>&nbsp;Stop Monitor</el-button></el-col>
          </el-row>
          <div v-show="controlVis">
            <el-progress :percentage="controlStatus.percent" striped striped-flow />
            <el-descriptions :column="1">
              <el-descriptions-item label="MSG">{{ controlStatus.msg }}</el-descriptions-item>
            </el-descriptions>
          </div>
        </el-col>
      </el-row>
    </el-card>
</template>
  
<script>
import VmPause from '@/assets/icons/VmPause';
import VmRestart from '@/assets/icons/VmRestart';
import VmStart from '@/assets/icons/VmStart';
import VmStop from '@/assets/icons/VmStop';

import { getHostInfo } from '@/api/info';
import { HostStartWebSocket, HostStopWebSocket } from '@/api/control';

import { ElNotification,ElProgress } from 'element-plus';

export default {
    props: {
        uuid: "",
    },
    data() {
        return {
            hostInfo: {
              "name": "",
              "hostname": "",
              "uuid": "",
              "ip": "",
              "os_type": "",
              "cpu": "",
              "cpu_name": "",
              "max_mem": "",
              "ssh_port": "",
              "libvirt_url": ""
            },
            controlVis: false,
            controlStatus: {
              "uuid": "",
              "percent": 0,
              "msg": "",
              "state": null,
            },
        }
    },
    watch: {
      uuid: {
        immediate: true,
        handler: function (val, oldVal) {
          if (oldVal == null || val == null || val == "") {
              return;
          }
          this.fetchHostInfo();
        },
      },
    },
    mounted() {
      // for reload
      this.fetchHostInfo();
    },
    methods: {
      fetchHostInfo() {
        if (this.uuid == "") {
          return;
        }

        getHostInfo(this.uuid)
        .then(response => {
            this.hostInfo = response.data;
            // render
        })
        .catch(error => {
            console.log(error);
        });
      },
      hostStart() {
        this.initWebSocket(HostStartWebSocket);
      },
      hostStop() {
        this.initWebSocket(HostStopWebSocket);
      },
      initWebSocket(socket) {
        if (this.uuid == "") {
          return;
        }

        if (typeof WebSocket === 'undefined') {
            return console.log('Your browser doesn\'t support WebSocket');
        }
        this.ws = socket(this.uuid);
        this.ws.onmessage = this.handleMessage;
        this.ws.onopen = this.handleOpen;
        this.ws.onclose = this.handleClose;
        this.ws.onerror = this.handleError;
      },
      handleMessage(data) {
          // open progress bar
          this.controlVis = true;

          let controlResp = JSON.parse(data.data);
          console.log("handleMessage", controlResp)
          this.controlStatus = controlResp;

          if (this.controlStatus.percent == 100 && this.controlStatus.state == true) {
            ElNotification({
              title: 'Host Status',
              message: this.controlStatus.msg,
              type: 'success',
              duration: 3000
            });

            // close progress bar
            setInterval(() => {
              this.controlVis = false;
            }, 3000);
          }

          if (this.controlStatus.percent == 100 && this.controlStatus.state == false) {
            ElNotification({
              title: 'Host Status',
              message: this.controlStatus.msg,
              type: 'error',
              duration: 3000
            });

            // close progress bar
            setInterval(() => {
              this.controlVis = false;
            }, 3000);
          }
      },
      handleOpen() {
          console.log('handleOpen');
      },
      handleClose() {
          console.log('handleClose');
      },
      handleError() {
          // reconnect
          this.initWebSocket();
          console.log('handleError');
      },
      handleSend(data) {
          this.ws.send(data);
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

.control_status {
  /* height: 150px; */
  display:flex;
  align-items: center;
  justify-content: center;
}

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

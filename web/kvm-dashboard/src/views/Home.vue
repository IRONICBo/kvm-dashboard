<template>
  <div class="home">
    <el-container>
      <el-header style="height: 6vh;" >
        <Header />
      </el-header>
      <el-container>
        <el-aside style="width: 250px; height: 94vh;">
          <Aside @updateUUID="updateUUID" />
        </el-aside>
        <el-main style="height:94vh">
          <system-info 
            v-if="!isHost"
            :uuid="uuid" ></system-info>
          <host-system-info 
            v-if="isHost"
            :uuid="uuid" ></host-system-info>
          <real-time-info 
            :realtimeInfo="realtimeInfo"
            :uuid="uuid"
            :isHost="isHost"></real-time-info>
          <metric-info
            v-if="!isHost"
            :key="uuid"
            :uuid="uuid" ></metric-info>
          <history-info 
            :key="uuid"
            :realtimeInfoWithTimestamp="realtimeInfoWithTimestamp"
            :uuid="uuid" ></history-info>
          <el-row :gutter="20">
            <el-col :span="16">
              <process-info
                :uuid="uuid" ></process-info>
            </el-col>
            <el-col :span="8">
              <alert-info
                :uuid="uuid" ></alert-info>
            </el-col>
          </el-row>
        </el-main>
      </el-container>
    </el-container>
  </div> 
</template>

<script>
import RealTimeInfo from '@/components/RealTimeInfo.vue'
import HistoryInfo from '@/components/HistoryInfo.vue'
import ProcessInfo from '@/components/ProcessInfo.vue'
import AlertInfo from '@/components/AlertInfo.vue'
import Header from '@/components/Header.vue'
import Aside from '@/components/Aside.vue'
import SystemInfo from '@/components/SystemInfo.vue'
import HostSystemInfo from '@/components/HostSystemInfo.vue'
import MetricInfo from '@/components/MetricInfo.vue'

import {SimpleWebSocket} from '@/api/realtime';

export default {
  components: { 
    Header,
    Aside,
    RealTimeInfo, 
    HistoryInfo, 
    ProcessInfo,
    AlertInfo,
    SystemInfo,
    HostSystemInfo,
    MetricInfo,
  },
  data() {
    return {
      realtimeInfoWithTimestamp: {
          "name": "",
          "unit": "",
          "data": [
              {
                  "cpu_usage": 0,
                  "mem_usage": 0,
                  "disk_usage": 0,
                  "net_rx_rate": 0,
                  "net_tx_rate": 0
              }
          ],
          "timestamp": [
              0
          ]
      },
      realtimeInfo: {
          cpu_usage: 0,
          memory_usage: 0,
          disk_usage: 0,
          net_rx_rate: 0,
          net_tx_rate: 0,
      },
      uuid: "",
      isHost: false,
    }
  },
  watch: {
    uuid: {
      immediate: true,
      handler: function (val, oldVal) {
        if (oldVal == null || val == null || val == "") {
            return;
        }
        this.closeWebSocket();
        this.initWebSocket();
      },
    },
  },
  mounted() {
  },
  unmounted() {
    // if the ws is null, can not use close method.
    if (this.ws) {
        this.ws.close();
    }
  },
  methods: {
    // ws
    initWebSocket() {
        if (typeof WebSocket === 'undefined') {
            return console.log('Your browser doesn\'t support WebSocket');
        }
        this.ws = SimpleWebSocket(this.uuid);
        this.ws.onmessage = this.handleMessage;
        this.ws.onopen = this.handleOpen;
        this.ws.onclose = this.handleClose;
        this.ws.onerror = this.handleError;
    },
    handleMessage(data) {
        let Resp = JSON.parse(data.data);
        // console.log("handleMessage", realtimeResp)
        // realtimeResp = realtimeResp.data[0]
        this.realtimeInfoWithTimestamp = Resp; // :value to update Rate
        this.realtimeInfo = Resp.data[0]; // :value to update Rate

        // console.log("realtimeInfo", this.realtimeInfoWithTimestamp)
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
    },
    closeWebSocket() {
        if (this.ws) {
            this.ws.close();
        }
    },
    updateUUID(data, isHost) {
      this.uuid = data;
      this.isHost = isHost;
      console.log("updateUUID", data, isHost)
    }
  }
}
</script>

<style scoped>
</style>

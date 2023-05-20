<template>
  <div class="home">
    <el-container>
      <el-header style="height: auto;">
        <Header />
      </el-header>
      <el-container>
        <el-aside width="200px">Aside</el-aside>
        <el-main style="height:90vh">
          <system-info></system-info>
          <real-time-info :realtimeInfo="realtimeInfo"></real-time-info>
          <history-info :realtimeInfoWithTimestamp="realtimeInfoWithTimestamp"></history-info>
          <el-row :gutter="20">
            <el-col :span="16">
              <process-info></process-info>
            </el-col>
            <el-col :span="8">
              <alert-info></alert-info>
            </el-col>
          </el-row>
        </el-main>
      </el-container>
    </el-container>
  </div> 
</template>

<script>
import RealTimeInfo from '../components/RealTimeInfo.vue'
import HistoryInfo from '../components/HistoryInfo.vue'
import ProcessInfo from '../components/ProcessInfo.vue'
import AlertInfo from '../components/AlertInfo.vue'
import Header from '../components/Header.vue'
import SystemInfo from '../components/SystemInfo.vue'

import {SimpleWebSocket} from '@/api/realtime';
import {TEMPINFO} from '@/constant/constant';

export default {
  components: { 
    Header,
    RealTimeInfo, 
    HistoryInfo, 
    ProcessInfo,
    AlertInfo,
    SystemInfo,
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
    }
  },
  mounted() {
      this.initWebSocket();
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
        this.ws = SimpleWebSocket(TEMPINFO.uuid);
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
    }
  }
}
</script>

<style scoped>
</style>

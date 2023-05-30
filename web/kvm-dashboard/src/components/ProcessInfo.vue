<template>
    <el-card class="box-card" >
      <template #header>
        <div class="card-header">
          <span>Process Workload</span>
        </div>
      </template>
      <el-table align="center" :data="processData" style="width: 100%" max-height="300px">
        <el-table-column fixed prop="command" label="COMMAND" width="180" />
        <el-table-column prop="pid" sortable label="PID" width="100"  />
        <el-table-column prop="user" sortable label="USER" width="100" />
        <el-table-column prop="pr" sortable label="PR" width="100" />
        <el-table-column prop="ni" sortable label="NI" width="100" />
        <el-table-column prop="virt" sortable label="VIRT" width="100" />
        <el-table-column prop="res" sortable label="RES" width="100" />
        <el-table-column prop="shr" sortable label="SHR" width="100" />
        <el-table-column prop="s" sortable label="S" width="100" />
        <el-table-column prop="cpu" sortable label="CPU" width="100" />
        <el-table-column prop="mem" sortable label="MEM" width="100" />
        <el-table-column prop="time" sortable label="TIME" width="100" />
      </el-table>
    </el-card>
</template>
  
<script>
import {ProcessWebSocket} from '@/api/realtime';

export default {
  props: {
    uuid: "",
  },
  data() {
    return {
      ws: null,
      processData: [],
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
        // this.initWebSocket();
  },
  unmounted() {
      // if the ws is null, can not use close method.
      this.closeWebSocket();
  },
  methods: {
    // ws
    initWebSocket() {
        if (typeof WebSocket === 'undefined') {
            return console.log('Your browser doesn\'t support WebSocket');
        }
        this.ws = ProcessWebSocket(this.uuid);
        this.ws.onmessage = this.handleMessage;
        this.ws.onopen = this.handleOpen;
        this.ws.onclose = this.handleClose;
        this.ws.onerror = this.handleError;
    },
    handleMessage(data) {
        let processResp = JSON.parse(data.data);
        // console.log("handleMessage", processResp)
        this.processData = processResp;
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
        
        // clean data
        this.processData = [];
    },
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

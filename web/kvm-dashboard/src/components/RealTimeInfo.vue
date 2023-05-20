<template>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>VM Workload - Real Time</span>
          <el-button class="button" text>xxx</el-button>
        </div>
      </template>
      <el-row>
        <el-col :span="6">
            <div ref="cpu_load" style="height: 150px;" />
            <span>CPU Load</span>
        </el-col>
        <el-col :span="6">
            <div ref="memory_load" style="height: 150px;" />
            <span>Mem Load</span>
        </el-col>
        <el-col :span="6">
            <div ref="disk_load" style="height: 150px;" />
            <span>Disk Load</span>
        </el-col>
        <el-col :span="5">
            <div class="net_load">
                <div>
                    <el-statistic title="RX Rate KB/s" :value="realtimeInfo.net_rx_rate" />
                    <el-statistic title="TX Rate KB/s" :value="realtimeInfo.net_tx_rate" />
                </div>
            </div>
            <span>Net Load</span>
        </el-col>
        </el-row>
    </el-card>
</template>
  
<script>
import * as echarts from 'echarts';
import {SimpleWebSocket} from '@/api/realtime';
import {TEMPINFO} from '@/constant/constant';

  
export default {
    props: {
    },
    data() {
        return {
            ws: null,
            realtimeInfo: {
                cpu_usage: 0,
                memory_usage: 0,
                disk_usage: 0,
                net_rx_rate: 0,
                net_tx_rate: 0,
            },
            cpu_chart: null,
            memory_chart: null,
            disk_chart: null,
        }
    },
    mounted() {
        this.initGraph();
        this.initWebSocket();
    },
    unmounted() {
        // if the ws is null, can not use close method.
        if (this.ws) {
            this.ws.close();
        }
    },
    methods: {
        initGraph() {
            // load chart
            this.cpu_chart = echarts.init(this.$refs.cpu_load);
            this.memory_chart = echarts.init(this.$refs.memory_load);
            this.disk_chart = echarts.init(this.$refs.disk_load);
            // define chart options
            const options = {
            tooltip: {
                    trigger: 'item',
                    formatter: '{b} : {d}%'
                },
            legend: {
                    orient: 'vertical',
                    left: 'left'
                },
            series: [{
                type: 'pie',
                avoidLabelOverlap: false,
                label: {
                    show: false,
                    position: 'left'
                },
                emphasis: {
                    itemStyle: {
                        shadowBlur: 10,
                        shadowOffsetX: 0,
                        shadowColor: 'rgba(0, 0, 0, 0.5)'
                    }
                },
                labelLine: {
                    show: false
                },
                data: [
                    {value: 0, name: 'usage', itemStyle: { color: '#67C23A' }},
                    {value: 100, name: 'free', itemStyle: { color: '#e1f3d8' }},
                ],
                itemStyle: {
                }
            }]
            }

            this.cpu_chart.setOption(options);
            this.memory_chart.setOption(options);
            this.disk_chart.setOption(options);
        },
        updateEchartsOption(chart, usage, threshold=80) {
            // console.log("updateEchartsOption", chart)
            // console.log("updateEchartsOption", usage)
            const data = [
                { value: usage, name: 'usage', itemStyle: { color: usage < threshold ? '#67C23A' : '#F56C6C'}},
                { value: 100 - usage, name: 'free', itemStyle: { color: '#e1f3d8' }}
            ];
            chart.setOption({
                tooltip: {
                    trigger: 'item',
                    formatter: '{b} : {d}%'
                },
                series: [{
                    data: data,
                }]
            });
        },
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
            let realtimeResp = JSON.parse(data.data);
            // console.log("handleMessage", realtimeResp)
            realtimeResp = realtimeResp.data[0]
            this.realtimeInfo = realtimeResp; // :value to update Rate

            // console.log("realtimeInfo", realtimeResp)
            // render
            this.updateEchartsOption(this.cpu_chart, this.realtimeInfo.cpu_usage);
            this.updateEchartsOption(this.memory_chart, this.realtimeInfo.mem_usage);
            this.updateEchartsOption(this.disk_chart, this.realtimeInfo.disk_usage);
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
};
</script>
  
<style scoped>  

.net_load {
    height: 150px;
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

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
import {TEMPINFO} from '@/constant/constant';
import { ElNotification } from 'element-plus';

  
export default {
    props: {
        realtimeInfo: {
            cpu_usage: 0,
            memory_usage: 0,
            disk_usage: 0,
            net_rx_rate: 0,
            net_tx_rate: 0,
        },
    },
    data() {
        return {
            ws: null,
            cpu_chart: null,
            mem_chart: null,
            disk_chart: null,
        }
    },
    mounted() {
        this.initGraph();
    },
    unmounted() {
    },
    watch: {
        realtimeInfo: {
            immediate: true,
            handler: function (val, oldVal) {
                if (oldVal == null || val == null) {
                    return;
                }
                this.updateEchartsOption(this.cpu_chart, val.cpu_usage, TEMPINFO.threshold.cpu_usage);
                this.updateEchartsOption(this.mem_chart, val.mem_usage, TEMPINFO.threshold.mem_usage);
                this.updateEchartsOption(this.disk_chart, val.disk_usage, TEMPINFO.threshold.disk_usage);

                // alert when usage > threshold
                if (val.cpu_usage > TEMPINFO.threshold.cpu_usage) {
                    ElNotification({
                        title: 'CPU Usage Alert',
                        message: `CPU Usage: ${val.cpu_usage}%`,
                        type: 'warning'
                    });
                }
                if (val.mem_usage > TEMPINFO.threshold.mem_usage) {
                    ElNotification({
                        title: 'Memory Usage Alert',
                        message: `Memory Usage: ${val.mem_usage}%`,
                        type: 'warning'
                    });
                }
                if (val.disk_usage > TEMPINFO.threshold.disk_usage) {
                    ElNotification({
                        title: 'Disk Usage Alert',
                        message: `Disk Usage: ${val.disk_usage}%`,
                        type: 'warning'
                    });
                }
            },
            // deep: true // watch object
        }
    },
    methods: {
        initGraph() {
            // load chart
            this.cpu_chart = echarts.init(this.$refs.cpu_load);
            this.mem_chart = echarts.init(this.$refs.memory_load);
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
            this.mem_chart.setOption(options);
            this.disk_chart.setOption(options);
        },
        updateEchartsOption(chart, usage, threshold=80) {
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

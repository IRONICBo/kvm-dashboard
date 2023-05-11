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
        <el-col :span="6">
            <div ref="net_load" style="height: 150px;" />
            <span>Net Load</span>
        </el-col>
        </el-row>
    </el-card>
</template>
  
<script>
import * as echarts from 'echarts';
  
export default {
    mounted() {
        // load chart
        const cpu_load = echarts.init(this.$refs.cpu_load);
        const memory_load = echarts.init(this.$refs.memory_load);
        const disk_load = echarts.init(this.$refs.disk_load);
        const net_load = echarts.init(this.$refs.net_load);

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

        cpu_load.setOption(options);
        memory_load.setOption(options);
        disk_load.setOption(options);
        net_load.setOption(options);
            
        // 模拟 CPU 占用率变化
        setInterval(() => {
            const usage = Math.random() * 100; // 随机生成 CPU 占用率
            const data = [
                { value: usage, name: 'usage', itemStyle: { color: usage < 80 ? '#67C23A' : '#F56C6C'}},
                { value: 100 - usage, name: 'free', itemStyle: { color: '#e1f3d8' }}
            ];
            cpu_load.setOption({
                tooltip: {
                    trigger: 'item',
                    formatter: '{b} : {d}%'
                },
                series: [{
                    data: data,
                }]
            });
        }, 500); // 每隔 1 秒更新一次 CPU 占用率

        setInterval(() => {
            const usage = Math.random() * 100; // 随机生成 CPU 占用率
            const data = [
                { value: usage, name: 'usage', itemStyle: { color: usage < 80 ? '#67C23A' : '#F56C6C'}},
                { value: 100 - usage, name: 'free', itemStyle: { color: '#e1f3d8' }}
            ];
            memory_load.setOption({
                tooltip: {
                    trigger: 'item',
                    formatter: '{b} : {d}%'
                },
                series: [{
                    data: data,
                }]
            });
        }, 500); // 每隔 1 秒更新一次 CPU 占用率

        setInterval(() => {
            const usage = Math.random() * 100; // 随机生成 CPU 占用率
            const data = [
               { value: usage, name: 'usage', itemStyle: { color: usage < 80 ? '#67C23A' : '#F56C6C'}},
                { value: 100 - usage, name: 'free', itemStyle: { color: '#e1f3d8' }}
            ];
            disk_load.setOption({
                tooltip: {
                    trigger: 'item',
                    formatter: '{b} : {d}%'
                },
                series: [{
                    data: data,
                }]
            });
        }, 500); // 每隔 1 秒更新一次 CPU 占用率

        setInterval(() => {
            const usage = Math.random() * 100; // 随机生成 CPU 占用率
            const data = [
                { value: usage, name: 'usage', itemStyle: { color: usage < 80 ? '#67C23A' : '#F56C6C'}},
                { value: 100 - usage, name: 'free', itemStyle: { color: '#e1f3d8' }}
            ];
            net_load.setOption({
                tooltip: {
                    trigger: 'item',
                    formatter: '{b} : {d}%'
                },
                series: [{
                    data: data,
                }]
            });
        }, 500); // 每隔 1 秒更新一次 CPU 占用率
    }
};
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

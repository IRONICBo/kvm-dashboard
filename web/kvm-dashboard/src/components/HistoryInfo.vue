<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>VM Workload - History</span>
        <el-button class="button" text>xxx</el-button>
      </div>
    </template>
    <el-row>
      <el-col>
      <div ref="cpu_load" style="height: 300px;" />
      <span>CPU Load</span>
      </el-col>
    </el-row>
    <el-row>
      <el-col>
      <div ref="memory_load" style="height: 300px;" />
      <span>Mem Load</span>
      </el-col>
    </el-row>
    <el-row>      
      <el-col>
      <div ref="disk_load" style="height: 300px;" />
      <span>Disk Load</span>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <div ref="rx_rate" style="height: 300px;" />
        <span>RX Rate</span>
      </el-col>
      <el-col :span="12">
        <div ref="tx_rate" style="height: 300px;" />
        <span>TX Rate</span>
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
      const rx_rate = echarts.init(this.$refs.rx_rate);
      const tx_rate = echarts.init(this.$refs.tx_rate);

      // define chart options
      const options = {
        tooltip: {
          trigger: 'axis',
          formatter: function (params) {
            var value = params[0].value;
            return 'Load: ' + value + '%';
          },
          },
          xAxis: {
              type: 'category',
              boundaryGap: [0, '10%'],
              name: "Time",
              data: [],
            },
          yAxis: {
            type: 'value',
            boundaryGap: [0, '50%'],
            name: "Usage/%",
          },
          series: [{
            name: '成交',
            type: 'line',
            smooth: true,
            symbol: 'none',
            stack: 'a',
            areaStyle: {
              normal: {}
            },
            data: [],
          }]
      }

      cpu_load.setOption(options);
      memory_load.setOption(options);
      disk_load.setOption(options);
      rx_rate.setOption(options);
      tx_rate.setOption(options);
          
      // 模拟 CPU 占用率变化
      setInterval(() => {
          let option = cpu_load.getOption();
          let data = option.series[0].data;
          let xAxisData = option.xAxis[0].data;


          let now = new Date();
          let time = now.getHours() + ':' + now.getMinutes() + ':' + now.getSeconds();
          const usage = Math.random() * 100; // 随机生成 CPU 占用率

          if (xAxisData.length >= 20) {
            xAxisData.shift();
          }
          if (data.length >= 20) {
            data.shift();
          }
          data.push(usage);
          xAxisData.push(time);
          
          cpu_load.setOption({
              xAxis: {
                data: xAxisData
              },
              series: [{
                  data: data,
              }]
          });
      }, 1000); // 每隔 1 秒更新一次 CPU 占用率

      setInterval(() => {
          let option = memory_load.getOption();
          let data = option.series[0].data;
          let xAxisData = option.xAxis[0].data;


          let now = new Date();
          let time = now.getHours() + ':' + now.getMinutes() + ':' + now.getSeconds();
          const usage = Math.random() * 100; // 随机生成 CPU 占用率

          if (xAxisData.length >= 10) {
            xAxisData.shift();
          }
          if (data.length >= 10) {
            data.shift();
          }
          data.push(usage);
          xAxisData.push(time);

          memory_load.setOption({
              xAxis: {
                data: xAxisData
              },
              series: [{
                  data: data,
              }]
          });
      }, 1000); // 每隔 1 秒更新一次 CPU 占用率

      setInterval(() => {
          let option = disk_load.getOption();
          let data = option.series[0].data;
          let xAxisData = option.xAxis[0].data;


          let now = new Date();
          let time = now.getHours() + ':' + now.getMinutes() + ':' + now.getSeconds();
          const usage = Math.random() * 100; // 随机生成 CPU 占用率

          if (xAxisData.length >= 10) {
            xAxisData.shift();
          }
          if (data.length >= 10) {
            data.shift();
          }
          data.push(usage);
          xAxisData.push(time);

          disk_load.setOption({
              xAxis: {
                data: xAxisData
              },
              series: [{
                  data: data,
              }]
          });
      }, 1000); // 每隔 1 秒更新一次 CPU 占用率

      setInterval(() => {
          let option = rx_rate.getOption();
          let data = option.series[0].data;
          let xAxisData = option.xAxis[0].data;


          let now = new Date();
          let time = now.getHours() + ':' + now.getMinutes() + ':' + now.getSeconds();
          const usage = Math.random() * 100; // 随机生成 CPU 占用率

          if (xAxisData.length >= 10) {
            xAxisData.shift();
          }
          if (data.length >= 10) {
            data.shift();
          }
          data.push(usage);
          xAxisData.push(time);
          
          rx_rate.setOption({
              xAxis: {
                data: xAxisData
              },
              series: [{
                  data: data,
              }]
          });
      }, 1000); // 每隔 1 秒更新一次 CPU 占用率

      setInterval(() => {
          let option = tx_rate.getOption();
          let data = option.series[0].data;
          let xAxisData = option.xAxis[0].data;


          let now = new Date();
          let time = now.getHours() + ':' + now.getMinutes() + ':' + now.getSeconds();
          const usage = Math.random() * 100; // 随机生成 CPU 占用率

          if (xAxisData.length >= 10) {
            xAxisData.shift();
          }
          if (data.length >= 10) {
            data.shift();
          }
          data.push(usage);
          xAxisData.push(time);
          
          tx_rate.setOption({
              xAxis: {
                data: xAxisData
              },
              series: [{
                  data: data,
              }]
          });
      }, 1000); // 每隔 1 秒更新一次 CPU 占用率
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

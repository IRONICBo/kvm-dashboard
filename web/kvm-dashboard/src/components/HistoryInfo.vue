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
import { TEMPINFO } from '@/constant/constant';
import { postSimpleHistory } from '@/api/history';


export default {
  props: {
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
  },
  // props: ['realtimeInfoWithTimestamp'],
  data() {
    return {
      cpu_chart: null,
      mem_chart: null,
      disk_chart: null,
      rx_rate: null,
      tx_rate: null,

      historyInfo: {
        "cpu_usage": {
          "name": "",
          "unit": "",
          "data": [],
          "timestamp": []
        },
        "mem_usage": {
          "name": "",
          "unit": "",
          "data": [],
          "timestamp": []
        },
        "disk_usage": {
          "name": "",
          "unit": "",
          "data": [],
          "timestamp": []
        },
        "net_rx_rate": {
          "name": "",
          "unit": "",
          "data": [],
          "timestamp": []
        },
        "net_tx_rate": {
          "name": "",
          "unit": "",
          "data": [],
          "timestamp": []
        }
      },
    }
  },
  mounted() {
      this.initGraph();
  },
  watch: {
    realtimeInfoWithTimestamp: {
      immediate: true,
      handler: function(val, oldVal) {
        if (oldVal == null || val == null) {
          return;
        }
        this.updateEchartsOption(this.cpu_chart, val.data[0].cpu_usage);
        this.updateEchartsOption(this.mem_chart, val.data[0].mem_usage);
        this.updateEchartsOption(this.disk_chart, val.data[0].disk_usage);
        this.updateEchartsOption(this.rx_chart, val.data[0].net_rx_rate);
        this.updateEchartsOption(this.tx_chart, val.data[0].net_tx_rate);
      }
    },
  },
  methods: {
    initGraph() {
      // init chart
      this.cpu_chart = echarts.init(this.$refs.cpu_load);
      this.mem_chart = echarts.init(this.$refs.memory_load);
      this.disk_chart = echarts.init(this.$refs.disk_load);
      this.rx_chart = echarts.init(this.$refs.rx_rate);
      this.tx_chart = echarts.init(this.$refs.tx_rate);

      // define chart options
      const options = {
        tooltip: {
          trigger: 'axis',
          formatter: function (params) {
            var value = params[0].value;
            return 'Load: ' + value;
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
            boundaryGap: ['50%', '100%'],
            name: "Usage",
          },
          series: [{
            name: 'metric',
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

      this.cpu_chart.setOption(options);
      this.mem_chart.setOption(options);
      this.disk_chart.setOption(options);
      this.rx_chart.setOption(options);
      this.tx_chart.setOption(options);

      // get history data
      this.getHistoryInfo();
    },
    getHistoryInfo() {
      let data = {
        "UUID": TEMPINFO.uuid,
        "Fields": [
          "cpu_usage",
          "mem_usage",
          "disk_usage",
          "net_rx_rate",
          "net_tx_rate"
        ]
      }

      postSimpleHistory(data)
        .then(response => {
          let resp = response.data;
          for (let i = 0; i < resp.length; i++) {
            this.historyInfo[resp[i].name] = resp[i];
            // update chart
            this.updateEchartsSeriesOption(this.cpu_chart, this.historyInfo.cpu_usage);
            this.updateEchartsSeriesOption(this.mem_chart, this.historyInfo.mem_usage);
            this.updateEchartsSeriesOption(this.disk_chart, this.historyInfo.disk_usage);
            this.updateEchartsSeriesOption(this.rx_chart, this.historyInfo.net_rx_rate);
            this.updateEchartsSeriesOption(this.tx_chart, this.historyInfo.net_tx_rate);
          }
          console.log(this.historyInfo);
        })
        .catch(error => {
          console.log(error);
        });
    },
    updateEchartsSeriesOption(chart, seriesData) {
      let option = chart.getOption();
      let data = option.series[0].data;
      let xAxisData = option.xAxis[0].data;

      for (let i = 0; i < seriesData.timestamp.length; i++) {
        let time = this.formatTimestamp(seriesData.timestamp[i])
        if (xAxisData.length >= 20) {
              xAxisData.shift();
        }
        if (data.length >= 20) {
          data.shift();
        }
        data.push(seriesData.data[i]);
        xAxisData.push(time);
      }

      console.log("xAxisData", xAxisData, "seriesData", data);

      chart.setOption({
          xAxis: {
            data: xAxisData
          },
          series: [{
              data: data,
          }]
      });
    },
    updateEchartsOption(chart, usage) {
      let option = chart.getOption();
      let data = option.series[0].data;
      let xAxisData = option.xAxis[0].data;

      let time = this.formatTimestamp(this.realtimeInfoWithTimestamp.timestamp[0])

      if (xAxisData.length >= 20) {
            xAxisData.shift();
      }
      if (data.length >= 20) {
        data.shift();
      }
      data.push(usage);
      xAxisData.push(time);

      chart.setOption({
          xAxis: {
            data: xAxisData
          },
          series: [{
              data: data,
          }]
      });
    },
    formatTimestamp(timestamp) {
      let date = new Date(timestamp * 1000); // to milliseconds

      const options = {
        timeZone: 'Asia/Shanghai',
        hour12: false,
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit'
      };

      const formattedTime = date.toLocaleString('en-US', options);
      return formattedTime;
    },
    padZero(value) {
      return String(value).padStart(2, '0'); // padding zero before single number
    }
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

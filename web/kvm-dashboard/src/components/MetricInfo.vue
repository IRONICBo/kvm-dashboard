<template>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>Metric - History</span>
          
          <el-row :gutter="20">
            <el-col :span="6">
              <el-dropdown max-height="300px">
                <el-button type="primary">
                  Metric <el-icon class="el-icon--right"><arrow-down /></el-icon>
                </el-button>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item v-for="m in metrics" :key="m" @click="selectMetrics(m)">{{ m }}</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </el-col>
            
            <el-col :span="6">
              <el-dropdown v-model="method">
                <el-button type="primary">
                  Method <el-icon class="el-icon--right"><arrow-down /></el-icon>
                </el-button>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item v-for="m in methods" :key="m" @click="selectMethod(m)">{{ m }}</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </el-col>

            <el-col :span="12">
              <el-radio-group v-model="period" >
                <el-radio-button label="1m" />
                <el-radio-button label="5m" />
                <el-radio-button label="1h" />
              </el-radio-group>
            </el-col>
          </el-row>
        </div>
      </template>
      <el-row>
        <el-col>
        <div ref="metric" style="height: 300px;" />
        <span>{{ metric }} - {{ method }} - {{ period }}</span>
        </el-col>
      </el-row>
    </el-card>
  </template>
  
  <script>
  import * as echarts from 'echarts';
  import { ArrowDown } from '@element-plus/icons-vue'
  import { postMetricHistory, getMetricList  } from '@/api/history';
  
  
  export default {
    props: {
      uuid: "",
    },
    data() {
      return {
        metric_chart: null,
        period: "1m",
        metric: "",
        method: "mean",
        metrics: [],
        methods: ["mean", "median", "last"],
        historyInfo:{},
      }
    },
    watch: {
      metric: {
        immediate: true,
        handler: function (val, oldVal) {
          if (oldVal == null || val == null) {
              return;
          }
          
          // get history data
          this.getHistoryInfo();
        },
      },
      method: {
        immediate: true,
        handler: function (val, oldVal) {
          if (oldVal == null || val == null) {
              return;
          }
          
          // get history data
          this.getHistoryInfo();
        },
      },
      period: {
        immediate: true,
        handler: function (val, oldVal) {
          if (oldVal == null || val == null) {
              return;
          }

          // get history data
          this.getHistoryInfo();
        },
      },
      uuid: {
        immediate: true,
        handler: function (val, oldVal) {
          if (oldVal == null || val == null) {
              return;
          }

          // get history data
          this.getHistoryInfo();
        },
      }
    },
    mounted() {
        this.initOptions();
        this.initGraph();
        this.getHistoryInfo();
    },
    methods: {
      selectMetrics(m) {
        this.metric = m;
      },
      selectMethod(m) {
        this.method = m;
      },
      initOptions() {
        getMetricList()
        .then(response => {
          this.metrics = response.data;
          this.metric = this.metrics[0]; // init metric
        })
        .catch(error => {
          console.log(error);
        });
      },
      initGraph() {
        // init chart
        this.metric_chart = echarts.init(this.$refs.metric);
  
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
  
        this.metric_chart.setOption(options);
      },
      getHistoryInfo() {
        if (this.uuid == null || this.uuid == "" || this.metric == null || this.method == null || this.period == null) {
          return; // not ready
        }

        let data = {
            "UUID": this.uuid,
            "Fields": [this.metric],
            "Period": this.period,
            "Func": this.method
        }
        
        postMetricHistory(data)
          .then(response => {
            let resp = response.data;
            for (let i = 0; i < resp.length; i++) {
              this.historyInfo = resp[i];
              // update chart
              this.updateEchartsSeriesOption(this.metric_chart, this.historyInfo);
            }
          })
          .catch(error => {
            console.log(error);
          });
      },
      updateEchartsSeriesOption(chart, seriesData) {
        let option = chart.getOption();
        let data = option.series[0].data;
        let xAxisData = option.xAxis[0].data;
  
        // clear data
        data.splice(0, data.length);
        xAxisData.splice(0, xAxisData.length);

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
  .example-showcase .el-dropdown-link {
    cursor: pointer;
    color: var(--el-color-primary);
    display: flex;
    align-items: center;
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
  
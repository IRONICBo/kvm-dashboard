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
          <real-time-info></real-time-info>
          <history-info></history-info>
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

export default {
  components: { 
    Header,
    RealTimeInfo, 
    HistoryInfo, 
    ProcessInfo,
    AlertInfo,
    SystemInfo,
  }
}
</script>

<style scoped>
</style>

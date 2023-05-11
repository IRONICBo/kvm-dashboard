<template>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>System Info</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="18">
          <el-descriptions :column="6">
            <!-- Use for each -->
            <el-descriptions-item label="STATUS">STARTED</el-descriptions-item>
            <el-descriptions-item label="OS">Kirin V10</el-descriptions-item>
            <el-descriptions-item label="CPU">Phytium D2000+</el-descriptions-item>
            <el-descriptions-item label="Mem">8012 MB</el-descriptions-item>
            <el-descriptions-item label="Disk">1000 GB</el-descriptions-item>
            <el-descriptions-item label="Net">1Mb/s</el-descriptions-item>
          </el-descriptions>
        </el-col>
        <el-col :span="6">
          <el-button type="primary"><el-icon><VmStart /></el-icon></el-button>
          <el-button type="primary"><el-icon><VmStop /></el-icon></el-button>
          <el-button type="primary"><el-icon><VmPause /></el-icon></el-button>
          <el-button type="primary"><el-icon><VmRestart /></el-icon></el-button>
        </el-col>
      </el-row>
    </el-card>
</template>
  
<script>
import VmPause from '@/assets/icons/VmPause';
import VmRestart from '@/assets/icons/VmRestart';
import VmStart from '@/assets/icons/VmStart';
import VmStop from '@/assets/icons/VmStop';


export default {
    components: {
        VmStart,
        VmStop,
        VmPause,
        VmRestart
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

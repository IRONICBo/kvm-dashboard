<template>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>Process Workload</span>
        </div>
      </template>
      <el-table :data="tableData" style="width: 100%">
        <el-table-column prop="pid" label="PID" width="180" />
        <el-table-column prop="command" label="COMMAND" width="180" />
        <el-table-column prop="cpu" label="%CPU" width="180" />
        <el-table-column prop="th" label="#TH" width="180" />
        <el-table-column prop="wq" label="#WQ" width="180" />
        <el-table-column prop="ports" label="#PORTS" width="180" />
        <el-table-column prop="mem" label="#MEM" width="180" />
      </el-table>
    </el-card>
</template>
  
<script>
const tableData = [
  {
    date: '2016-05-03',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-02',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-04',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
]
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

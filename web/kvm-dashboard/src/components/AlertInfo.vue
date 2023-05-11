<template>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>Alert Info</span>
        </div>
      </template>
      <el-table :data="tableData" style="width: 100%">
        <el-table-column prop="id" label="ID" width="180" />
        <el-table-column prop="info" label="INFO" width="180" />
        <el-table-column prop="timestamp" label="TIMESTAMP" width="180" />
      </el-table>
      <div class="example-pagination-block">
        <el-pagination small layout="prev, pager, next" :total="50" />
      </div>
    </el-card>
</template>
<script>

export default ({

})
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

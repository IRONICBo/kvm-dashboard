<template>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>Alert Info</span>
        </div>
      </template>
      <el-table :data="tableData" style="width: 100%"  max-height="300px">
        <el-table-column sortable prop="metric" label="Metric" width="100" />
        <el-table-column sortable prop="value" label="Value" width="100" />
        <el-table-column sortable prop="timestamp" label="TIMESTAMP" width="180" />
      </el-table>
      <div class="example-pagination-block">
        <el-pagination v-model:current-page="currentPage"  @current-change="handleCurrentChange" small layout="prev, pager, next" :total="pageCount" />
      </div>
    </el-card>
</template>

<script>
import { TEMPINFO } from '@/constant/constant'
import { postAlertHistory } from '@/api/history'

export default ({
  data() {
    return {
      tableData: [],
      pageCount: 0,
    }
  },
  mounted() {
    this.fetchData()
  },
  methods: {
    handleCurrentChange(value) {
      let data = {
        "UUID": TEMPINFO.uuid,
        "PageSize": TEMPINFO.page.pagesize,
        "Page": value,
      }
      postAlertHistory(data)
      .then((response) => {
        let resp = response.data
        resp.data.forEach((item) => {
          item.timestamp = new Date(item.timestamp * 1000).toLocaleString()
        })

        this.tableData = resp.data
        this.pageCount = resp.total
      })
      .catch((error) => {
        console.log(error)
      })
    },
    fetchData() {
      let data = {
        "UUID": TEMPINFO.uuid,
        "PageSize": TEMPINFO.page.pagesize,
        "Page": 1,
      }
      postAlertHistory(data)
      .then((response) => {
        let resp = response.data
        resp.data.forEach((item) => {
          item.timestamp = new Date(item.timestamp * 1000).toLocaleString()
        })

        this.tableData = resp.data
        console.log(resp.total)
        this.pageCount = resp.total
      })
      .catch((error) => {
        console.log(error)
      })
    }
  },
})
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

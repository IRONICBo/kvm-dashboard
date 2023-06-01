<script>
import { defineComponent } from "vue";
import { CirclePlus, Refresh } from "@element-plus/icons-vue";
import { getRequest, postRequest } from "@/utils/axios";
import { ElMessage, ElMessageBox } from "element-plus";

export default defineComponent({
  name: "hostManageView",
  computed: {
    Refresh() {
      return Refresh;
    },
    CirclePlus() {
      return CirclePlus;
    },
  },
  data() {
    return {
      hostList: [],
      hostListCurrent: [],
      dialogFormVisible: false,
      dialogFormVisible2: false,
      currentPage: 1,
      pageSize: 20,
      total: "",
      hostInfo: {
        hostZzid: "",
        hostId: "",
        hostName: "",
        hostDescription: "",
        hostState: "",
        hostIp: "",
        hostPort: "",
        hostPortSsh: "",
        hostCpus: "",
        hostMemory: "",
        hostLoginUsername: "",
        hostLoginPassword: "",
      },
    };
  },
  methods: {
    handleEdit(number, row) {
      this.hostInfo.hostId = row.hostId;
      this.hostInfo.hostName = row.hostName;
      this.hostInfo.hostDescription = row.hostDescription;
      this.hostInfo.hostIp = row.hostIp;
      this.hostInfo.hostPort = row.hostPort;
      this.hostInfo.hostPortSsh = row.hostPortSsh;
      this.hostInfo.hostCpus = row.hostCpus;
      this.hostInfo.hostMemory = row.hostMemory;
      this.hostInfo.hostLoginUsername = row.hostLoginUsername;
      this.hostInfo.hostLoginPassword = row.hostLoginPassword;
      this.dialogFormVisible2 = true;
    },
    updateHostInfo() {
      this.$refs["ruleForm2"].validate((valid) => {
        if (valid) {
          postRequest("/api/host/updatehost", this.hostInfo).then((resp) => {
            if (resp.code == 200) {
              this.refreshHostList();
              ElMessage.success("宿主机信息修改成功！");
              this.dialogFormVisible2 = false;
            } else {
              this.refreshHostList();
              ElMessage.error("宿主机信息修改失败！");
              ElMessageBox.confirm(resp.message, "Error", {
                cancelButtonText: "Cancel",
                type: "error",
              });
              this.dialogFormVisible2 = false;
            }
            // 清空 hostInfo 信息
            for (let hostInfoKey in this.hostInfo) {
              delete this.hostInfo[hostInfoKey];
            }
          });
        } else {
          ElMessage.error("表单检验未通过，请检查所填内容！");
          return false;
        }
      });
    },
    cancelUpdateHostInfo() {
      this.dialogFormVisible2 = false;
      // 清空 hostInfo 信息
      for (let hostInfoKey in this.hostInfo) {
        delete this.hostInfo[hostInfoKey];
      }
    },
    requestDelete(number, row) {
      ElMessageBox.confirm("数据将会被删除，是否确认?", "Warning", {
        confirmButtonText: "OK",
        cancelButtonText: "Cancel",
        type: "warning",
      })
        .then(() => {
          this.handleDelete(number, row);
        })
        .catch(() => {
          ElMessage({
            type: "info",
            message: "Delete canceled",
          });
        });
    },
    handleDelete(number, row) {
      postRequest("/api/host/deletehost?hostId=" + row.hostId).then((resp) => {
        if (resp.code == 200) {
          this.refreshHostList();
          ElMessage.success("宿主机数据删除成功！");
        } else {
          ElMessage.error("数据删除失败!");
          ElMessageBox.confirm(resp.message, "Error", {
            cancelButtonText: "Cancel",
            type: "error",
          });
        }
      });
    },
    pageSizeChange(number) {
      this.pageSize = number;
      this.refreshHostList();
    },
    getCurrentPageList(currentPageTemp) {
      this.hostListCurrent = [];
      if (currentPageTemp * this.pageSize <= this.hostList.length) {
        // 当前页面能够全部填满
        for (
          let i = (currentPageTemp - 1) * this.pageSize;
          i < currentPageTemp * this.pageSize;
          i++
        ) {
          this.hostListCurrent.push(this.hostList.at(i));
        }
      } else {
        // 当前页面不能够填满
        for (
          let i = (currentPageTemp - 1) * this.pageSize;
          i < this.hostList.length;
          i++
        ) {
          this.hostListCurrent.push(this.hostList.at(i));
        }
      }
    },
    refreshHostList() {
      getRequest("/api/host/gethostlist").then((resp) => {
        if (resp.code == 200) {
          this.hostList = resp.data;
          this.total = parseInt(this.hostList.length);
          this.getCurrentPageList(1);
          ElMessage.success({
            message: "宿主机列表刷新成功！",
            duration: 1500,
          });
        } else {
          this.total = 1;
          ElMessage.error("宿主机列表刷新失败：" + resp.message);
        }
        this.currentPage = 1;
      });
    },
    addHostShow() {
      // 清空 hostInfo 信息
      for (let hostInfoKey in this.hostInfo) {
        delete this.hostInfo[hostInfoKey];
      }
      this.dialogFormVisible = true;
    },
    addHost() {
      this.$refs["ruleForm1"].validate((valid) => {
        if (valid) {
          postRequest("/api/host/addhost", this.hostInfo).then((resp) => {
            if (resp.code == 200) {
              ElMessage.success("宿主机新增成功！");
              for (let hostInfoKey in this.hostInfo) {
                delete this.hostInfo[hostInfoKey];
              }
              this.refreshHostList();
              this.dialogFormVisible = false;
            } else {
              this.refreshHostList();
              ElMessage.error("宿主机新增失败！");
              ElMessageBox.confirm(resp.message, "Error", {
                cancelButtonText: "Cancel",
                type: "error",
              });
            }
            // 清空 hostInfo 信息
            for (let hostInfoKey in this.hostInfo) {
              delete this.hostInfo[hostInfoKey];
            }
          });
        } else {
          ElMessage.error("表单检验未通过，请检查所填内容！");
          return false;
        }
      });
    },
    cancleAddHost() {
      this.dialogFormVisible = false;
      // 清空 hostInfo 信息
      for (let hostInfoKey in this.hostInfo) {
        delete this.hostInfo[hostInfoKey];
      }
    },
  },
  mounted() {
    this.refreshHostList();
  },
});
</script>

<template>
  <el-dialog
    v-model="this.dialogFormVisible"
    title="新增宿主机"
    width="500"
    top="12vh"
  >
    <el-form
      ref="ruleForm1"
      :model="this.hostInfo"
      style="margin-left: 53px"
      label-width="auto"
      label-position="left"
    >
      <el-form-item
        label="宿主机名称"
        prop="hostName"
        :rules="[{ required: true, message: '名称为必须项', trigger: 'blur' }]"
      >
        <el-input
          v-model="this.hostInfo.hostName"
          style="width: 240px"
          placeholder="请输入宿主机名称"
        />
      </el-form-item>
      <el-form-item
        label="宿主机 IP"
        prop="hostIp"
        :rules="[
          { required: true, message: '宿主机IP为必须项', trigger: 'blur' },
        ]"
      >
        <el-input
          v-model="this.hostInfo.hostIp"
          style="width: 240px"
          placeholder="请输入宿主机 IP 地址"
        />
      </el-form-item>
      <el-form-item
        label="QEMU 端口"
        prop="hostPort"
        :rules="[
          { required: true, message: 'QEMU端口为必须项', trigger: 'blur' },
        ]"
      >
        <el-input
          v-model="this.hostInfo.hostPort"
          style="width: 240px"
          placeholder="请输入宿主机 QEMU 端口号"
        />
      </el-form-item>
      <el-form-item
        label="SSH 端口"
        prop="hostPortSsh"
        :rules="[
          { required: true, message: 'SSH端口为必须项', trigger: 'blur' },
        ]"
      >
        <el-input
          v-model="this.hostInfo.hostPortSsh"
          style="width: 240px"
          placeholder="请输入宿主机 SSH 端口号"
        />
      </el-form-item>
      <el-form-item
        label="管理员用户名"
        prop="hostLoginUsername"
        :rules="[
          { required: true, message: '管理员用户名为必须项', trigger: 'blur' },
        ]"
      >
        <el-input
          v-model="this.hostInfo.hostLoginUsername"
          style="width: 240px"
          placeholder="请输入登录宿主机的管理员用户名"
        />
      </el-form-item>
      <el-form-item
        label="管理员密码"
        prop="hostLoginPassword"
        :rules="[
          { required: true, message: '管理员密码为必须项', trigger: 'blur' },
        ]"
      >
        <el-input
          v-model="this.hostInfo.hostLoginPassword"
          style="width: 240px"
          type="password"
          show-password
          placeholder="请输入登录宿主机的管理员密码"
        />
      </el-form-item>
      <el-form-item label="描述信息" prop="hostDescription">
        <el-input
          v-model="this.hostInfo.hostDescription"
          :rows="2"
          type="textarea"
          style="width: 240px"
          placeholder="请输入该宿主机的描述信息"
        />
      </el-form-item>
    </el-form>

    <el-row style="margin-top: 30px; margin-left: 280px">
      <el-button @click="this.cancleAddHost">Cancel</el-button>
      <el-button type="primary" @click="this.addHost">Confirm</el-button>
    </el-row>
  </el-dialog>

  <!--    <el-dialog v-model="this.dialogFormVisible" title="新增宿主机" width="500" top="12vh">-->
  <!--        <el-form :model="this.hostInfo">-->
  <!--            <el-form-item label="宿主机名称" style="margin-left: 50px;" prop="hostName" :rules="[{required: true, message: '名称为必须项', trigger: 'blur'}]">-->
  <!--                <el-input v-model="this.hostInfo.hostName" style="width: 220px; margin-left: 36px" placeholder="请输入宿主机名称"/>-->
  <!--            </el-form-item>-->
  <!--            <el-form-item label="宿主机 IP" style="margin-left: 50px" prop="hostIp" :rules="[{required: true, message: '宿主机IP为必须项', trigger: 'blur'}]">-->
  <!--                <el-input v-model="this.hostInfo.hostIp" style="width: 220px; margin-left: 47px" placeholder="请输入宿主机 IP 地址"/>-->
  <!--            </el-form-item>-->
  <!--            <el-form-item label="QEMU 端口" style="margin-left: 50px" prop="hostPort" :rules="[{required: true, message: 'QEMU端口为必须项', trigger: 'blur'}]">-->
  <!--                <el-input v-model="this.hostInfo.hostPort" style="width: 220px; margin-left: 32px" placeholder="请输入宿主机 QEMU 端口号"/>-->
  <!--            </el-form-item>-->
  <!--            <el-form-item label="SSH 端口" style="margin-left: 50px" prop="hostPortSsh" :rules="[{required: true, message: 'SSH端口为必须项', trigger: 'blur'}]">-->
  <!--                <el-input v-model="this.hostInfo.hostPortSsh" style="width: 220px; margin-left: 46px" placeholder="请输入宿主机 SSH 端口号"/>-->
  <!--            </el-form-item>-->
  <!--            <el-form-item label="管理员用户名" style="margin-left: 50px" prop="hostLoginUsername" :rules="[{required: true, message: '管理员用户名为必须项', trigger: 'blur'}]">-->
  <!--                <el-input v-model="this.hostInfo.hostLoginUsername" style="width: 220px; margin-left: 25px" placeholder="请输入登录宿主机的管理员用户名"/>-->
  <!--            </el-form-item>-->
  <!--            <el-form-item label="管理员密码" style="margin-left: 50px" prop="hostLoginPassword" :rules="[{required: true, message: '管理员密码为必须项', trigger: 'blur'}]">-->
  <!--                <el-input v-model="this.hostInfo.hostLoginPassword" style="width: 220px; margin-left: 39px" type="password" show-password placeholder="请输入登录宿主机的管理员密码"/>-->
  <!--            </el-form-item>-->
  <!--            <el-form-item label="描述信息" style="margin-left: 50px" prop="hostDescription">-->
  <!--                <el-input v-model="this.hostInfo.hostDescription" :rows="2" type="textarea" style="width: 220px; margin-left: 52px" placeholder="请输入该宿主机的描述信息"/>-->
  <!--            </el-form-item>-->
  <!--        </el-form>-->

  <!--        <el-row style="margin-top: 30px; margin-left: 280px">-->
  <!--            <el-button @click="this.cancleAddHost">Cancel</el-button>-->
  <!--            <el-button type="primary" @click="this.addHost">Confirm</el-button>-->
  <!--        </el-row>-->
  <!--    </el-dialog>-->

  <el-dialog
    v-model="this.dialogFormVisible2"
    title="修改宿主机信息"
    width="500"
    top="12vh"
  >
    <el-form
      ref="ruleForm2"
      :model="this.hostInfo"
      style="margin-left: 53px"
      label-width="auto"
      label-position="left"
    >
      <el-form-item
        label="宿主机名称"
        prop="hostName"
        :rules="[{ required: true, message: '名称为必须项', trigger: 'blur' }]"
      >
        <el-input v-model="this.hostInfo.hostName" style="width: 240px" />
      </el-form-item>
      <el-form-item
        label="宿主机 IP"
        prop="hostIp"
        :rules="[
          { required: true, message: '宿主机IP为必须项', trigger: 'blur' },
        ]"
      >
        <el-input v-model="this.hostInfo.hostIp" style="width: 240px" />
      </el-form-item>
      <el-form-item
        label="QEMU 端口"
        prop="hostPort"
        :rules="[
          { required: true, message: 'QEMU端口为必须项', trigger: 'blur' },
        ]"
      >
        <el-input v-model="this.hostInfo.hostPort" style="width: 240px" />
      </el-form-item>
      <el-form-item
        label="SSH 端口"
        prop="hostPortSsh"
        :rules="[
          { required: true, message: 'SSH端口为必须项', trigger: 'blur' },
        ]"
      >
        <el-input v-model="this.hostInfo.hostPortSsh" style="width: 240px" />
      </el-form-item>
      <el-form-item
        label="管理员用户名"
        prop="hostLoginUsername"
        :rules="[
          { required: true, message: '管理员用户名为必须项', trigger: 'blur' },
        ]"
      >
        <el-input
          v-model="this.hostInfo.hostLoginUsername"
          style="width: 240px"
        />
      </el-form-item>
      <el-form-item label="管理员密码" prop="hostLoginPassword">
        <el-input
          v-model="this.hostInfo.hostLoginPassword"
          style="width: 240px"
          type="password"
          show-password
          placeholder="原密码已隐藏,为空则表示不修改"
        />
      </el-form-item>
      <el-form-item label="描述信息" prop="hostDescription">
        <el-input
          v-model="this.hostInfo.hostDescription"
          :rows="2"
          type="textarea"
          style="width: 240px"
        />
      </el-form-item>
    </el-form>
    <el-row style="margin-top: 30px; margin-left: 280px">
      <el-button @click="this.cancelUpdateHostInfo">Cancel</el-button>
      <el-button type="primary" @click="this.updateHostInfo">Confirm</el-button>
    </el-row>
  </el-dialog>
  <el-card class="box-card">
    <el-row>
      <el-button
        type="primary"
        :icon="CirclePlus"
        style="margin-left: 20px; width: 125px"
        @click="this.addHostShow"
        >新增宿主机</el-button
      >
      <el-button
        plain
        :icon="Refresh"
        @click="this.refreshHostList"
        style="width: 125px"
        >刷新列表</el-button
      >
    </el-row>

    <el-row>
      <el-table
        :data="this.hostListCurrent"
        height="540"
        style="width: 95%; margin-top: 20px; margin-left: 20px"
        border
        stripe
        highlight-current-row
        scrollbar-always-on
        :header-cell-style="{ background: '#eef5fe', color: '#64a2f8' }"
      >
        <el-table-column
          prop="hostName"
          label="宿主机名称"
          width="160"
          show-overflow-tooltip
        />
        <el-table-column prop="hostIp" label="IP地址" width="160" />
        <el-table-column prop="hostPort" label="QEMU端口" width="100" />
        <el-table-column prop="hostPortSsh" label="SSH端口" width="100" />
        <el-table-column
          prop="hostLoginUser"
          label="管理员用户名"
          width="120"
          show-overflow-tooltip
        />
        <el-table-column
          prop="hostId"
          label="宿主机ID"
          width="300"
          show-overflow-tooltip
        />
        <el-table-column
          prop="hostDescription"
          label="简要描述"
          width="260"
          show-overflow-tooltip
        />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)"
              >Edit</el-button
            >
            <el-button
              size="small"
              type="danger"
              @click="requestDelete(scope.$index, scope.row)"
              plain
              >Delete</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </el-row>
    <el-row>
      <el-pagination
        v-model:current-page="this.currentPage"
        v-model:page-size="this.pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :page-sizes="[10, 20, 50]"
        :total="this.total"
        @current-change="this.getCurrentPageList(this.currentPage)"
        @size-change="this.pageSizeChange"
        style="margin-top: 25px; margin-left: 20px"
        background
      />
    </el-row>
  </el-card>
</template>

<style scoped>
.box-card {
  margin-bottom: 14px;
}
</style>
<script>
import { defineComponent } from "vue";
import { FolderOpened, Refresh } from "@element-plus/icons-vue";
import { getRequest, postRequest } from "@/utils/axios";
import { ElMessage, ElMessageBox, ElNotification } from "element-plus";
import VueJsonPretty from "vue-json-pretty";
import "vue-json-pretty/lib/styles.css";

import "element-plus/theme-chalk/el-loading.css";

export default defineComponent({
  name: "domainManageView",
  components: {
    VueJsonPretty,
  },
  computed: {
    FolderOpened() {
      return FolderOpened;
    },
    Refresh() {
      return Refresh;
    },
  },
  data() {
    return {
      // ---------------------container-------------------------
      activate: 1,
      containerHostName: "",
      collapse: [],
      VmParamInfo: {
        vmId: "",
        vmHostId: "",
        vmType: "",
        vmCreateForm: "",
        vmName: "",
        vmDescription: "",
        vmCpus: "",
        vmMemory: "",
        vmMemoryCurrent: "",
        vmOsType: "",
        vmOsTypeArch: "",
        vmOsTypeMachine: "",
        vmCpuMode: "",
        vmCpuModelFallback: "",
        vmCpuVendorid: "",
        vmCpuModelValue: "",
        vmCpuTopologySocket: "",
        vmCpuTopologyDies: "",
        vmCpuTopologyCores: "",
        vmCpuTopologyThreads: "",
        vmCpuFeaturePolicy: "",
        vmCpuFeatureName: "",
        vmDevicesDiskType: "",
        vmDevicesDiskDevice: "",
        vmDevicesDiskDriverType: "",
        vmDevicesDiskDriverIo: "",
        vmDevicesDiskDriverCache: "",
        vmDevicesDiskDriverIothread: "",
        vmDevicesDiskDriverErrorPolicy: "",
        vmDevicesDiskDriverRerrorPolicy: "",
        vmDevicesDiskDriverRetryInterval: "",
        vmDevicesDiskDriverRetryTimeout: "",
        vmDevicesDiskSourcepath: "",
        vmDevicesDiskTargetDev: "",
        vmDevicesDiskTargetBus: "",
        vmDevicesDiskBootOrder: "",
        vmDevicesDiskSize: "",
        vmDevicesDiskImageType: "",
        vmDevicesDiskImageDevice: "",
        vmDevicesDiskImageSourcepath: "",
        vmDevicesDiskImageTargetDev: "",
        vmDevicesDiskImageTargetBus: "",
        vmDevicesDiskImageBootOrder: "",
        vmDevicesDiskImageReadonly: "",
        vmThvCpu: "",
        vmThvMemory: "",
        vmThvDisk: "",
        vmInterfaceType: "",
        vmInterfaceMacAddress: "",
        vmInterfaceModelType: "",
        vmInterfaceTargetDev: "",
        vmInterfaceBandwidthInboundAverage: "",
        vmInterfaceBandwidthOutboundAverage: "",
        vmInterfaceBridgeSourceBridge: "",
        vmInterfaceBridgeBootOrder: "",
        vmInterfaceNatSourceNetwork: "",
        vmInterfaceNatBootOrder: "",
        vmGraphicsVncPort: "",
      },
      typeOptions: [
        { value: "kvm" },
        { value: "xen" },
        { value: "qemu" },
        { value: "lxc" },
      ],
      osTypeArchOptions: [{ value: "aarch64" }, { value: "x86_64" }],
      osTypeMachineOptions: [
        { value: "virt" },
        { value: "pc" },
        { value: "q35" },
        { value: "pc-q35-5.2" },
        { value: "pc-q35-4.2" },
        { value: "pc-q35-2.7" },
        { value: "pc-q35-7.1" },
        { value: "pc-i440fx-7.1" },
        { value: "pc-i440fx-2.12" },
        { value: "pc-i440fx-2.0" },
        { value: "pc-i440fx-6.2" },
      ],
      osTypeOptions: [
        { value: "hvm" },
        { value: "xen" },
        { value: "exe" },
        { value: "linux" },
        { value: "xenpvh" },
      ],
      interfaceTypeOptions: [
        { value: "bridge", disabled: false },
        { value: "network", disabled: false },
        { value: "user", disabled: true },
        { value: "ethernet", disabled: true },
        { value: "direct", disabled: true },
        { value: "vhostuser", disabled: true },
        { value: "hostdev", disabled: true },
        { value: "udp", disabled: true },
      ],
      cpuModeOptions: [
        { value: "host-passthrough" },
        { value: "host-model" },
        { value: "custom" },
      ],
      cpuFeaturePilicyOptions: [
        { value: "force" },
        { value: "require" },
        { value: "optional" },
        { value: "disable" },
        { value: "forbid" },
      ],
      diskDriverCacheOptions: [
        { value: "default" },
        { value: "none" },
        { value: "writethrough" },
        { value: "writeback" },
        { value: "directsync" },
        { value: "unsafe" },
      ],
      diskDriverErrorPolicyOptions: [
        { value: "stop" },
        { value: "report" },
        { value: "ignore" },
        { value: "enospace" },
        { value: "retry" },
      ],

      // -----------------------------------------------------------
      currentPage: 1,
      pageSize: 20,
      total: "",

      dialogVisible1: false,
      dialogVisible2: false,
      dialogVisible2Loading: false,
      domainStrObject: {},
      domainList: [],
      domainListCurrent: [],
      hostList: [],
      hostOptions: [],
      domainInfo: {},
      drawer1: false,
      drawer2: false,
      drawer3: false,
      drawer2Loading: false,
      drawer3Loading: false,
      confirmModifyLoading: false,
      notify1: "",
      ifDeleteDisk: true,
      deleteVmId: "",
      deleteVmName: "",
      deleteDiskPath: "",
      currentVmId: "",
      waitHostName: "",
      tempPath: "/",
      migrateOldVmName: "",
      migrateInfo: {
        vmId: "",
        newVmName: "",
        newHostId: "",
        newDiskPath: "",
      },
      getFolderListRequest: {
        uuid: "",
        path: "",
      },
      resourceInfoList: [],
    };
  },
  methods: {
    // -----------------------------------container------------------------------------------
    showConfirmInfo() {
      // this.VmParamInfo_CN.虚拟机ID = this.VmParamInfo.vmId;
      // this.VmParamInfo_CN.宿主机ID = this.VmParamInfo.vmHostId;
      this.activate++;
    },
    async confirmModify() {
      this.confirmModifyLoading = true;
      let resp = await this.postRequestConfirmModify();
      this.confirmModifyLoading = false;
      if (resp.code == 200) {
        this.activate++;
        this.refreshDomainList();
      } else {
        ElMessage.error("虚拟机信息修改失败：" + resp.message);
      }
    },
    postRequestConfirmModify() {
      return postRequest("/api/vm/modifydomain", this.VmParamInfo);
    },

    // --------------------------------------------------------------------------------------

    pageSizeChange(number) {
      this.pageSize = number;
      this.refreshDomainList();
    },
    getCurrentPageList(currentPageTemp) {
      this.domainListCurrent = [];
      if (currentPageTemp * this.pageSize <= this.domainList.length) {
        // 当前页面能够全部填满
        for (
          let i = (currentPageTemp - 1) * this.pageSize;
          i < currentPageTemp * this.pageSize;
          i++
        ) {
          this.domainListCurrent.push(this.domainList.at(i));
        }
      } else {
        // 当前页面不能够填满
        for (
          let i = (currentPageTemp - 1) * this.pageSize;
          i < this.domainList.length;
          i++
        ) {
          this.domainListCurrent.push(this.domainList.at(i));
        }
      }
    },

    refreshDomainList() {
      getRequest("/api/vm/domainlist").then((resp) => {
        if (resp.code == 200) {
          this.domainList = resp.data;
          for (let i = 0; i < this.domainList.length; i++) {
            if (this.domainList.at(i).vmCreateForm == "0") {
              this.domainList.at(i).vmCreateForm = "本地安装介质";
            } else if (this.domainList.at(i).vmCreateForm == "1") {
              this.domainList.at(i).vmCreateForm = "现有磁盘镜像";
            }
          }
          this.total = parseInt(this.domainList.length);
          this.getCurrentPageList(1);

          ElMessage.success({ message: "虚拟机列表刷新成功!", duration: 1500 });
        } else {
          this.total = 1;
          ElMessage.error({ message: "虚拟机列表刷新失败！", duration: 1500 });
        }
        this.currentPage = 1;
      });
    },
    showDrawer1() {
      if (this.migrateInfo.newHostId == "") {
        return;
      }
      this.getFileListByPath(this.migrateInfo.newHostId, this.tempPath);
      this.drawer1 = true;
    },
    handleDetails(number, row) {
      this.domainStrObject = row;
      this.dialogVisible1 = true;
    },
    handleMigrate(number, row) {
      this.tempPath = "/";
      this.migrateInfo.vmId = row.vmId;
      this.migrateOldVmName = row.vmName;
      this.getHostList();
      this.dialogVisible2 = true;
    },
    // async confirmMigrate() {
    //     this.dialogVisible2Loading = true;
    //     let resp = await this.postRequestConfirmMigrate();
    //     this.dialogVisible2Loading = false;
    //     if (resp.code == 200) {
    //         ElMessage.success("虚拟机迁移成功！");
    //         this.dialogVisible2 = false;
    //     } else {
    //         ElMessage.error("虚拟机迁移失败，错误原因: " + resp.message);
    //     }
    // },
    // postRequestConfirmMigrate() {
    //     return postRequest('/api/vm/migrate', this.migrateInfo);
    // },
    confirmMigrate() {
      if (
        this.migrateInfo.vmId == "" ||
        this.migrateInfo.newVmName == "" ||
        this.migrateInfo.newHostId == "" ||
        this.migrateInfo.newDiskPath == ""
      ) {
        ElMessage.error("信息不完全，请检查所填写的信息！");
        return;
      }
      this.dialogVisible2 = false;
      var newVmName = this.migrateInfo.newVmName;
      this.notify1 = ElNotification({
        title: "Prompt",
        message: "正在将虚拟机迁移到 [" + newVmName + "], 请耐心等待！",
        duration: 0,
      });
      this.postRequestConfirmMigrate();
    },
    postRequestConfirmMigrate() {
      postRequest("/api/vm/migrate", this.migrateInfo).then((resp) => {
        this.notify1.close();
        if (resp.code == 200) {
          ElMessage.success("虚拟机迁移成功！");
        } else {
          ElMessageBox.confirm(resp.message, "Error", {
            cancelButtonText: "Cancel",
            type: "error",
          });
        }
      });
    },
    confirmDrawer1() {
      this.migrateInfo.newDiskPath = this.tempPath;
      this.drawer1 = false;
    },
    getHostList() {
      getRequest("/api/host/gethostlist").then((resp) => {
        this.hostOptions = [];
        if (resp.code == 200) {
          this.hostList = resp.data;
          for (let i = 0; i < this.hostList.length; i++) {
            var temp = this.hostList.at(i);
            var tempList = new Object();
            tempList.value = temp["hostName"];
            this.hostOptions.push(tempList);
          }
        }
      });
    },
    async handleDelete() {
      this.drawer3Loading = true;
      let resp = await this.getRequestHandleDelete();
      this.drawer3Loading = false;
      this.drawer3 = false;

      if (resp.code == 200) {
        ElMessage.success("虚拟机删除成功！");
        this.refreshDomainList();
      } else {
        ElMessage.error("虚拟机删除失败：" + resp.message);
      }
    },
    getRequestHandleDelete() {
      return getRequest(
        "/api/vm/delete?uuid=" +
          this.deleteVmId +
          "&ifdelete=" +
          this.ifDeleteDisk
      );
    },
    requestDelete(number, row) {
      this.deleteDiskPath = row.vmDevicesDiskSourcepath;
      this.deleteVmName = row.vmName;
      this.deleteVmId = row.vmId;
      this.drawer3 = true;
    },
    updateFolderPath1(row) {
      if (row.isDirectory == false) {
        return;
      }
      if (row.resourceName == ".." && this.tempPath == "/") {
        return;
      }
      if (row.resourceName == "..") {
        var lastIndexOf = this.tempPath.lastIndexOf("/");
        if (lastIndexOf == 0) this.tempPath = "/";
        else this.tempPath = this.tempPath.substring(0, lastIndexOf);
        this.getFileListByPath(this.migrateInfo.newHostId, this.tempPath);
      } else {
        if (this.tempPath == "/") {
          this.tempPath = this.tempPath + row.resourceName;
        } else {
          this.tempPath = this.tempPath + "/" + row.resourceName;
        }
        this.getFileListByPath(this.migrateInfo.newHostId, this.tempPath);
      }
    },
    getFileListByPath(uuid, path) {
      this.getFolderListRequest.uuid = uuid;
      this.getFolderListRequest.path = path;
      postRequest("/api/host/getFolderList", this.getFolderListRequest).then(
        (resp) => {
          if (resp.code == 200) {
            this.resourceInfoList = resp.data;
          }
        }
      );
    },
    updateVmHostId() {
      this.tempPath = "/";
      for (var i = 0; i < this.hostList.length; i++) {
        if (this.hostList.at(i)["hostName"] == this.waitHostName) {
          this.migrateInfo.newHostId = this.hostList.at(i)["hostId"];
          break;
        }
      }
    },
    async handleModify(number, row) {
      this.drawer2 = true;
      this.activate = 1;
      this.drawer2Loading = true;
      let resp = await this.getRequestGetLibvirtXml(row.vmId);
      this.drawer2Loading = false;
      if (resp.code == 200) {
        ElMessage.success("虚拟机信息获取成功！");
        this.VmParamInfo = resp.data;
        this.VmParamInfo.vmThvCpu = parseFloat(this.VmParamInfo.vmThvCpu);
        this.VmParamInfo.vmThvMemory = parseFloat(this.VmParamInfo.vmThvMemory);
        this.VmParamInfo.vmThvDisk = parseFloat(this.VmParamInfo.vmThvDisk);
        this.VmParamInfo.vmDevicesDiskSize = parseInt(
          this.VmParamInfo.vmDevicesDiskSize
        );
        this.VmParamInfo.vmCpus = parseInt(this.VmParamInfo.vmCpus);
        this.VmParamInfo.vmMemory = parseInt(this.VmParamInfo.vmMemory);
      }
    },
    getRequestGetLibvirtXml(uuid) {
      return getRequest("/api/vm/getlibvirtxml?uuid=" + uuid);
    },
  },
  mounted() {
    // getRequest('/api/vm/domainlist').then(resp=>{
    //     if (resp.code == 200) {
    //         this.domainList = resp.data;
    //     } else {
    //         ElMessage.error("虚拟机数据获取失败！");
    //     }
    // })
    this.refreshDomainList();
  },
});
</script>

<template>
  <el-dialog
    v-model="this.drawer3"
    title="警告：是否确认删除虚拟机"
    style="width: 600px"
  >
    <el-form style="margin-left: 16px" v-loading="this.drawer3Loading">
      <el-form-item label="虚拟机名称">
        <el-input
          v-model="this.deleteVmName"
          style="width: 400px; margin-left: 30px"
          disabled
        />
      </el-form-item>
      <el-form-item label="虚拟机 ID">
        <el-input
          v-model="this.deleteVmId"
          style="width: 400px; margin-left: 40px"
          disabled
        />
      </el-form-item>
      <el-form-item label="虚拟机Disk路径">
        <el-input
          v-model="this.deleteDiskPath"
          style="width: 400px; margin-left: 3px"
          disabled
        />
      </el-form-item>
      <el-form-item label="是否删除镜像">
        <el-switch
          v-model="this.ifDeleteDisk"
          class="ml-2"
          style="margin-left: 20px"
        />
      </el-form-item>
    </el-form>

    <span style="margin-left: 65%">
      <el-button @click="this.drawer3 = false">Cancel</el-button>
      <el-button type="primary" @click="this.handleDelete">Confirm</el-button>
    </span>
  </el-dialog>

  <el-drawer
    title="磁盘存储路径"
    v-model="this.drawer1"
    :direction="'rtl'"
    size="55%"
  >
    <el-text type="primary">当前路径: {{ this.tempPath }}</el-text>

    <el-table
      :data="this.resourceInfoList"
      border
      stripe
      scrollbar-always-on
      class="drawer1-table"
      highlight-current-row
      @row-dblclick="this.updateFolderPath1"
    >
      <el-table-column property="resourceName" label="name" width="240" />
      <el-table-column property="isDirectory" label="isDir" width="160" />
      <el-table-column property="resourceSize" label="size" width="160" />
      <el-table-column property="resourceModifyTime" label="time" width="210" />
    </el-table>

    <div class="drawer1-footer" align="right">
      <el-button @click="this.drawer1 = false">cancel</el-button>
      <el-button type="primary" @click="confirmDrawer1">confirm</el-button>
    </div>
  </el-drawer>

  <!--  ========================================================配置修改========================================================================  -->
  <el-drawer
    title="虚拟机配置信息修改"
    v-model="this.drawer2"
    :direction="'rtl'"
    size="86%"
  >
    <div class="container">
      <el-steps :active="this.activate" align-center>
        <el-step title="基本信息" description="Basic information" />
        <el-step title="镜像信息" description="Image information" />
        <el-step title="高级配置" description="Advanced information" />
        <el-step title="信息确认" description="Confirm the information" />
        <el-step title="修改完成" description="Vm states information" />
      </el-steps>

      <div v-if="activate === 1" v-loading="this.drawer2Loading">
        <el-row gutter="20">
          <el-col :span="11">
            <el-form class="create-mode1">
              <el-form-item label="虚拟机名称" style="margin-left: 12px">
                <el-input
                  v-model="VmParamInfo.vmName"
                  style="width: 250px"
                  placeholder="请输入虚拟机名称"
                  disabled
                />
              </el-form-item>
              <el-form-item label="虚拟机方案" style="margin-left: 12px">
                <el-tooltip
                  effect="dark"
                  content="Default Scheme"
                  placement="right"
                >
                  <el-select v-model="VmParamInfo.vmType" style="width: 250px">
                    <el-option
                      v-for="item in typeOptions"
                      :key="item.value"
                      :value="item.value"
                    />
                  </el-select>
                </el-tooltip>
              </el-form-item>
              <el-form-item label="宿主机选择" style="margin-left: 12px">
                <el-input
                  v-model="this.containerHostName"
                  disabled
                  style="width: 250px"
                ></el-input>
              </el-form-item>
              <el-form-item label="虚拟机简介" style="margin-left: 12px">
                <el-input
                  v-model="this.VmParamInfo.vmDescription"
                  :rows="4"
                  type="textarea"
                  placeholder="请输入虚拟机简介"
                  style="width: 350px"
                />
              </el-form-item>
            </el-form>
          </el-col>

          <el-col :span="1">
            <el-divider
              direction="vertical"
              border-style="dashed"
              style="height: 100%; margin-top: 30px"
            />
          </el-col>

          <el-col :span="12">
            <el-form class="create-mode2">
              <el-form-item label="虚拟机 CPU 架构" style="margin-left: 12px">
                <el-tooltip
                  effect="dark"
                  content="Kylin OS Default"
                  placement="right"
                >
                  <el-select
                    v-model="VmParamInfo.vmOsTypeArch"
                    style="width: 250px; margin-left: 10px"
                  >
                    <el-option
                      v-for="item in osTypeArchOptions"
                      :key="item.value"
                      :value="item.value"
                    />
                  </el-select>
                </el-tooltip>
              </el-form-item>
              <el-form-item label="虚拟机 机器 类型" style="margin-left: 12px">
                <el-tooltip
                  effect="dark"
                  content="Kylin OS Default"
                  placement="right"
                >
                  <el-select
                    v-model="VmParamInfo.vmOsTypeMachine"
                    style="width: 250px; margin-left: 11px"
                  >
                    <el-option
                      v-for="item in osTypeMachineOptions"
                      :key="item.value"
                      :value="item.value"
                    />
                  </el-select>
                </el-tooltip>
              </el-form-item>
              <el-form-item label="虚拟化 OS 类型" style="margin-left: 12px">
                <el-tooltip
                  effect="dark"
                  content="Kylin OS Default"
                  placement="right"
                >
                  <el-select
                    v-model="VmParamInfo.vmOsType"
                    style="width: 250px; margin-left: 20px"
                  >
                    <el-option
                      v-for="item in osTypeOptions"
                      :key="item.value"
                      :value="item.value"
                    />
                  </el-select>
                </el-tooltip>
              </el-form-item>
              <el-form-item
                label="虚拟机 CPU 使用率告警阈值"
                style="margin-left: 12px; margin-top: 30px"
              >
                <el-slider
                  v-model="VmParamInfo.vmThvCpu"
                  step="0.01"
                  max="1.00"
                  show-input
                  input-size="small"
                  style="margin-left: 10px"
                />
              </el-form-item>
              <el-form-item
                label="虚拟机 内存 使用率告警阈值"
                style="margin-left: 12px"
              >
                <el-slider
                  v-model="VmParamInfo.vmThvMemory"
                  step="0.01"
                  max="1.00"
                  show-input
                  input-size="small"
                  style="margin-left: 10px"
                />
              </el-form-item>
              <el-form-item
                label="虚拟机 磁盘 使用率告警阈值"
                style="margin-left: 12px"
              >
                <el-slider
                  v-model="VmParamInfo.vmThvDisk"
                  step="0.01"
                  max="1.00"
                  show-input
                  input-size="small"
                  style="margin-left: 10px"
                />
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>
        <el-row>
          <el-form-item
            style="margin-left: 38%; margin-bottom: 40px; margin-top: 110px"
          >
            <el-button
              type="primary"
              @click="activate++"
              style="margin-right: 20px; width: 100px"
              >下一步</el-button
            >
            <el-button style="width: 100px" @click="this.drawer2 = false"
              >取 消</el-button
            >
          </el-form-item>
        </el-row>
      </div>

      <div v-if="activate === 2">
        <el-row gutter="20">
          <el-col :span="11">
            <el-form class="create-mode3">
              <el-form-item label="虚拟机安装途径">
                <el-radio-group
                  v-model="VmParamInfo.vmCreateForm"
                  style="margin-left: 10px"
                  disabled
                >
                  <el-radio label="0" border style="margin-bottom: 20px"
                    >本地安装介质 (IOS Image or CDROM)</el-radio
                  >>
                  <el-radio
                    label="1"
                    border
                    style="margin-bottom: 20px; width: 290px"
                    >现有磁盘镜像 (Disk Image)</el-radio
                  >>
                  <el-radio label="2" border disabled style="width: 290px"
                    >网络安装 (HTTP, HTTPS or FTP)</el-radio
                  >>
                </el-radio-group>
              </el-form-item>
              <div v-if="VmParamInfo.vmCreateForm == '0'">
                <el-form-item label="ISO/CDROM 镜像路径">
                  <el-input
                    v-model="VmParamInfo.vmDevicesDiskImageSourcepath"
                    style="width: 250px"
                    disabled
                  />
                </el-form-item>
                <el-form-item label="Disk 磁盘 安装路径">
                  <el-input
                    v-model="VmParamInfo.vmDevicesDiskSourcepath"
                    style="width: 250px; margin-left: 21px"
                    disabled
                  />
                </el-form-item>
                <el-form-item label="Disk 磁盘类型选择">
                  <el-radio-group
                    v-model="this.VmParamInfo.vmDevicesDiskDriverType"
                    disabled
                  >
                    <el-radio label="raw" style="margin-left: 30px" />
                    <el-radio label="qcow2" />
                  </el-radio-group>
                </el-form-item>
                <el-form-item label="虚拟机磁盘容量 (G)">
                  <el-input-number
                    v-model="VmParamInfo.vmDevicesDiskSize"
                    max="800"
                    min="8"
                    style="margin-left: 22px"
                  />
                </el-form-item>
              </div>

              <div v-if="VmParamInfo.vmCreateForm == '1'">
                <el-form-item
                  label="现有磁盘镜像路径"
                  :rules="[{ required: true, message: '提醒: 请选择镜像路径' }]"
                >
                  <el-input
                    v-model="VmParamInfo.vmDevicesDiskSourcepath"
                    style="width: 250px; margin-left: 31px"
                    placeholder="请选择镜像路径"
                  />
                </el-form-item>
                <el-form-item label="虚拟机磁盘容量 (G)">
                  <el-input-number
                    v-model="VmParamInfo.vmDevicesDiskSize"
                    max="800"
                    min="8"
                    style="margin-left: 33px"
                  />
                </el-form-item>
              </div>
            </el-form>
          </el-col>

          <el-col :span="1">
            <el-divider
              direction="vertical"
              border-style="dashed"
              style="height: 100%; margin-top: 30px"
            />
          </el-col>

          <el-col :span="11">
            <el-form class="create-mode4">
              <el-form-item label="虚拟机 CPU 数量">
                <el-input-number
                  v-model="this.VmParamInfo.vmCpus"
                  :min="1"
                  :max="128"
                  style="margin-left: 75px"
                />
              </el-form-item>
              <el-form-item label="虚拟机内存分配(MB)">
                <el-input-number
                  v-model="this.VmParamInfo.vmMemory"
                  :min="128"
                  style="margin-left: 55px"
                />
              </el-form-item>
              <el-form-item label="虚拟机 输入带宽限制(KB/s)">
                <el-input
                  v-model="this.VmParamInfo.vmInterfaceBandwidthInboundAverage"
                  style="margin-left: 12px; width: 220px"
                  placeholder="默认为空，不限制输入带宽"
                />
              </el-form-item>
              <el-form-item label="虚拟机 输出带宽限制(KB/s)">
                <el-input
                  v-model="this.VmParamInfo.vmInterfaceBandwidthOutboundAverage"
                  style="margin-left: 12px; width: 220px"
                  placeholder="默认为空，不限制输出带宽"
                />
              </el-form-item>
              <el-form-item label="虚拟机 虚拟网卡类型">
                <el-select
                  v-model="VmParamInfo.vmInterfaceType"
                  style="width: 220px; margin-left: 53px"
                >
                  <el-option
                    v-for="item in interfaceTypeOptions"
                    :key="item.value"
                    :value="item.value"
                    :disabled="item.disabled"
                  />
                </el-select>
              </el-form-item>
              <el-form-item label="虚拟机网卡 MAC 地址">
                <el-input
                  v-model="this.VmParamInfo.vmInterfaceMacAddress"
                  placeholder="如果为空，则已经自动生成!"
                  style="width: 220px; margin-left: 44px"
                ></el-input>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>
        <el-row>
          <el-form-item style="margin-left: 38%; margin-top: 40px">
            <el-button
              type="primary"
              @click="activate--"
              style="margin-right: 20px; width: 100px"
              >上一步</el-button
            >
            <el-button type="primary" @click="activate++" style="width: 100px"
              >下一步</el-button
            >
          </el-form-item>
        </el-row>
      </div>

      <div v-if="activate === 3">
        <el-row>
          <el-collapse
            v-model="this.collapse"
            style="margin-left: 80px; margin-top: 30px; width: 1080px"
          >
            <el-collapse-item title="虚拟机 CPU 高级配置" name="1">
              <el-form style="margin-left: 22%">
                <el-form-item label="CPU Mode">
                  <el-select
                    v-model="this.VmParamInfo.vmCpuMode"
                    style="width: 300px; margin-left: 120px"
                    placeholder="请选择虚拟机 CPU 模式"
                  >
                    <el-option
                      v-for="item in cpuModeOptions"
                      :key="item.value"
                      :value="item.value"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item label="CPU Model-Fallback">
                  <el-radio-group v-model="this.VmParamInfo.vmCpuModelFallback">
                    <el-radio label="allow" style="margin-left: 58px" />
                    <el-radio label="forbid" />
                  </el-radio-group>
                </el-form-item>
                <el-form-item label="CPU Model-Value">
                  <el-input
                    v-model="this.VmParamInfo.vmCpuModelValue"
                    placeholder="请设置 CPU 模型值"
                    style="margin-left: 73px; width: 300px"
                  ></el-input>
                </el-form-item>
                <el-form-item label="CPU Vendor-ID">
                  <el-input
                    v-model="this.VmParamInfo.vmCpuVendorid"
                    placeholder="请设置 CPU 供应商 ID"
                    style="margin-left: 90px; width: 300px"
                  ></el-input>
                </el-form-item>
                <el-form-item label="CPU Topology-Sockets">
                  <el-input
                    v-model="this.VmParamInfo.vmCpuTopologySocket"
                    placeholder="请设置 CPU 插槽总数"
                    style="margin-left: 37px; width: 300px"
                  />
                </el-form-item>
                <el-form-item label="CPU Topology-Dies">
                  <el-input
                    v-model="this.VmParamInfo.vmCpuTopologyDies"
                    placeholder="请设置 CPU 每个插槽的裸片数"
                    style="margin-left: 61px; width: 300px"
                  />
                </el-form-item>
                <el-form-item label="CPU Topology-Cores">
                  <el-input
                    v-model="this.VmParamInfo.vmCpuTopologyCores"
                    placeholder="请设置 CPU 每个裸片的内核数"
                    style="margin-left: 52px; width: 300px"
                  />
                </el-form-item>
                <el-form-item label="CPU Topology-Threads">
                  <el-input
                    v-model="this.VmParamInfo.vmCpuTopologyThreads"
                    placeholder="请设置 CPU 每个内核的线程数"
                    style="margin-left: 38px; width: 300px"
                  />
                </el-form-item>
                <el-form-item label="CPU Feature-Policy">
                  <el-select
                    v-model="VmParamInfo.vmCpuFeaturePolicy"
                    style="width: 300px; margin-left: 63px"
                    placeholder="请选择虚拟机 CPU 特性策略"
                  >
                    <el-option
                      v-for="item in cpuFeaturePilicyOptions"
                      :key="item.value"
                      :value="item.value"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item label="CPU Feature-Name">
                  <el-input
                    v-model="this.VmParamInfo.vmCpuFeatureName"
                    placeholder="请为 CPU 特性指定名称"
                    style="margin-left: 64px; width: 300px"
                  />
                </el-form-item>
              </el-form>
            </el-collapse-item>

            <el-collapse-item title="虚拟机 Disk 磁盘高级配置" name="2">
              <el-form style="margin-left: 22%">
                <el-form-item label="Disk Type">
                  <el-input
                    v-model="this.VmParamInfo.vmDevicesDiskType"
                    style="width: 300px; margin-left: 128px"
                    disabled
                  ></el-input>
                </el-form-item>
                <el-form-item label="Disk Device">
                  <el-radio-group
                    v-model="this.VmParamInfo.vmDevicesDiskDevice"
                    disabled
                  >
                    <el-radio label="disk" style="margin-left: 120px" />
                    <el-radio label="cdrom" />
                    <el-radio label="floppy" />
                  </el-radio-group>
                </el-form-item>
                <el-form-item label="Disk Driver-IO">
                  <el-radio-group
                    v-model="this.VmParamInfo.vmDevicesDiskDriverIo"
                  >
                    <el-radio label="native" style="margin-left: 103px" />
                    <el-radio label="threads" />
                  </el-radio-group>
                </el-form-item>
                <el-form-item label="Disk Driver-Cache">
                  <el-select
                    v-model="this.VmParamInfo.vmDevicesDiskDriverCache"
                    style="width: 300px; margin-left: 75px"
                    placeholder="请设置 Disk 磁盘缓存模式"
                  >
                    <el-option
                      v-for="item in diskDriverCacheOptions"
                      :key="item.value"
                      :value="item.value"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item label="Disk Driver-IOThread">
                  <el-input
                    v-model="this.VmParamInfo.vmDevicesDiskDriverIothread"
                    placeholder="请设置 Disk 磁盘 IO 线程"
                    style="margin-left: 57px; width: 300px"
                  />
                </el-form-item>
                <el-form-item label="Disk Driver-Error-Policy">
                  <el-select
                    v-model="this.VmParamInfo.vmDevicesDiskDriverErrorPolicy"
                    style="width: 300px; margin-left: 37px"
                    placeholder="请设置 Disk IO写错误处理策略"
                  >
                    <el-option
                      v-for="item in diskDriverErrorPolicyOptions"
                      :key="item.value"
                      :value="item.value"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item label="Disk Driver-Rerror-Policy">
                  <el-select
                    v-model="this.VmParamInfo.vmDevicesDiskDriverRerrorPolicy"
                    style="width: 300px; margin-left: 30px"
                    placeholder="请设置 Disk IO读错误处理策略"
                  >
                    <el-option
                      v-for="item in diskDriverErrorPolicyOptions"
                      :key="item.value"
                      :value="item.value"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item label="Disk Driver-Retry-Interval">
                  <el-input
                    v-model="this.VmParamInfo.vmDevicesDiskDriverRetryInterval"
                    placeholder="请设置 Disk IO错误重试间隔时间(ms)"
                    style="margin-left: 29px; width: 300px"
                  />
                </el-form-item>
                <el-form-item label="Disk Driver-Retry-Timeout">
                  <el-input
                    v-model="this.VmParamInfo.vmDevicesDiskDriverRetryTimeout"
                    placeholder="请设置 Disk IO错误重试超时时间(ms)"
                    style="margin-left: 24px; width: 300px"
                  />
                </el-form-item>
                <el-form-item label="Disk Target-Dev">
                  <el-input
                    v-model="this.VmParamInfo.vmDevicesDiskTargetDev"
                    placeholder="请指定磁盘的逻辑设备名称"
                    style="margin-left: 90px; width: 300px"
                    disabled
                  />
                </el-form-item>
                <el-form-item label="Disk Target-Bus">
                  <el-input
                    v-model="this.VmParamInfo.vmDevicesDiskTargetBus"
                    style="width: 300px; margin-left: 90px"
                    disabled
                  ></el-input>
                </el-form-item>
              </el-form>
            </el-collapse-item>
            <el-collapse-item title="虚拟机 ISO/CDROM 镜像高级配置" name="3">
              <el-form style="margin-left: 22%">
                <el-form-item label="Image Disk Type">
                  <el-input
                    v-model="this.VmParamInfo.vmDevicesDiskImageType"
                    style="width: 300px; margin-left: 90px"
                    disabled
                  ></el-input>
                </el-form-item>
                <el-form-item label="Image Disk Device">
                  <el-radio-group
                    v-model="this.VmParamInfo.vmDevicesDiskImageDevice"
                    disabled
                  >
                    <el-radio label="disk" style="margin-left: 90px" />
                    <el-radio label="cdrom" />
                    <el-radio label="floppy" />
                  </el-radio-group>
                </el-form-item>
                <el-form-item label="Image Target-Dev">
                  <el-input
                    v-model="this.VmParamInfo.vmDevicesDiskImageTargetDev"
                    placeholder="请指定 Image 磁盘的逻辑设备名称"
                    style="margin-left: 80px; width: 300px"
                    disabled
                  />
                </el-form-item>
                <el-form-item label="Image Target-Bus">
                  <el-input
                    v-model="this.VmParamInfo.vmDevicesDiskImageTargetBus"
                    style="width: 300px; margin-left: 80px"
                    disabled
                  ></el-input>
                </el-form-item>
              </el-form>
            </el-collapse-item>

            <el-collapse-item title="虚拟机 网卡 高级配置" name="4">
              <el-form style="margin-left: 22%">
                <el-form-item label="Interface Model-Type">
                  <el-input
                    v-model="this.VmParamInfo.vmInterfaceModelType"
                    placeholder="请指定 Interface 驱动类型"
                    style="margin-left: 60px; width: 300px"
                  />
                </el-form-item>
                <el-form-item label="Interface Target-Dec">
                  <el-input
                    v-model="this.VmParamInfo.vmInterfaceTargetDev"
                    placeholder="请指定 Interface 虚拟网卡名称"
                    style="margin-left: 64px; width: 300px"
                  />
                </el-form-item>
                <el-form-item label="Interface Source-Bridge">
                  <el-input
                    v-model="this.VmParamInfo.vmInterfaceBridgeSourceBridge"
                    placeholder="(Bridge类型填写) 请指定 Interface 网桥名称"
                    style="margin-left: 45px; width: 300px"
                  />
                </el-form-item>
                <el-form-item label="Interface Source-Network">
                  <el-input
                    v-model="this.VmParamInfo.vmInterfaceNatSourceNetwork"
                    placeholder="(NAT类型填写) 请指定 Interface 网桥名称"
                    style="margin-left: 33px; width: 300px"
                  />
                </el-form-item>
              </el-form>
            </el-collapse-item>
            <el-collapse-item title="虚拟机 Graphics VNC 高级配置" name="5">
              <el-form style="margin-left: 22%">
                <el-form-item label="Customize the VNC Port">
                  <el-input
                    v-model="this.VmParamInfo.vmGraphicsVncPort"
                    placeholder="如果为空，则已经自动生成！"
                    style="margin-left: 140px; width: 300px"
                  />
                </el-form-item>
              </el-form>
            </el-collapse-item>
          </el-collapse>
        </el-row>
        <el-row>
          <el-form-item
            style="margin-left: 38%; margin-bottom: 40px; margin-top: 30px"
          >
            <el-button
              type="primary"
              @click="activate--"
              style="margin-right: 20px; width: 100px"
              >上一步</el-button
            >
            <el-button
              type="primary"
              @click="this.showConfirmInfo"
              style="width: 100px"
              >下一步</el-button
            >
          </el-form-item>
        </el-row>
      </div>

      <div v-if="activate === 4" v-loading="this.confirmModifyLoading">
        <el-row>
          <el-form-item style="margin-left: 32%; margin-top: 20px">
            <vue-json-pretty
              :data="this.VmParamInfo"
              style="width: 700px"
            ></vue-json-pretty>
          </el-form-item>
        </el-row>
        <el-row>
          <el-form-item style="margin-left: 38%; margin-top: 20px">
            <el-button
              type="primary"
              @click="activate--"
              style="margin-right: 20px; width: 100px"
              >上一步</el-button
            >
            <el-button
              type="primary"
              @click="this.confirmModify"
              style="width: 100px"
              >确认修改</el-button
            >
          </el-form-item>
        </el-row>
      </div>

      <div v-if="activate === 5">
        <el-row>
          <el-icon
            size="200px"
            color="#32cd32"
            style="margin-left: 39%; margin-top: 120px"
            @click="this.drawer2 = false"
            ><CircleCheck
          /></el-icon>
        </el-row>
      </div>
    </div>
  </el-drawer>
  <!--  ========================================================================================================================  -->

  <el-dialog
    v-model="this.dialogVisible1"
    title="详细信息"
    width="50%"
    top="5vh"
  >
    <el-form>
      <el-form-item style="margin-left: 60px">
        <vue-json-pretty :data="this.domainStrObject"></vue-json-pretty>
      </el-form-item>
    </el-form>
    <el-button @click="this.dialogVisible1 = false" style="margin-left: 600px"
      >Cancel</el-button
    >
  </el-dialog>

  <el-dialog
    v-model="this.dialogVisible2"
    title="虚拟机静态迁移"
    width="40%"
    show-close
  >
    <el-form
      style="margin-left: 64px"
      :model="this.migrateInfo"
      label-width="auto"
      label-position="left"
    >
      <el-form-item label="原虚拟机名称">
        <el-input
          v-model="this.migrateOldVmName"
          style="width: 250px; margin-left: 20px"
          disabled
        />
      </el-form-item>
      <el-form-item label="宿主机选择">
        <el-select
          v-model="this.waitHostName"
          style="width: 250px; margin-left: 19px"
          @change="this.updateVmHostId"
          placeholder="请选择目的Host，原Host不可用"
        >
          <el-option
            v-for="item in hostOptions"
            :key="item.value"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="新虚拟机名称"
        prop="newVmName"
        :rules="[{ required: true, trigger: 'blur' }]"
        :show-message="false"
      >
        <el-input
          v-model="this.migrateInfo.newVmName"
          style="width: 250px; margin-left: 20px"
          placeholder="请输入新虚拟机名称"
        />
      </el-form-item>
      <el-form-item
        label="磁盘存储路径"
        prop="newDiskPath"
        :rules="[{ required: true, trigger: 'blur' }]"
        :show-message="false"
      >
        <el-input
          v-model="this.migrateInfo.newDiskPath"
          style="width: 250px; margin-left: 20px"
          placeholder="请输入磁盘存储路径"
        />
        <el-button
          type="primary"
          :icon="FolderOpened"
          round
          plain
          style="margin-left: 25px"
          @click="this.showDrawer1"
        ></el-button>
      </el-form-item>
    </el-form>
    <el-row style="margin-top: 30px; margin-left: 200px">
      <el-button @click="this.dialogVisible2 = false">Cancel</el-button>
      <el-button type="primary" @click="this.confirmMigrate">Confirm</el-button>
    </el-row>
  </el-dialog>

  <el-card class="box-card">
    <el-row>
      <el-button
        plain
        :icon="Refresh"
        style="margin-left: 20px; width: 125px"
        @click="this.refreshDomainList"
        >刷新列表</el-button
      >
    </el-row>
    <el-row>
      <el-table
        :data="this.domainList"
        style="width: 97%; height: 540px; margin-top: 20px; margin-left: 20px"
        border
        stripe
        highlight-current-row
        scrollbar-always-on
        :header-cell-style="{ background: '#eef5fe', color: '#64a2f8' }"
      >
        <el-table-column
          prop="vmName"
          label="虚拟机名称"
          width="140"
          show-overflow-tooltip
        />
        <el-table-column prop="vmType" label="虚拟机类型" width="100" />
        <el-table-column
          prop="vmCreateForm"
          label="创建形式"
          width="120"
          show-overflow-tooltip
        />
        <el-table-column prop="vmCpus" label="CPU数量" width="90" />
        <el-table-column prop="vmMemory" label="内存空间(MB)" width="120" />
        <el-table-column prop="vmOsType" label="OS引导类型" width="110" />
        <el-table-column
          prop="vmOsTypeArch"
          label="CPU架构"
          width="90"
          show-overflow-tooltip
        />
        <el-table-column
          prop="vmOsTypeMachine"
          label="机器类型"
          width="90"
          show-overflow-tooltip
        />
        <el-table-column
          prop="vmDevicesDiskSize"
          label="磁盘容量(GB)"
          width="120"
        />
        <el-table-column
          prop="vmId"
          label="虚拟机ID"
          width="330"
          show-overflow-tooltip
        />
        <el-table-column
          prop="vmHostId"
          label="宿主机ID"
          width="330"
          show-overflow-tooltip
        />
        <el-table-column prop="vmThvCpu" label="CPU告警阈值" width="120" />
        <el-table-column prop="vmThvMemory" label="内存告警阈值" width="120" />
        <el-table-column prop="vmThvDisk" label="磁盘告警阈值" width="120" />
        <el-table-column label="操作" width="305" fixed="right">
          <template #default="scope">
            <el-button
              size="small"
              @click="this.handleDetails(scope.$index, scope.row)"
              >Details</el-button
            >
            <el-button
              size="small"
              type="primary"
              @click="this.handleModify(scope.$index, scope.row)"
              plain
              >Modify</el-button
            >
            <el-button
              size="small"
              type="primary"
              @click="this.handleMigrate(scope.$index, scope.row)"
              plain
              >Migrate</el-button
            >
            <el-button
              size="small"
              type="danger"
              @click="this.requestDelete(scope.$index, scope.row)"
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
.create-mode1 {
  width: 500px;
  margin-top: 80px;
  margin-left: 90px;
  .el-form-item {
    margin-bottom: 30px;
  }
}
.create-mode2 {
  width: 550px;
  margin-top: 70px;
}
.create-mode3 {
  width: 500px;
  margin-top: 70px;
  margin-left: 100px;
  .el-form-item {
    margin-bottom: 25px;
  }
}
.create-mode4 {
  width: 550px;
  margin-top: 80px;
  margin-left: 50px;
  .el-form-item {
    margin-bottom: 28px;
  }
}

.drawer1-table {
  width: 770px;
  height: 550px;
  margin-top: 13px;
  margin-left: 10px;
}
.drawer1-footer {
  margin-top: 30px;
  margin-right: 15px;
}
.drawer2-footer {
  margin-top: 30px;
  margin-right: 15px;
}
.box-card {
  margin-bottom: 14px;
}
</style>
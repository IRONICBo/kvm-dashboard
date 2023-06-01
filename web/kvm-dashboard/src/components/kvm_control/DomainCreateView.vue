<script>
import { defineComponent } from "vue";
import { getRequest, postRequest } from "@/utils/axios";
import { Document, FolderOpened } from "@element-plus/icons-vue";
import { ElMessage, ElMessageBox } from "element-plus";
import VueJsonPretty from "vue-json-pretty";

export default defineComponent({
  name: "domainCreateView",
  components: { VueJsonPretty },
  computed: {
    Document() {
      return Document;
    },
    FolderOpened() {
      return FolderOpened;
    },
  },
  data() {
    return {
      // 步骤条数据
      activate: 1,
      drawer1: false,
      diskpath1: "/",
      drawer2: false,
      diskpath2: "/",
      drawer3: false,
      createDomainLoading: false,
      collapse: [],
      vnc: "no",
      VmParamInfo: {
        vmId: "",
        vmHostId: "localhost",
        vmType: "kvm",
        vmCreateForm: "0",
        vmName: "",
        vmDescription: "",
        vmCpus: 1,
        vmMemory: 2048,
        vmMemoryCurrent: "",
        vmOsType: "hvm",
        vmOsTypeArch: "aarch64",
        vmOsTypeMachine: "virt",
        vmCpuMode: "host-passthrough",
        vmCpuModelFallback: "",
        vmCpuVendorid: "",
        vmCpuModelValue: "",
        vmCpuTopologySocket: "",
        vmCpuTopologyDies: "",
        vmCpuTopologyCores: "",
        vmCpuTopologyThreads: "",
        vmCpuFeaturePolicy: "",
        vmCpuFeatureName: "",
        vmDevicesDiskType: "file",
        vmDevicesDiskDevice: "disk",
        vmDevicesDiskDriverType: "qcow2",
        vmDevicesDiskDriverIo: "",
        vmDevicesDiskDriverCache: "",
        vmDevicesDiskDriverIothread: "",
        vmDevicesDiskDriverErrorPolicy: "",
        vmDevicesDiskDriverRerrorPolicy: "",
        vmDevicesDiskDriverRetryInterval: "",
        vmDevicesDiskDriverRetryTimeout: "",
        vmDevicesDiskSourcepath: "",
        vmDevicesDiskTargetDev: "vda",
        vmDevicesDiskTargetBus: "virtio",
        vmDevicesDiskBootOrder: "",
        vmDevicesDiskSize: 20,
        vmDevicesDiskImageType: "file",
        vmDevicesDiskImageDevice: "cdrom",
        vmDevicesDiskImageSourcepath: "",
        vmDevicesDiskImageTargetDev: "sda",
        vmDevicesDiskImageTargetBus: "scsi",
        vmDevicesDiskImageBootOrder: "",
        vmDevicesDiskImageReadonly: "",
        vmThvCpu: 0.9,
        vmThvMemory: 0.9,
        vmThvDisk: 0.9,
        vmInterfaceType: "network",
        vmInterfaceMacAddress: "",
        vmInterfaceModelType: "virtio",
        vmInterfaceTargetDev: "",
        vmInterfaceBandwidthInboundAverage: "",
        vmInterfaceBandwidthOutboundAverage: "",
        vmInterfaceBridgeSourceBridge: "br0",
        vmInterfaceBridgeBootOrder: "",
        vmInterfaceNatSourceNetwork: "default",
        vmInterfaceNatBootOrder: "",
        vmGraphicsVncPort: "",
      },
      VmParamInfo_CN: {
        虚拟机ID: "",
        宿主机ID: "",
        虚拟机创建类型: "",
        虚拟机创建形式: "",
        虚拟机名称: "",
        虚拟机描述: "",
        虚拟机CPU数量: "",
        虚拟机内存容量: "",
        虚拟机引导操作系统类型: "",
        虚拟机CPU架构: "",
        虚拟机机器类型: "",
        虚拟机CPU属性: "",
        虚拟机CPU供应商ID: "",
        虚拟机CPU模型允许回退: "",
        虚拟机CPU模型模型值: "",
        虚拟机CPU插槽总数: "",
        虚拟机CPU每个插槽裸片数: "",
        虚拟机CPU每个裸片内核数: "",
        虚拟机CPU每个内核线程数: "",
        虚拟机CPU特性策略: "",
        虚拟机CPU特性名称: "",
      },
      HostParamInfo: {
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
      },
      resourceInfoListImage: [],
      resourceInfoListDisk: [],
      typeOptions: [
        { value: "kvm" },
        { value: "xen" },
        { value: "qemu" },
        { value: "lxc" },
      ],
      hostList: [],
      hostOptionDefault: "localhost",
      hostOptions: [],
      getFolderListRequest: {
        uuid: "",
        path: "",
      },
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
      diskTypeOptions: [
        { value: "file", disabled: false },
        { value: "block", disabled: false },
        { value: "dir", disabled: false },
        { value: "network", disabled: true },
        { value: "volume", disabled: true },
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
      diskTargetBusOptions: [
        { value: "virtio" },
        { value: "scsi" },
        { value: "usb" },
        { value: "sata" },
      ],
    };
  },
  methods: {
    showDrawer1() {
      this.getFileListByPath(this.VmParamInfo.vmHostId, this.diskpath1, 1);
      this.drawer1 = true;
    },
    showDrawer2() {
      this.getFileListByPath(this.VmParamInfo.vmHostId, this.diskpath2, 2);
      this.drawer2 = true;
    },
    showDrawer3() {
      this.getFileListByPath(this.VmParamInfo.vmHostId, this.diskpath1, 1);
      this.drawer3 = true;
    },
    updateFolderPath1(row) {
      if (row.isDirectory == false) {
        if (this.diskpath1.endsWith("/")) {
          this.VmParamInfo.vmDevicesDiskImageSourcepath =
            this.diskpath1 + row.resourceName;
        } else {
          this.VmParamInfo.vmDevicesDiskImageSourcepath =
            this.diskpath1 + "/" + row.resourceName;
        }
        ElMessage.success(
          "已选中镜像文件: " + this.VmParamInfo.vmDevicesDiskImageSourcepath
        );
        return;
      }
      if (row.resourceName == ".." && this.diskpath1 == "/") {
        return;
      }
      if (row.resourceName == "..") {
        var lastIndexOf = this.diskpath1.lastIndexOf("/");
        if (lastIndexOf == 0) this.diskpath1 = "/";
        else this.diskpath1 = this.diskpath1.substring(0, lastIndexOf);
        this.getFileListByPath(this.VmParamInfo.vmHostId, this.diskpath1, 1);
      } else {
        if (this.diskpath1 == "/") {
          this.diskpath1 = this.diskpath1 + row.resourceName;
        } else {
          this.diskpath1 = this.diskpath1 + "/" + row.resourceName;
        }
        this.getFileListByPath(this.VmParamInfo.vmHostId, this.diskpath1, 1);
      }
    },
    updateFolderPath2(row) {
      if (row.isDirectory == false) {
        return;
      }
      if (row.resourceName == ".." && this.diskpath2 == "/") {
        return;
      }
      if (row.resourceName == "..") {
        var lastIndexOf = this.diskpath2.lastIndexOf("/");
        if (lastIndexOf == 0) this.diskpath2 = "/";
        else this.diskpath2 = this.diskpath2.substring(0, lastIndexOf);
        this.getFileListByPath(this.VmParamInfo.vmHostId, this.diskpath2, 2);
      } else {
        if (this.diskpath2 == "/") {
          this.diskpath2 = this.diskpath2 + row.resourceName;
        } else {
          this.diskpath2 = this.diskpath2 + "/" + row.resourceName;
        }
        this.getFileListByPath(this.VmParamInfo.vmHostId, this.diskpath2, 2);
      }
    },
    getFileListByPath(uuid, path, index) {
      this.getFolderListRequest.uuid = uuid;
      this.getFolderListRequest.path = path;
      postRequest("/api/host/getFolderList", this.getFolderListRequest).then(
        (resp) => {
          if (resp.code == 200) {
            if (index == 1) {
              this.resourceInfoListImage = resp.data;
            } else if (index == 2) {
              this.resourceInfoListDisk = resp.data;
            }
          }
        }
      );
    },
    updateVmHostId() {
      for (var i = 0; i < this.hostList.length; i++) {
        if (this.hostOptionDefault == "localhost") {
          this.VmParamInfo.vmHostId = "localhost";
          break;
        }
        if (this.hostList.at(i)["hostName"] == this.hostOptionDefault) {
          this.VmParamInfo.vmHostId = this.hostList.at(i)["hostId"];
          break;
        }
      }
      console.log(this.VmParamInfo.vmHostId);
    },
    showConfirmInfo() {
      // this.VmParamInfo_CN.虚拟机ID = this.VmParamInfo.vmId;
      // this.VmParamInfo_CN.宿主机ID = this.VmParamInfo.vmHostId;
      this.activate++;
    },
    cancelDrawer1() {
      this.drawer1 = false;
    },
    cancelDrawer2() {
      this.drawer2 = false;
    },
    cancelDrawer3() {
      this.drawer3 = false;
    },
    confirmDrawer2() {
      this.VmParamInfo.vmDevicesDiskSourcepath = this.diskpath2;
      ElMessage.success(
        "已选择 Disk 磁盘安装路径: " + this.VmParamInfo.vmDevicesDiskSourcepath
      );
      this.cancelDrawer2();
    },
    async createDomain() {
      this.createDomainLoading = true;
      let resp = await this.postRequestCreateDomain();
      this.createDomainLoading = false;
      if (resp.code == 200) {
        this.activate++;
        ElMessage.success("虚拟机创建成功！");
      } else {
        ElMessageBox.confirm(resp.message, "Error", {
          cancelButtonText: "Cancel",
          type: "error",
        });
      }
    },
    postRequestCreateDomain() {
      return postRequest("/api/vm/createbyxml", this.VmParamInfo);
    },
    clearInitInfo() {
      this.VmParamInfo.vmDevicesDiskDriverType = "";
      this.VmParamInfo.vmDevicesDiskSize = "";
    },
  },
  mounted() {
    this.createDomainLoading = false;
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
});
</script>

<template>
  <el-card class="box-card">
    <div class="create-container">
      <el-steps :active="activate" align-center>
        <el-step
          title="基本信息"
          description="Basic information about the VM"
        />
        <el-step
          title="镜像选择"
          description="Virtual machine image selection"
        />
        <el-step
          title="高级配置"
          description="Advanced configuration information"
        />
        <el-step title="信息确认" description="Confirm the VM information" />
        <el-step title="创建完成" description="Vm states information" />
      </el-steps>

      <el-drawer
        title="ISO/CDROM 镜像路径选择"
        v-model="drawer1"
        :direction="'rtl'"
        size="55%"
      >
        <el-text class="drawer1-path" type="primary"
          >当前路径: {{ this.diskpath1 }}</el-text
        >

        <el-table
          :data="this.resourceInfoListImage"
          border
          stripe
          scrollbar-always-on
          class="drawer1-table"
          highlight-current-row
          @row-dblclick="this.updateFolderPath1"
        >
          <el-table-column property="resourceName" label="name" width="240" />
          <el-table-column property="isDirectory" label="isDir" width="160" />
          <el-table-column
            property="resourceSize"
            label="size (MB)"
            width="160"
          />
          <el-table-column
            property="resourceModifyTime"
            label="time"
            width="210"
          />
        </el-table>

        <div class="drawer1-footer" align="right">
          <el-button @click="cancelDrawer1">cancel</el-button>
        </div>
      </el-drawer>

      <el-drawer
        title="Disk 磁盘 安装路径选择"
        v-model="drawer2"
        :direction="'rtl'"
        size="55%"
      >
        <el-text class="drawer2-path" type="primary"
          >当前路径: {{ this.diskpath2 }}</el-text
        >

        <el-table
          :data="this.resourceInfoListDisk"
          border
          stripe
          scrollbar-always-on
          class="drawer2-table"
          highlight-current-row
          @row-dblclick="this.updateFolderPath2"
        >
          <el-table-column property="resourceName" label="name" width="240" />
          <el-table-column property="isDirectory" label="isDir" width="160" />
          <el-table-column property="resourceSize" label="size" width="160" />
          <el-table-column
            property="resourceModifyTime"
            label="time"
            width="210"
          />
        </el-table>

        <div class="drawer2-footer" align="right">
          <el-button @click="cancelDrawer2">cancel</el-button>
          <el-button type="primary" @click="confirmDrawer2">confirm</el-button>
        </div>
      </el-drawer>

      <el-drawer
        title="现有磁盘镜像选择"
        v-model="drawer3"
        :direction="'rtl'"
        size="55%"
      >
        <el-text class="drawer1-path" type="primary"
          >当前路径: {{ this.diskpath1 }}</el-text
        >

        <el-table
          :data="this.resourceInfoListImage"
          border
          stripe
          scrollbar-always-on
          class="drawer1-table"
          highlight-current-row
          @row-dblclick="this.updateFolderPath1"
        >
          <el-table-column property="resourceName" label="name" width="240" />
          <el-table-column property="isDirectory" label="isDir" width="160" />
          <el-table-column
            property="resourceSize"
            label="size (MB)"
            width="160"
          />
          <el-table-column
            property="resourceModifyTime"
            label="time"
            width="210"
          />
        </el-table>

        <div class="drawer1-footer" align="right">
          <el-button @click="cancelDrawer3">cancel</el-button>
        </div>
      </el-drawer>

      <div v-if="activate === 1">
        <el-row gutter="20">
          <el-col :span="11">
            <el-form class="create-mode1" :model="this.VmParamInfo">
              <el-form-item
                label="虚拟机名称"
                prop="vmName"
                :rules="[{ required: true, message: '提醒: 请输入虚拟机名称' }]"
              >
                <el-input
                  v-model="VmParamInfo.vmName"
                  style="width: 250px"
                  placeholder="请输入虚拟机名称"
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
                <el-select
                  v-model="this.hostOptionDefault"
                  style="width: 250px"
                  @change="this.updateVmHostId"
                >
                  <el-option
                    v-for="item in hostOptions"
                    :key="item.value"
                    :value="item.value"
                  />
                </el-select>
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
            style="margin-left: 38%; margin-bottom: 40px; margin-top: 40px"
          >
            <el-button
              type="primary"
              @click="activate++"
              style="margin-right: 20px; width: 100px"
              >下一步</el-button
            >
            <el-button style="width: 100px">取 消</el-button>
          </el-form-item>
        </el-row>
      </div>

      <div v-if="activate === 2">
        <el-row gutter="20">
          <el-col :span="11">
            <el-form class="create-mode3" :model="this.VmParamInfo">
              <el-form-item label="虚拟机安装途径">
                <el-radio-group
                  v-model="VmParamInfo.vmCreateForm"
                  style="margin-left: 10px"
                >
                  <el-radio label="0" border style="margin-bottom: 20px"
                    >本地安装介质 (IOS Image or CDROM)</el-radio
                  >>
                  <el-radio
                    label="1"
                    border
                    style="margin-bottom: 20px; width: 290px"
                    @click="this.clearInitInfo"
                    >现有磁盘镜像 (Disk Image)</el-radio
                  >>
                  <el-radio label="2" border disabled style="width: 290px"
                    >网络安装 (HTTP, HTTPS or FTP)</el-radio
                  >>
                </el-radio-group>
              </el-form-item>
              <div v-if="VmParamInfo.vmCreateForm == '0'">
                <el-form-item
                  label="ISO/CDROM 镜像路径"
                  prop="vmDevicesDiskImageSourcepath"
                  :rules="[{ required: true, message: '提醒: 请选择镜像路径' }]"
                >
                  <el-input
                    v-model="VmParamInfo.vmDevicesDiskImageSourcepath"
                    style="width: 250px"
                    placeholder="请选择镜像路径"
                  />
                  <el-button
                    type="primary"
                    :icon="Document"
                    round
                    plain
                    style="margin-left: 20px"
                    @click="this.showDrawer1"
                  ></el-button>
                </el-form-item>
                <el-form-item
                  label="Disk 磁盘 安装路径"
                  prop="vmDevicesDiskSourcepath"
                  :rules="[{ required: true }]"
                  :show-message="false"
                >
                  <el-input
                    v-model="VmParamInfo.vmDevicesDiskSourcepath"
                    style="width: 250px; margin-left: 21px"
                    placeholder="请输入磁盘安装路径"
                  />
                  <el-button
                    type="primary"
                    :icon="FolderOpened"
                    round
                    plain
                    style="margin-left: 20px"
                    @click="this.showDrawer2"
                  ></el-button>
                </el-form-item>
                <el-form-item
                  label="Disk 磁盘类型选择"
                  style="margin-left: 12px"
                >
                  <el-radio-group
                    v-model="this.VmParamInfo.vmDevicesDiskDriverType"
                  >
                    <el-radio label="raw" style="margin-left: 30px" />
                    <el-radio label="qcow2" />
                  </el-radio-group>
                </el-form-item>
                <el-form-item
                  label="虚拟机磁盘容量 (G)"
                  style="margin-left: 10px"
                >
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
                    v-model="VmParamInfo.vmDevicesDiskImageSourcepath"
                    style="width: 250px; margin-left: 31px"
                    placeholder="请选择镜像路径"
                  />
                  <el-button
                    type="primary"
                    :icon="Document"
                    round
                    plain
                    style="margin-left: 20px"
                    @click="this.showDrawer3"
                  ></el-button>
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
                  placeholder="默认为空，自动生成"
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
                  <el-select
                    v-model="this.VmParamInfo.vmDevicesDiskType"
                    style="width: 300px; margin-left: 128px"
                    placeholder="请选择 Disk 磁盘设备类型"
                  >
                    <el-option
                      v-for="item in diskTypeOptions"
                      :key="item.value"
                      :value="item.value"
                      :disabled="item.disabled"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item label="Disk Device">
                  <el-radio-group
                    v-model="this.VmParamInfo.vmDevicesDiskDevice"
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
                  />
                </el-form-item>
                <el-form-item label="Disk Target-Bus">
                  <el-select
                    v-model="this.VmParamInfo.vmDevicesDiskTargetBus"
                    style="width: 300px; margin-left: 90px"
                    placeholder="请指定磁盘设备类型"
                  >
                    <el-option
                      v-for="item in diskTargetBusOptions"
                      :key="item.value"
                      :value="item.value"
                    />
                  </el-select>
                </el-form-item>
              </el-form>
            </el-collapse-item>
            <el-collapse-item title="虚拟机 ISO/CDROM 镜像高级配置" name="3">
              <el-form style="margin-left: 22%">
                <el-form-item label="Image Disk Type">
                  <el-select
                    v-model="this.VmParamInfo.vmDevicesDiskImageType"
                    style="width: 300px; margin-left: 90px"
                    placeholder="请选择 Image 磁盘设备类型"
                  >
                    <el-option
                      v-for="item in diskTypeOptions"
                      :key="item.value"
                      :value="item.value"
                      :disabled="item.disabled"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item label="Image Disk Device">
                  <el-radio-group
                    v-model="this.VmParamInfo.vmDevicesDiskImageDevice"
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
                  />
                </el-form-item>
                <el-form-item label="Image Target-Bus">
                  <el-select
                    v-model="this.VmParamInfo.vmDevicesDiskImageTargetBus"
                    style="width: 300px; margin-left: 80px"
                    placeholder="请指定 Image 磁盘设备类型"
                  >
                    <el-option
                      v-for="item in diskTargetBusOptions"
                      :key="item.value"
                      :value="item.value"
                    />
                  </el-select>
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
                  <el-radio-group v-model="this.vnc">
                    <el-radio label="yes" style="margin-left: 55px" />
                    <el-radio label="no" />
                  </el-radio-group>
                </el-form-item>

                <div v-if="this.vnc == 'yes'">
                  <el-form-item label="VNC Port">
                    <el-input
                      v-model="this.VmParamInfo.vmGraphicsVncPort"
                      placeholder="请手动指定 VNC 端口号"
                      style="margin-left: 140px; width: 300px"
                    />
                  </el-form-item>
                </div>
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

      <div v-if="activate === 4" v-loading="this.createDomainLoading">
        <el-row>
          <el-form-item style="margin-left: 31%; margin-top: 20px">
            <vue-json-pretty
              :data="this.VmParamInfo"
              style="width: 740px"
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
              @click="this.createDomain"
              style="width: 100px"
              >确认创建</el-button
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
            ><CircleCheck
          /></el-icon>
        </el-row>
      </div>
    </div>
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
  margin-left: 40px;
  .el-form-item {
    margin-bottom: 25px;
  }
}
.create-mode4 {
  width: 550px;
  margin-top: 80px;
  margin-left: 20px;
  .el-form-item {
    margin-bottom: 28px;
  }
}

.drawer1-path {
  margin-left: 13px;
}
.drawer1-table {
  width: 770px;
  height: 550px;
  margin-top: 13px;
  margin-left: 10px;
}
.drawer2-table {
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
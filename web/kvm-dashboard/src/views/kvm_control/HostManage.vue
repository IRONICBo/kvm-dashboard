<script>
import { defineComponent } from "vue";
import HostManageView from "@/components/kvm_control/HostManageView.vue";

export default defineComponent({
  name: "HostManage",
  components: { HostManageView },
});
</script>

<template>
  <host-manage-view></host-manage-view>
</template>

<style scoped>
</style>
<script>
import { defineComponent } from "vue";
import DomainManageView from "@/components/kvm_control/DomainManageView.vue";

export default defineComponent({
  name: "DomainManage",
  components: { DomainManageView },
});
</script>

<template>
  <domain-manage-view></domain-manage-view>
</template>

<style scoped>
</style>
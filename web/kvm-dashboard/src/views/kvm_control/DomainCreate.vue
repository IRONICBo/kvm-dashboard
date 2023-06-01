<script>
import { defineComponent } from "vue";
import DomainCreateView from "@/components/kvm_control/DomainCreateView.vue";

export default defineComponent({
  name: "DomainCreate",
  components: { DomainCreateView },
});
</script>

<template>
  <domain-create-view></domain-create-view>
</template>

<style scoped>
</style>
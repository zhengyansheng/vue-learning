<template>

<div>
    <CmdbListPage></CmdbListPage>
</div>
<br>

<div class="mb-4">
    <CmdbAddPage ref="InsanceInfo"></CmdbAddPage><br>
</div>

</template>

<script setup lang="ts">

import { ref } from "vue";
import CmdbListPage from "@/components/cmdb/CmdbList.vue"

let InsanceInfo = ref();

</script>

<style scoped>


</style>

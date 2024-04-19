
<template>

<MySalary ref="SalaryInfo"></MySalary><br>
<button @click="addSalary">提交</button>

</template>

<script setup lang="ts">
import { ref } from "vue";
import MySalary from "./components/MySalary.vue"

let SalaryInfo = ref()

function addSalary() {
    console.log(SalaryInfo)
    console.log(SalaryInfo.value)
    SalaryInfo.value.salary+=1
}

</script>

<style scoped>


</style>

<template>
    <el-card>
    <template #header>
      <div class="card-header">
        <span>CMDB数据展示 </span>
        <DialogForm></DialogForm>
        
      </div>
    </template>

    <el-card>
    <el-table :data="responseData" style="width: 100%">
      <el-table-column prop="instanceID" label="实例ID" width="180" />
      <el-table-column prop="hostname" label="主机名" width="180" />
      <el-table-column prop="instanceType" label="机型" width="180" />
      <el-table-column prop="platform" label="云平台" width="180" />
      <el-table-column prop="date1" label="创建时间" width="220" />
      <el-table-column fixed="right" label="操作" width="120">
      <template #default="{ row }">
        <el-button link type="primary" size="small" @click="handleDelete(row)">
          删除
        </el-button>
        <el-button link type="primary" size="small">修改</el-button>
      </template>
    </el-table-column>
    </el-table>
    </el-card>
</el-card>
</template>

<script setup lang="ts">
import axios from 'axios';
import { ref, onMounted } from 'vue';

// import DialogForm from "./components/cmdb/DialogForm.vue"
import DialogForm from "@/components/cmdb/DialogForm.vue";

const responseData = ref('');

onMounted(() => {
    fetchData()
});

const handleDelete = (row) => {
    // 在这里处理删除操作，row参数就是点击行的数据
    console.log("Delete row:", row.instanceID);

    // 定义查询参数
    let timeout = { timeout: 3000 }

    // 在页面加载完成后发起HTTP GET请求
    let url = "http://127.0.0.1:8080/api/v1/cmdb/"+row.instanceID

    // 使用axios.get发起HTTP GET请求，并传递查询参数
    axios.delete(url, {
        ...timeout
    })
    .then(response => {
        // 请求成功，更新响应数据
        let responseData = response.data;
        if (responseData.code === 0) {
            // 处理成功情况
            console.log(responseData.data);
            fetchData()
        } else {
            // 处理失败情况
            console.error("Failed to fetch data:", responseData.message);
        }
    })
    .catch(error => {
        // 请求失败，输出错误信息
        console.error("Failed to fetch data:", error);
    });
};

// 发送 HTTP GET 请求获取数据的函数
const fetchData = () => {

    // 在页面加载完成后发起HTTP GET请求
    let url = "http://127.0.0.1:8080/api/v1/cmdb"

    // 设置超时时间为3秒
    let timeout = { timeout: 3000 }
    axios.get(url, timeout)
    .then(response => {
        // 请求成功，更新响应数据
    //   console.log(response.data.data)
        response = response.data
    //   responseData.value = response.data.data;
        if (response.code === 0) {
            // 处理成功情况
            console.log(response.data)
            responseData.value = response.data
        } else {
            // 处理失败情况
            responseData.value = null
        }
    })
    .catch(error => {
        // 请求失败，输出错误信息
        console.error('Failed to fetch data:', error);
    });

};




</script>

<style scoped>

</style>
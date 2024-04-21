<template>
    <el-card>
    <template #header>
      <div class="card-header">
        <span>CMDB数据展示 </span>
        <DialogForm></DialogForm>
        <el-input
            v-model="input"
            style="width: 240px"
            placeholder="请搜索"
            clearable
            @input="handleSearch"
        />
        
      </div>
    </template>

    <el-card>
        <el-table :data="responseData" style="width: 100%">
        <el-table-column prop="instance_id" label="实例ID" width="180" />
        <el-table-column prop="hostname" label="主机名" width="180" />
        <el-table-column prop="instance_type" label="机型" width="180" />
        <el-table-column prop="platform" label="云平台" width="180" />
        <el-table-column prop="date1" label="创建时间" width="220" />
        <el-table-column fixed="right" label="操作" width="200">
        <template #default="{ row }">
            <!-- <el-button plain type="warning" size="small" @click="showModal = true">修改</el-button> -->
            <el-button plain type="warning" size="small" @click="handleUpdate(row)">修改</el-button>
            <el-dialog v-model="showModal" title="修改主机" width="500">
                <!-- 子页面 -->
                <ChildUpdate  :jsonData="jsonData"></ChildUpdate>
            </el-dialog>

            <el-button type="danger" size="small" @click="handleDelete(row)">删除</el-button>
        </template>
        </el-table-column>
        </el-table>
    </el-card>


</el-card>
</template>

<script setup lang="ts">

import axios from 'axios';
import { ref, onMounted } from 'vue';

import DialogForm from "@/components/cmdb/DialogForm.vue";
import ChildUpdate from "@/components/cmdb/ChildUpdate.vue";


// 在父组件中引用子组件，并访问被暴露的属性
const childRef = ref();
const responseData = ref();
const input = ref('')
const showModal = ref(false);
const jsonData = ref();

onMounted(() => {
    fetchData("")
});

function handleUpdate(row: Map<string, any>) {
    showModal.value = true
    console.log("Update row:", row.instance_id);
    singleQuery(row.instance_id)

}

function handleSearch() {
    console.log(input.value)
    fetchData(input.value)
}

const handleDelete = (row: Map<string, any>) => {
    // 在这里处理删除操作，row参数就是点击行的数据
    console.log("Delete row:", row.instance_id);

    // 定义查询参数
    let timeout = { timeout: 3000 }

    // 在页面加载完成后发起HTTP GET请求
    let url = "http://127.0.0.1:8080/api/v1/cmdb/"+row.instance_id

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
const fetchData = (id: string) => {
    // 在页面加载完成后发起HTTP GET请求
    let url = "http://127.0.0.1:8080/api/v1/cmdb"  // 将url的声明移动到外部
    let is_multiple_data = true

    if (id !== "") {
        is_multiple_data = false
        url = url+"/"+id
    }
    console.log(url)
    

    // 设置超时时间
    let timeout = { timeout: 3000 }

    // 发起GET请求
    axios.get(url, timeout).then(response => {
        if (response.data.code === 0) {
            // 处理成功情况
            responseData.value = response.data.data
            if (is_multiple_data == false) {
                responseData.value = [response.data.data]
            }
           
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

const singleQuery = (id: string) => {
    // 在页面加载完成后发起HTTP GET请求
    let url = "http://127.0.0.1:8080/api/v1/cmdb/"+id
    console.log(url)
    

    // 设置超时时间
    let timeout = { timeout: 3000 }

    // 发起GET请求
    axios.get(url, timeout).then(response => {
        if (response.data.code === 0) {
            // 处理成功情况
            jsonData.value = response.data.data
            jsonData.value = {"instance_id":response.data.data.instance_id}
            console.log("single query", jsonData.value)
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
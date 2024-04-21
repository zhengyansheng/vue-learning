<template>
     <el-button plain @click="dialogFormVisible = true">
        添加
     </el-button>

    <el-dialog v-model="dialogFormVisible" title="添加主机" width="500">
        <!-- 子页面 -->
        <CmdbAddPage ref="childRef"></CmdbAddPage>

        <el-form-item>
            <el-button type="primary" @click="onSubmit">提交</el-button>
            <el-button @click="onCancel">取消</el-button>
        </el-form-item>
    </el-dialog>

</template>


<script lang="ts" setup>

import { ref, reactive } from "vue";
import axios from "axios";
import { ElMessage } from 'element-plus'

import CmdbAddPage from "@/components/cmdb/CmdbAdd.vue"

const dialogFormVisible = ref(false)

// 在父组件中引用子组件，并访问被暴露的属性
const childRef = ref();

// 提交书籍
const onSubmit = () =>{
    createPost(childRef.value.form)
}

// 模态窗取消
function onCancel() {
  dialogFormVisible.value = false
}

// 定义创建Post的函数
const createPost = async (data) => {
  try {
    let url = "http://127.0.0.1:8080/api/v1/cmdb"

     // 设置超时时间为3秒
    let timeout = { timeout: 3000 }
    
    const { response } = await axios.post(url, data, timeout);
    if (response.code === 0) {
      // 处理成功情况
      ElMessage({
        message: '提交成功.',
        type: 'success',
      })
    } else {
      // 处理失败情况
      ElMessage.error('Oops, 提交失败:', response.message)
    }
  } catch (error) {
    // 异常
    ElMessage({
      message: 'Oops, this is a error message:' + error,
      type: 'error',
      plain: true,
    })
  }
};



</script>

<style scoped>
</style>

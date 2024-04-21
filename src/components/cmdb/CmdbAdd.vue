
<template>
  <div >
    <el-form :model="form" label-width="auto" style="max-width: 600px">
      <el-form-item label="实例 ID">
        <el-input v-model="form.instanceID" />
      </el-form-item>
      <el-form-item label="主机名">
        <el-input v-model="form.hostname" />
      </el-form-item>
      <el-form-item label="机型">
        <el-select v-model="form.instanceType" placeholder="请选择你的机型">
        <el-option label="1 vCPU | 2 GiB" value="ecs.t5-lc1m2.small" />
        <el-option label="2 vCPU | 4 GiB" value="ecs.c7.large" />
        <el-option label="4 vCPU | 8 GiB" value="ecs.c7.xlarge" />
        <el-option label="8 vCPU | 16 GiB" value="ecs.c7.2xlarge" />
      </el-select>
      </el-form-item>
      <el-form-item label="创建时间">
        <el-col :span="11">
          <el-date-picker
            v-model="form.date1"
            type="date"
            placeholder="Pick a date"
            style="width: 100%"
          />
        </el-col>
        <el-col :span="2" class="text-center">
          <span class="text-gray-500">-</span>
        </el-col>
        <el-col :span="11">
          <el-time-picker
            v-model="form.date2"
            placeholder="Pick a time"
            style="width: 100%"
          />
        </el-col>
      </el-form-item>
      <el-form-item label="云平台">
        <el-radio-group v-model="form.platform">
          <el-radio value="aliyun">阿里云</el-radio>
          <el-radio value="tengxunyun">腾讯云</el-radio>
          <el-radio value="huaweiyun">华为云</el-radio>
          <el-radio value="aws">aws</el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts" setup>

import { ref  } from 'vue';
import { reactive } from 'vue';

// 使用 ref 创建一个变量控制窗口的显示和隐藏
const dialogFormVisible = ref(true);

// do not use same name with ref
const form = reactive({
  instanceID: '',
  hostname: '',
  instanceType: '',
  date1: '',
  date2: '',
  platform: '',
})



// 将 数据 暴露给父组件
defineExpose({
  form, 
});

</script>


<style scoped>

.center-content {
    flex-wrap: wrap; /* 添加这个属性以允许换行 */
    justify-content: center;
    align-items: center;
    background-color: #f0f0f0; /* 设置背景颜色，以便看到居中效果 */
}

</style>

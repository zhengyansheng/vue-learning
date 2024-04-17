<template>
  <div>
    用户: <input v-model="userName"><br>
    薪水: <input type="number" v-model="salary"><br>
    <button @click="addSalary">提交</button>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        userName: "张三",
        salary: 1000
      }
    },
    methods: {
      addSalary() {
        this.salary++
      }
    },

  }

</script>

<style scoped>

</style>


<template>
  <div>
    用户: <input v-model="userName">{{ userName}}<br>
    薪水: <input type="number" v-model="salary">{{ salary }}<br>
    <button @click="addSalary">提交</button>
    <br />

    <div>
        <h2>个人信息</h2>

        <p> 年龄: 
          <input type="number" v-model="userInfo.age">
        </p>
        <p>
          性别: 
          <input type="radio" value="1" v-model="userInfo.sex">男
          <input type="radio" value="0" v-model="userInfo.sex">女
        </p>
        <p>岗位: 
          <select v-model="userInfo.job">
            <option value="dev">开发</option>
            <option value="test">测试</option>
            <option value="op">运维</option>
          </select>
        </p>
        <p>技术: 
          <span v-for="skill in userInfo.skills" v-bind:key="skill" class="skill-item">{{ skill }}</span>
        </p>
        <p>新技术: 
          <input v-model="userInfo.skill">
          <button @click="addSkill">新增</button>
        </p>
        <p>个人信息汇总: <br>{{ userInfo }}</p>
        <p>
          <ul>
            <li v-for="(value, key) in userInfo" :key="key">{{ key }} {{value}}</li>
          </ul>
        </p>
    </div>

  </div>
</template>

<script>
// import { userInfo } from 'os';

  export default {
    // 模版数据
    data() {
      return {
        userName: "张三",
        salary: 1000,
        userInfo: {
          age:18,
          sex:1,
          job: "test",
          skills: ["python", "java"],
          skill: "",
        }
      }
    },
    // 模版方法
    methods: {
      addSalary() {
        this.salary++
      },
      addSkill() {
        if(this.userInfo.skill){
          this.userInfo.skills.push(this.userInfo.skill)
        }
      }
    },
  }

</script>

<style scoped>

.skill-item {
  background-color:bisque;
  margin-right: 10px; /* 设置间距 */
}

</style>

<template>
  <el-container>
    <el-main style="margin-left: 30%;margin-right: 30%">
      <el-form :model="form">
        <el-form-item label="Username">
          <el-input v-model="form.username" placeholder="Please input username"/>
        </el-form-item>

        <el-form-item label="Password">
          <el-input show-password v-model="form.password" placeholder="Please input password"/>
        </el-form-item>

        <el-form-item style="margin-left: 14%">
          <el-button type="primary" @click="onSubmit">Login</el-button>
          <el-button>Cancel</el-button>
        </el-form-item>
      </el-form>
    </el-main>
  </el-container>
</template>

<script lang="ts" setup>
import {reactive} from 'vue'
import {ElMessage} from "element-plus";
import axios from "axios";



// do not use same name with ref
const form = reactive({
  username: '',
  password: '',
})

const data = reactive({
  name: 'test'
})

const onSubmit = () => {
  console.log('submit!')
  let postData = {
    'username': form.username,
    'password': form.password
  }
  axios.post(`/api/v1/sign-in`, postData).then((res) => {
    if (res.data.code === 0) {
      localStorage.setItem('token', res.data.data)
    } else {
      ElMessage.error(res.data.msg)
    }
  })
}


</script>

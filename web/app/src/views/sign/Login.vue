<template>
  <el-form :model="form" label-width="120px">
    <el-form-item label="Activity name">
      <el-input v-model="form.email"/>
    </el-form-item>
    <el-form-item label="Activity name">
      <el-input v-model="form.password"/>
    </el-form-item>

    <el-form-item label="Instant delivery">
      <el-switch v-model="form.delivery"/>
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="onSubmit">Create</el-button>
      <el-button>Cancel</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>

import {reactive} from 'vue'
import router from "@/router";
import tokenApi from "@/api/modules/token";
import userApi from "@/api/modules/user";
import userStore from "@/store/user";

const form = reactive({
  email: '',
  password: '',
  delivery: false,

})

const onSubmit = () => {
  tokenApi.postToken(form).then((v) => {
    if (v.code == 0) {
      userApi.getUserInfo().then((v) => {
        if (v.code == 0) {
          userStore().setUserInfo(v.data)
          console.log(userStore().getUserInfo().email)
        }
      })
      ElMessage.success(`${'登陆成功'}`)
      router.push({name: 'home'})
    } else {
      ElMessage.error(`${v.msg}`)
    }
  }).catch((err) => {
    ElMessage.error(err.message)
  })
  console.log()
}

onMounted(() => {
  console.log(userStore().getUserInfo())
  if (Object.keys(userStore().getUserInfo()).length > 0) {
    router.push({name: 'home'})
  }
})
</script>
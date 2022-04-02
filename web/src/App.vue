<template>
  <el-menu mode="horizontal" @select="handleSelect">
    <el-menu-item index="/">Alchemy Furnace</el-menu-item>
    <el-menu-item index="/about">About</el-menu-item>
  </el-menu>
  <router-view/>
</template>

<script lang="ts" setup>
import {useRouter} from 'vue-router'

const router = useRouter();

const handleSelect = (key: string) => {
  router.push({path: key})
}
</script>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()
</script>

<template>
  <div class="home">
    <h1>Ag Araclari</h1>
    <div class="butonlar">
      <button @click="router.push('/subnetting')">IP Subnetting</button>
      <button @click="router.push('/portchecker')">Port Checker</button>
    </div>
  </div>
</template>

<style scoped>
.home {
  text-align: center;
  margin-top: 60px;
}
.butonlar {
  display: flex;
  justify-content: center;
  gap: 16px;
  margin-top: 24px;
}
button {
  padding: 16px 32px;
  font-size: 16px;
  cursor: pointer;
}
</style>
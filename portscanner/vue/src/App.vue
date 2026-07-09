<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <nav>
    <RouterLink to="/">Ana Sayfa</RouterLink> |
    <RouterLink to="/subnetting">IP Subnetting</RouterLink> |
    <RouterLink to="/portchecker">Port Checker</RouterLink>
  </nav>

  <RouterView />
</template>

<style>
nav {
  padding: 16px;
  text-align: center;
  border-bottom: 1px solid #ccc;
}
nav a {
  margin: 0 8px;
  text-decoration: none;
  color: #333;
}
</style>
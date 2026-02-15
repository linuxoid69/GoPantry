<template>
  <div>
    <h1>Данные с бэкенда</h1>
    <DataTable :api-url="apiUrl" />
  </div>
</template>

<script setup>
const apiUrl = 'https://jsonplaceholder.typicode.com/users' // Пример внешнего API

// Или если ваш бэкенд локально:
// const apiUrl = 'http://localhost:3001/api/users'
</script>

<template>
  <div class="table-container">
    <!-- Состояние загрузки -->
    <div v-if="pending" class="loading">
      Загрузка пользователей...
    </div>

    <!-- Ошибка -->
    <div v-else-if="error" class="error">
      Ошибка загрузки: {{ error.message }}
    </div>

    <!-- Таблица с данными -->
    <table v-else-if="users && users.length" class="data-table">
      <thead>
        <tr>
          <th>ID</th>
          <th>Имя</th>
          <th>Username</th>
          <th>Email</th>
          <th>Город</th>
          <th>Телефон</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td>{{ user.id }}</td>
          <td>{{ user.name }}</td>
          <td>{{ user.username }}</td>
          <td>{{ user.email }}</td>
          <!-- Обратите внимание, как мы получаем город из вложенного объекта address -->
          <td>{{ user.address.city }}</td>
          <td>{{ user.phone }}</td>
        </tr>
      </tbody>
    </table>

    <!-- Нет данных -->
    <div v-else class="no-data">
      Нет данных для отображения
    </div>
  </div>
</template>

<script setup>
// Используем useFetch для получения данных
const { data: users, pending, error } = await useFetch('https://jsonplaceholder.typicode.com/users')

// Определяем заголовки таблицы (автоматически или вручную)
const headers = computed(() => {
    if (data.value && data.value.length > 0) {
        // Автоматически создаём заголовки из первого элемента
        return Object.keys(data.value[0]).map(key => ({
            key,
            label: key.charAt(0).toUpperCase() + key.slice(1) // Первая буква заглавная
        }))
    }
    return []
})

// Функция для ручного обновления
const refreshData = () => {
    refresh()
}
</script>

<style scoped>
.table-container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 20px;
    background: white;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.data-table {
    width: 100%;
    border-collapse: collapse;
    margin: 20px 0;
}

.data-table th {
    background-color: #00C58E;
    color: white;
    padding: 12px;
    text-align: left;
    font-weight: bold;
}

.data-table td {
    padding: 10px 12px;
    border-bottom: 1px solid #ddd;
}

.data-table tr:hover {
    background-color: #f5f5f5;
}

.loading,
.error,
.no-data {
    text-align: center;
    padding: 40px;
    font-size: 18px;
}

.error {
    color: #ff4444;
}

.refresh-btn {
    background-color: #00C58E;
    color: white;
    border: none;
    padding: 10px 20px;
    border-radius: 4px;
    cursor: pointer;
    font-size: 16px;
    margin-top: 20px;
}

.refresh-btn:hover {
    background-color: #00a076;
}
</style>

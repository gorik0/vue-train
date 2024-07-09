<script setup>
import Header from './Components/Header.vue'
import CardList from './Components/CardList.vue'
import Drawer from './Components/Drawer.vue'
</script>

<template>
  <Drawer />
  <div class="bg-white w-4/5 m-auto mt-10 rounded shadow-xl">
    <Header />
    <div class="p-10">
      <h2 class="text-3xl bold">Кроссовки</h2>

      <CardList />
    </div>
  </div>
</template>

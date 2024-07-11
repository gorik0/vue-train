<script setup>
import { inject } from 'vue'
import CardItem from './CardItem.vue'
const { itemsInBasket } = inject('basket')




</script>
<template>
  <div class="flex flex-col gap-4">
    <CardItem v-for="item in itemsInBasket" :key="item.id" :title="item.title" :price="item.price" :img="item.img" />
  </div>
</template>

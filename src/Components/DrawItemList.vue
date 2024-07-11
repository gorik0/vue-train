<script setup>
import { inject } from 'vue'
import CardItem from './CardItem.vue'
const { itemsInBasket, addToBasket } = inject('basket')




</script>
<template>
  <div class="flex flex-col gap-4" v-auto-animate>
    <CardItem v-for="item in itemsInBasket" :addToBasket="() => addToBasket(item)" :key="item.id" :title="item.title"
      :price="item.price" :img="item.img" />
  </div>
</template>

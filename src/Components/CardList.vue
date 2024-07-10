<script setup>
defineProps({
  items: Array
})
import Card from './Card.vue'

const addToBasket = () => {
  console.log('add to basket')
}
</script>

<template>
  <div class="grid grid-cols-2 gap-2">
    <Card v-for="item in items" :key="item.id" :title="item.title" :price="item.price" :img="item.img"
      :isAddedToFavorite="item.isFavorite" :isAddedToBasket="true" :addToBasket="addToBasket" />
  </div>
</template>

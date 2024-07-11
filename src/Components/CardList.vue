<script setup>


import { inject } from 'vue';

const { addToBasket } = inject('basket')
const emit = defineEmits(['addToFavorites'])


defineProps({
  items: Array
})
import Card from './Card.vue'




</script>

<template>
  <div class="grid grid-cols-2 gap-2">
    <Card v-for="item in items" :key="item.id" :title="item.title" :price="item.price" :img="item.img"
      :isAddedToFavorite="item.isFavorite" :isAddedToBasket="item.isAddedToBasket"
      :addToFavorite="() => emit('addToFavorites', item)" :addToBasket="() => addToBasket(item)" />
  </div>
</template>

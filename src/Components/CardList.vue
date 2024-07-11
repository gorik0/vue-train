<script setup>


import { inject, onMounted } from 'vue';



defineProps({
  items: Array,

  isFavoriteList: Boolean
})
const { addToBasket } = inject('basket')
const emit = defineEmits(['addToFavorites'])



import Card from './Card.vue'



</script>

<template>
  <div class="grid grid-cols-2 gap-2" v-auto-animate>
    <Card v-for="item in items" :key="item.id" :title="item.title" :price="item.price" :img="item.img"
      :isAddedToFavorite="item.isFavorite" :isAddedToBasket="item.isAddedToBasket"
      :addToFavorite="isFavoriteList ? null : () => emit('addToFavorites', item)"
      :addToBasket="() => addToBasket(item)" />
  </div>
</template>

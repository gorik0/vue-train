<script setup>
import Card from './Card.vue'

const addToBasket = () => {
  console.log('add to basket')
}
</script>

<template>
  <div class="grid grid-cols-2 gap-2">
    <Card
      title="Nike pro airBASS"
      :price="1000"
      :img="'/sneakers/sneakers-1.jpg'"
      :isAddedToFavorite="true"
      :isAddedToBasket="true"
      :addToBasket="addToBasket"
    />
    <Card />
    <Card />
    <Card />
  </div>
</template>

<script setup>
import { inject } from "vue"

defineProps({
  title: String,
  price: Number,
  img: String,
  isAddedToBasket: Boolean,
  isAddedToFavorite: Boolean,
  addToBasket: Function,
  addToFavorite: Function
})

</script>

<template>
  <div
    class="relative border-4 border-slate-300 m-10 p-5 rounded-3xl cursor-pointer ease-out transition hover:-translate-y-3 hover:shadow-2xl">
    <img class="absolute top-3 left-3" @click="addToFavorite" :src="isAddedToFavorite ? '/like-2.svg' : '/like-1.svg'"
      alt="like" />
    <img :src="img" alt="sn-1" />
    <p>{{ title }}</p>

    <div class="flex justify-between mt-5">
      <span>Цена</span>
      <b>{{ price }} rub.</b>
      <img @click="addToBasket" :src="!isAddedToBasket ? '/plus.svg' : '/checked.svg'" alt="add_to" />
    </div>
  </div>
</template>

<script setup>
import { inject, onMounted, reactive, provide, ref, watch, computed } from 'vue'

import axios from 'axios'
import Header from './Components/Header.vue'

import Drawer from './Components/Drawer.vue'



// :::: OBSERVABLES props ::::

const itemsInBasket = ref([])

const makeOrderInProccess = ref(false)


const tax = computed(() => {
  return priceTotal.value * 0.2
})

const priceTotal = computed(() => {
  return itemsInBasket.value.reduce((acc, item) => acc + item.price, 0)
})

const isDrawerOpen = ref(false)











// ::: HANDLER user utils actions ::::


const addToBasket = (item) => {

  console.log(`"Add to basket!" ${item.isAddedToBasket}`);

  if (!item.isAddedToBasket) {
    item.isAddedToBasket = true
    itemsInBasket.value.push(item)
  } else {

    item.isAddedToBasket = false
    itemsInBasket.value = itemsInBasket.value.filter((itemInBasket) => itemInBasket.id !== item.id)
  }
  console.log(`"Add to basket!" ${item.isAddedToBasket}`);

}

const openDrawer = () => {
  console.log("open");
  isDrawerOpen.value = true
}
const closeDrawer = () => {

  isDrawerOpen.value = false
}


const makeOrder = async () => {
  const url = `http://localhost:8080/sneakers/order`
  console.log(makeOrderInProccess);
  try {
    makeOrderInProccess.value = true
    const { data } = await axios.post(url, {
      items: itemsInBasket.value,
      price: priceTotal.value
    })

    for (const item of itemsInBasket.value) {
      item.isAddedToBasket = false
    }
    itemsInBasket.value = []
  } catch (err) {
    console.log(err);
  } finally {
    makeOrderInProccess.value = false
  }
}






// ::: MAKE QUERY ::::








// ::::PROVIDERS ::::

provide('basket', { openDrawer, closeDrawer, addToBasket, itemsInBasket, makeOrder })


</script>

<template>
  <Drawer v-if="isDrawerOpen" :total="priceTotal" :tax="tax" @makeOrder="makeOrder"
    :makeOrderInProccess="makeOrderInProccess" />
  <div class="bg-white w-4/5 m-auto mt-10 rounded shadow-xl">
    <Header @open-drawer="openDrawer" :price="priceTotal" />
    <div class="p-10">

      <!-- :HOME.vue -->
      <router-view></router-view>

    </div>
  </div>
</template>

<script setup>
import { inject, onMounted, reactive, provide, ref, watch } from 'vue'

import axios from 'axios'
import Header from './Components/Header.vue'
import CardList from './Components/CardList.vue'
import Drawer from './Components/Drawer.vue'



// :::: OBSERVABLES props ::::

const items = ref([])

const itemsInBasket = ref([])

const isDrawerOpen = ref(false)

const filters = reactive({
  sortBy: '',
  title: ''
})





// ::: FETCH DATA  functions  :::

const fetchSneakers = async () => {

  const url = `http://localhost:8080/sneakers`
  const { data } = await axios.get(url, { params: filters })
  items.value = data.map((item) => ({

    ...item,
    isFavorite: false,
    favoriteId: null,
    isAddedToBasket: false
  }))
}

const fetchSneakersFavorites = async () => {
  const url = `http://localhost:8080/sneakers/favorites`
  const { data: favorites } = await axios.get(url)
  console.log(favorites);
  items.value = items.value.map((item) => {
    const isFavorite = favorites.find((favorite) => favorite.parentId === item.id)

    if (isFavorite) {
      return {
        ...item,
        isFavorite: true,
        favoriteId: isFavorite.id
      }
    } else {
      return item;
    }
  })
}





// :::: HANDLER user filter actions ::::
function makeSortBy(event) {
  filters.sortBy = event.target.value
}
function makeQueryBy(event) {

  filters.title = event.target.value
}



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

const addToFavorites = async (item) => {
  item.isFavorite = !item.isFavorite


  if (item.isFavorite) {

    const url = `http://localhost:8080/sneakers/favorites/add`

    try {

      const { data } = await axios.post(url, {
        parentId: item.id
      })

      item.favoriteId = data.id
      item.isFavorite
      console.log(data);
    } catch (err) {
      console.log(err);
    }
  } else {

    // ::: DELETE FAVORITE  ::::


    const url = `http://localhost:8080/sneakers/favorites/delete`

    try {

      await axios.get(url, {
        params: {
          id: item.favoriteId
        }
      })
    } catch (err) {
      console.log(err);
    }
  }


}





// ::: MAKE QUERY ::::
onMounted(async () => {
  await fetchSneakers()
  await fetchSneakersFavorites()
})

watch(filters, fetchSneakers)


// ::::PROVIDERS ::::

provide('basket', { openDrawer, closeDrawer, addToBasket, itemsInBasket })

</script>

<template>
  <Drawer v-if="isDrawerOpen" />
  <div class="bg-white w-4/5 m-auto mt-10 rounded shadow-xl">
    <Header @open-drawer="openDrawer" />
    <div class="p-10">
      <div class="flex items-center justify-between">
        <h2 class="text-3xl bold">Кроссовки</h2>

        <div class="flex gap-5">
          <select @change="makeSortBy" class="outline-none border-2 rounded p-1">
            <option value="title">ПО title</option>
            <option value="price">ПО price</option>
            <option value="-price">ПО -price</option>
          </select>
          <div class="flex gap-3">
            <img src="/search.svg" alt="search" />
            <input @input="makeQueryBy" type="text" class="border-b focus:border-slate-400 outline-none"
              placeholder="INPUT" />
          </div>
        </div>
      </div>

      <CardList :items="items" @addToFavorites="addToFavorites" />
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref, watch } from 'vue'

import axios from 'axios'
import Header from './Components/Header.vue'
import CardList from './Components/CardList.vue'
import Drawer from './Components/Drawer.vue'

const makeSneakersFromServer = (async) => {
  const url = `http://localhost:8080/sneakers`
  console.log(filters)
  axios.get(url, { params: filters }).then((response) => {
    items.value = response.data
    console.log(items.value)
  })
}
const items = ref([])

const filters = reactive({
  sortBy: '',
  title: ''
})

function makeSortBy(event) {
  filters.sortBy = event.target.value
}
function makeQueryBy(event) {
  filters.title = event.target.value
}
onMounted(makeSneakersFromServer)
watch(filters, makeSneakersFromServer)
</script>

<template>
  <!-- <Drawer /> -->
  <div class="bg-white w-4/5 m-auto mt-10 rounded shadow-xl">
    <Header />
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
            <input
              @input="makeQueryBy"
              type="text"
              class="border-b focus:border-slate-400 outline-none"
              placeholder="INPUT"
            />
          </div>
        </div>
      </div>

      <CardList :items="items" />
    </div>
  </div>
</template>

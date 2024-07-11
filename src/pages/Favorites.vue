

<script setup>
import CardList from '../Components/CardList.vue';
import { inject, onMounted, ref } from "vue";
import axios from 'axios';


const items = ref([])


const fetchFavorites = async () => {
    try {
        const url = `http://localhost:8080/sneakers/favorites`

        const { data } = await axios.get(url)

        const { data: allItems } = await axios.get('http://localhost:8080/sneakers')
        items.value = data.map((item) => {
            const itemWithInfo = allItems.find((itemInArray) => itemInArray.id === item.parentId)

            return itemWithInfo
        })


    } catch (err) {
        console.log(err);
    }
}


onMounted(() => {

    fetchFavorites()
})

const addToFavorites = () => { }


</script>



<template>
    <div class="flex items-center justify-between">
        <h2 class="text-3xl bold">Кроссовки</h2>

    </div>

    <CardList :items="items" @addToFavorites="addToFavorites" :isFavoriteList="true" />
</template>

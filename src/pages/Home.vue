<script setup >
import { inject, onMounted, reactive, provide, ref, watch, computed } from 'vue'

import axios from 'axios'
import CardList from '../Components/CardList.vue'


const filters = reactive({
    sortBy: '',
    title: ''
})

const items = ref([])



const { addToBasket, itemsInBasket } = inject('basket')




// :::: HANDLER user filter actions ::::

function makeSortBy(event) {
    filters.sortBy = event.target.value
}
function makeQueryBy(event) {

    filters.title = event.target.value
}



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

onMounted(async () => {
    // :::: GET items from localstorage

    const fromLocalStorage = localStorage.getItem("basket")

    itemsInBasket.value = fromLocalStorage ? JSON.parse(fromLocalStorage) : []
    itemsInBasket.value


    await fetchSneakers()
    await fetchSneakersFavorites()


    items.value = items.value.map((item) => ({
        ...item,
        isAddedToBasket: itemsInBasket.value.some((basketItem) => (basketItem.id === item.id))
    }))


})





// ::: WATCHERS ::: 
watch(filters, fetchSneakers)

watch(itemsInBasket, () => {

    localStorage.setItem("basket", JSON.stringify(itemsInBasket.value))

}, {
    deep: true
})



</script>


<template>
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
</template>

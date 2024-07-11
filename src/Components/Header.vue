<script setup>


const emit = defineEmits(['open-drawer'])


defineProps({
  price: Number
})






</script>

<template>
  <div class="bg-white">
    <header class="flex justify-between border-b p-5">
      <router-link to="/">
        <div class="flex items-center gap-4">


          <img src="/logo.png" alt="Logo" />
          <h2>Vue SNICKERS</h2>
          <p>Магазин кроссовок</p>
        </div>
      </router-link>

      <ul class="flex gap-10 center">
        <li @click="() => emit('open-drawer')" class="flex items-center gap5 cursor-pointer">
          <img src="/cart.svg" alt="cart" />
          <b>Корзина {{ price }} руб.</b>
        </li>
        <li class="flex items-center gap5 cursor-pointer">
          <router-link to="/favorites">

            <img src="/heart.svg" alt="heart" />
            <b>Закладки</b>

          </router-link>
        </li>
        <li class="flex items-center gap5 cursor-pointer">
          <img src="/profile.svg" alt="profile" />
          <b>Профиль</b>
        </li>
      </ul>
    </header>
  </div>
</template>

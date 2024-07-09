<script setup></script>

<template>
  <div class="bg-white">
    <header class="flex justify-between border-b p-5">
      <div class="flex items-center gap-4">
        <img src="/logo.png" alt="Logo" />
        <h2>Vue SNICKERS</h2>
        <p>Магазин кроссовок</p>
      </div>

      <ul class="flex gap-10 center">
        <li class="flex items-center gap5 cursor-pointer">
          <img src="/cart.svg" alt="cart" />
          <b>Корзина</b>
        </li>
        <li class="flex items-center gap5 cursor-pointer">
          <img src="/heart.svg" alt="heart" />
          <b>Закладки</b>
        </li>
        <li class="flex items-center gap5 cursor-pointer">
          <img src="/profile.svg" alt="profile" />
          <b>Профиль</b>
        </li>
      </ul>
    </header>
  </div>
</template>

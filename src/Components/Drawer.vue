<script setup>
import DrawerHead from './DrawerHead.vue'
import CardItem from './CardItem.vue'
</script>
<template>
  <div class="fixed w-full h-full bg-black/50 z-10">
    <div class="fixed w-2/5 p-7 right-5 z-20 h-full bg-white">
      <DrawerHead />

      <CardItem />
      <CardItem />
      <CardItem />
      <CardItem />
      <CardItem />
    </div>
  </div>
</template>

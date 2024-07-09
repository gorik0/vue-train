<script setup>
import DrawerHead from './DrawerHead.vue'
import DrawItemList from './DrawItemList.vue'
</script>
<template>
  <div class="fixed w-full h-full bg-black/50 z-10">
    <div class="fixed w-2/5 p-7 right-5 z-20 h-full bg-white">
      <DrawerHead />

      <DrawItemList />

      <div class="flex flex-col gap-5 mt-10">
        <div class="flex">
          <b>1000 rubs</b>
          <div class="flex-1 border-10 border-dashed border-b mx-2"></div>
          <span>ИТОГО</span>
        </div>
        <div class="flex justify-between">
          <b>1000 rubs</b>
          <div class="flex-1 border-10 border-dashed border-b mx-2"></div>
          <span>НАЛОГ</span>
        </div>
        <button
          disabled="true"
          class="transition disabled:opacity-50 cursor-pointer w-full bg-black opacity-80 hover:opacity-100 rounded-xl text-white p-3 rounded"
        >
          Оформить
        </button>
      </div>
    </div>
  </div>
</template>

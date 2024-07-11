
<script setup>

defineProps({

  title: String,
  price: Number,
  img: String


})

</script>


<template>
  <div class="flex border p-5 gap-3 rounded-xl border-slate-300 mt-2">
    <img class="mt-0 w-16 h-16" :src="img" alt="" />
    <div class="flex flex-col mt-3">
      <p>{{ title }}</p>
      <div class="flex justify-between">
        <b>{{ price }} rub</b>
        <img class="opacity-30 hover:opacity-100 hover:translate-x-3 transition cursor-pointer" src="/close.svg"
          alt="delete_item" />
      </div>
    </div>
  </div>
</template>

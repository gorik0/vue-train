<!-- ::: Compositional API (OLD) -->

<script setup>
defineProps({
  model: String,
  fatAmount: Number
})
</script>

<template>
  <div class="kefir">
    <h2>KEFIR ::::</h2>
    <p>model :::: {{ model }}</p>
    <p>FAT :::: {{ fatAmount }}</p>
  </div>
</template>

<style scoped>
.kefir {
  border: 2px solid darkblue;
  padding: 10px;
}
</style>
<!-- ::: OPtional API (OLD) -->
<!-- <script>
export default {
  data() {
    return {
      count: 0
    }
  },
  methods: {
    incr() {
      this.count++
    }
  }
}
</script> -->

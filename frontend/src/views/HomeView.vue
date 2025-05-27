<script setup>
import Button from "primevue/button";
import Sidebar from "primevue/sidebar";
import { ref } from "vue";

const visible = ref(false);
</script>

<template>
  <main>
    <h1>Hello World</h1>
    <Button label="Click me" @click="visible = true" />
    <Sidebar v-model:visible="visible" position="right">
      <h1>Hello World</h1>
    </Sidebar>
  </main>
</template>

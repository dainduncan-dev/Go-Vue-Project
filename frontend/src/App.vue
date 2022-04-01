<template>
  <ObjectData />
</template>

<script>
import ObjectData from './components/ObjectData.vue'

export default {
  name: 'App',
  components: {
    ObjectData
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  color: #2c3e50;
  display: flex;
  align-items: center;
  justify-content: center;
}
html {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background: rgb(0,0,0);
  background: linear-gradient(0deg, rgba(0,0,0,1) 14%, rgba(37,52,221,0.9136029411764706) 55%, rgba(9,154,64,1) 94%);
}
</style>

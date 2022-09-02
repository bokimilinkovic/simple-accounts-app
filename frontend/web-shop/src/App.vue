<template>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.0.0/dist/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
  <Header />
  <div id="nav" class="center-block text-center px-4 py-5 p ">
    <h3>My App</h3>
      <router-link class="col" class-active="active" exact  to="/login">Login</router-link> 
      <router-link class="col" class-active="active" exact to="/counter">Counter</router-link>
      <router-link class="col" class-active="active" exact to="/movies">Movies</router-link>

  </div>
  <router-view />
  <Footer />
</template>

<script>
import Header from "./views/Header.vue"
import Footer from "./views/Footer.vue"

export default {
    setup() {
        return {};
    },
    watch: {},
    components: { Footer, Header }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

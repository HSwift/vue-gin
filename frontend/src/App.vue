<template>
  <div id="app">
    <keep-alive>
      <component v-bind:is="currentTabComponent" v-on:switch="switchTab"></component>
    </keep-alive>
  </div>
</template>
<script>
import Login from "./components/Login.vue";
import Register from "./components/Register.vue";

export default {
  name: "app",
  data() {
    return {
      tab: Login
    };
  },
  methods: {
    switchTab: function(a) {
      if (a == "login") this.tab = Login;
      else this.tab = Register;
    }
  },
  computed: {
    currentTabComponent: function() {
      return this.tab;
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 10%;
}
</style>

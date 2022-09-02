<template>
  <b-container>
    <b-row align-v="center" class="gradient">
      <b-col class="p-0">
        <b-img alt="Mo Egger" src="@/assets/mo.jpeg" fluid></b-img>
      </b-col>
      <b-col cols="10">
        <h1 class="py-3 text-center font-bold">Let's Count Mo's Tweets!</h1>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
export default {
  name: 'Header',
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.gradient {
  background: linear-gradient(
    90deg,
    rgb(255, 255, 255) 20%,
    rgba(251, 79, 20, 1) 100%
  );
}
.font-bold {
  font-weight: bolder;
}
</style>

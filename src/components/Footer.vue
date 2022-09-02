<template>
  <div class="footer pt-2">
    <p>Copyright &copy; 2022</p>
    <p class="mx-2">|</p>
    <p>
      Built by
      <a target="_blank" href="https://www.linkedin.com/in/jgoney/"
        >Justin Goney</a
      >
    </p>
  </div>
</template>

<script>
export default {
  name: 'Footer',
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.footer {
  display: flex;
  align-items: center;
  justify-content: center;
}

a {
  color: rgba(255, 255, 255, 0.8);
}
</style>

<template>
  <div class="layout">
    <Header></Header>
    <App class="content"></App>
    <Footer></Footer>
  </div>
</template>

<script>
import App from '@/App';
import Header from '@/components/Header';
import Footer from '@/components/Footer';

export default {
  name: 'Layout',
  components: {
    App,
    Header,
    Footer,
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
.footer {
  color: rgba(255, 255, 255, 0.8);
  background-color: black;
}

.layout {
  display: flex;
  min-height: 100vh;
  flex-direction: column;
}

.content {
  flex: 1;
}
</style>

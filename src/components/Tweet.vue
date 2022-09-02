<template>
  <blockquote class="twitter-tweet">
    <a :href="tweetURL"></a>
  </blockquote>
</template>

<script>
import TwitterWidgetsLoader from 'twitter-widgets';

export default {
  name: 'Tweet',
  props: {
    tweetURL: {
      type: String,
    },
  },
  watch: {
    tweetURL() {
      this.updateTweetURL();
    },
  },
  methods: {
    updateTweetURL() {
      TwitterWidgetsLoader.load((err, twttr) => {
        if (err) {
          console.error('err', err);
          return;
        }
        twttr.widgets.load();
      });
    },
  },
  mounted() {
    this.updateTweetURL();
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>

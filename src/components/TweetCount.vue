<template>
  <h3>
    From <strong>{{ start }}</strong> to <strong>{{ end }}</strong
    >, Mo tweeted <span class="tweet-count">{{ times }}</span> time(s).
  </h3>
</template>

<script>
import TwitterWidgetsLoader from 'twitter-widgets';

export default {
  name: 'TweetCount',
  props: {
    start: {
      type: String,
    },
    end: {
      type: String,
    },
    times: {
      type: Number,
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
<style scoped>
.tweet-count {
  color: #c6011f;
  font-weight: bold;
}
</style>

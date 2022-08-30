<template>
  <b-container>
    <b-row>
      <h1 class="py-3">Let's Count Mo's Tweets</h1>
    </b-row>
    <b-row>
      <b-col>
        <label for="example-datepicker">Gather Tweets from...</label>
        <b-form-datepicker
          id="start-date"
          v-model="startDate"
          class="mb-2"
        ></b-form-datepicker>
        <b-form-timepicker v-model="startTime" locale="en"></b-form-timepicker>
      </b-col>
      <b-col>
        <label for="example-datepicker">...until</label>
        <b-form-datepicker
          id="end-date"
          v-model="endDate"
          class="mb-2"
        ></b-form-datepicker>
        <b-form-timepicker v-model="endTime" locale="en"></b-form-timepicker>
      </b-col>
    </b-row>
    <b-row>
      <b-col>
        <b-button
          @click="getTweets"
          :disabled="loading"
          size="lg"
          class="mt-3"
          :variant="loading ? 'secondary' : 'primary'"
          >{{ loading ? 'Loading...' : 'Fetch Tweet totals' }}</b-button
        >
        <div v-if="debug">
          <p v-for="tweet in tweets" :key="tweet.id">{{ tweet.text }}</p>
          <p>{{ this.tweetCount }}</p>
          <p>{{ this.tweets.length }}</p>
          <p>{{ this.tweets.length === this.tweetCount }}</p>
        </div>
      </b-col>
    </b-row>
    <b-row v-if="tweetCount" class="mt-3">
      <p size="lg">
        From <strong>{{ startDatetimeFriendly }}</strong> to
        <strong>{{ endDatetimeFriendly }}</strong
        >, Mo tweeted <span class="tweet-count">{{ tweetCount }}</span> time(s).
      </p>
    </b-row>
    <b-row>
      <b-col v-if="mostLiked">
        <a target="_blank" :href="mostLikedURL">Mo's most liked Tweet...</a>
        <blockquote class="twitter-tweet">
          <a :href="mostLikedURL"></a>
        </blockquote>
      </b-col>
      <b-col v-if="mostReplied">
        <a target="_blank" :href="mostRepliedURL"
          >Mo's most talked about Tweet...</a
        >
        <blockquote class="twitter-tweet">
          <a :href="mostRepliedURL"></a>
        </blockquote>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
import dayjs from 'dayjs';
import customParseFormat from 'dayjs/esm/plugin/customParseFormat';
import localizedFormat from 'dayjs/esm/plugin/localizedFormat';

dayjs.extend(customParseFormat);
dayjs.extend(localizedFormat);

import TwitterWidgetsLoader from 'twitter-widgets';

export default {
  name: 'App',
  data() {
    return {
      startDate: '',
      endDate: '',
      startTime: '',
      endTime: '',
      loading: false,
      tweetCount: 0,
      tweets: [],
      mostLiked: '',
      mostReplied: '',
      debug: false,
    };
  },
  computed: {
    startDatetimeISO8601() {
      if (this.startDate && this.startTime) {
        const concat = `${this.startDate} ${this.startTime}`;
        return dayjs(concat, 'YYYY-MM-DD HH:mm:ss').toISOString();
      }
      return '';
    },
    endDatetimeISO8601() {
      if (this.endDate && this.endTime) {
        const concat = `${this.endDate} ${this.endTime}`;
        return dayjs(concat, 'YYYY-MM-DD HH:mm:ss').toISOString();
      }
      return '';
    },
    startDatetimeFriendly() {
      return dayjs(this.startDatetimeISO8601).format('L LT');
    },
    endDatetimeFriendly() {
      return dayjs(this.endDatetimeISO8601).format('L LT');
    },
    mostRepliedURL() {
      return 'https://twitter.com/x/status/' + this.mostReplied;
    },
    mostLikedURL() {
      return 'https://twitter.com/x/status/' + this.mostLiked;
    },
  },
  methods: {
    async getTweets() {
      this.loading = true;

      const response = await fetch(
        '/api/tweets?' +
          new URLSearchParams({
            end_time: this.endDatetimeISO8601,
            start_time: this.startDatetimeISO8601,
          })
      );
      const json = await response.json();

      this.tweets = json.tweets;
      this.tweetCount = json.tweet_count;

      this.mostLiked = json.most_liked;
      this.mostReplied = json.most_replied;

      this.loading = false;
    },
  },
  mounted() {
    TwitterWidgetsLoader.load();

    const now = dayjs();
    const then = now.subtract(7, 'day');

    this.startDate = then.toISOString();
    this.endDate = now.toISOString();

    this.startTime = then.format('HH:mm:ss');
    this.endTime = now.format('HH:mm:ss');
  },
};
</script>

<style>
.tweet-count {
  color: red;
  font-weight: bold;
}
</style>

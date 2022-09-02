<template>
  <b-container>
    <b-row class="grey py-4 px-3">
      <b-row>
        <b-col>
          <label for="start-date" class="bold">Gather Tweets from...</label>
          <b-form-datepicker
            id="start-date"
            v-model="startDate"
            class="mb-2"
          ></b-form-datepicker>
          <b-form-timepicker
            v-model="startTime"
            locale="en"
          ></b-form-timepicker>
        </b-col>
        <b-col>
          <label for="end-date" class="bold">...until</label>
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
    </b-row>

    <b-row v-if="tweetCount" class="p-5 text-center">
      <TweetCount
        :start="startDatetimeFriendly"
        :end="endDatetimeFriendly"
        :times="tweetCount"
      ></TweetCount>
    </b-row>
    <b-row v-if="errorText" class="p-5">
      <p>{{ errorText }}</p>
    </b-row>
    <b-row class="grey py-4 px-3" align-h="between">
      <b-col v-if="mostLiked">
        <a target="_blank" :href="mostLikedURL">Mo's most liked Tweet...</a>
        <Tweet :tweetURL="mostLikedURL"></Tweet>
      </b-col>
      <b-col v-if="mostReplied">
        <a target="_blank" :href="mostRepliedURL"
          >Mo's most talked about Tweet...</a
        >
        <Tweet :tweetURL="mostRepliedURL"></Tweet>
      </b-col>
    </b-row>
    <b-row v-if="tweets.length" class="py-5">
      <b-col class="p-0">
        <h4>Mo's tweets in detail...</h4>
        <b-table striped hover :fields="tableFields" :items="tweetTable">
          <!-- A virtual column -->
          <template #cell(index)="data">
            {{ data.index + 1 }}
          </template>

          <template #cell(id)="data">
            <a
              target="_blank"
              :href="'https://twitter.com/x/status/' + data.item.id"
              >{{ data.item.id }}</a
            >
          </template>
        </b-table>
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

import Tweet from '@/components/Tweet';
import TweetCount from '@/components/TweetCount';

export default {
  name: 'App',
  components: {
    Tweet,
    TweetCount,
  },
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
      errorText: '',
      tableFields: [
        // A virtual column that doesn't exist in items
        'index',
        { key: 'id', sortable: true },
        { key: 'text', sortable: true },
        { key: 'retweet_count', sortable: true },
        { key: 'reply_count', sortable: true },
        { key: 'like_count', sortable: true },
        { key: 'quote_count', sortable: true },
      ],
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
    tweetTable() {
      return this.tweets.map((tweet) => {
        return {
          id: tweet.id,
          text: tweet.text,
          retweet_count: tweet.public_metrics.retweet_count,
          reply_count: tweet.public_metrics.reply_count,
          like_count: tweet.public_metrics.like_count,
          quote_count: tweet.public_metrics.quote_count,
        };
      });
    },
  },
  methods: {
    async getTweets() {
      this.loading = true;
      this.errorText = '';

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

      if (this.tweets.length === 0) {
        this.errorText = 'No Tweets retrieved, possible error. :-(';
      }

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

<style scoped>
a {
  color: black;
}

.grey {
  background-color: rgb(242, 242, 242);
}

.bold {
  font-weight: bolder;
}
</style>

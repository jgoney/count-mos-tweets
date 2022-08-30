# count-mos-tweets

Simple web app for counting Mo's tweets.

## Quickstart:

```
npm install
npm run build  # You have to build it the first time, as the server looks for dist/index.html, and will panic if not found
npm run serve
```

Run the server in a separate terminal:

```
go run app.go
```

To compile and minify frontend for production:

```
npm run build
```

To lint and fix files:

```
npm run lint
```

## Key points:

The port you run the Gin server on must match the port you define in `vue.config.js`

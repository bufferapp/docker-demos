const express = require('express');

const app = express();

// Logging middleware
app.use((req, res, next) => {
  console.log(`${req.method} ${req.url} ${res.statusCode}`);
  next();
});

// Routes
app.get('/', (req, res) => {
  res.send('OK!');
});

app.get('/user/:name', (req, res) => {
  res.send(`Hello ${req.params.name}`);
});

app.listen(process.env.PORT);

console.log(`Now listening on port ${process.env.PORT}`)

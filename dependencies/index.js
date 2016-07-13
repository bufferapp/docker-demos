const request = require('request');

request.get('https://api.bufferapp.com/', (err, response, body) => {
  if (!err && response.statusCode === 200) {
    console.log(body);
  }
  process.exit(0);
});

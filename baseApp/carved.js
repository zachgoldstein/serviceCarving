const protobuf = require("protobufjs");
const express = require('express');
const bodyParser = require('body-parser');
const request = require("request");

const app = express();
app.use(bodyParser.json());
app.use(express.json());

app.get('/', (req, res) => {
  res.send('Hello World!');
})

function getSumBuffer(data) {
  const errMsg = Sum.verify(data);
  if (errMsg) {
    throw errMsg;
  }
  const sum = Sum.fromObject(data);
  return Sum.encode(sum).finish();
}

function processCount(requestData, callback) {
  var times = Times.decode(requestData);
  request({
    method: 'POST',
    uri: 'http://localhost:6666/twirp/twirp.carved.Carved/Count',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(times)
  }, function(error, response, body) {
    const buffer = getSumBuffer(JSON.parse(body));
    callback(error, buffer);
  });
}

function rpcImpl(method, requestData, callback) {
  if (method.name === 'Count') {
    processCount(requestData, callback);
  }
}

app.post('/Count', (req, res) => {
  carvedService.count(req.body, function(err, response) {
    res.send(`Counted to ${response.Sum} \n`)
  });
})

const root = protobuf.loadSync("../carvedService.proto");
const CarvedService = root.lookupService("Carved");
const carvedService = CarvedService.create(rpcImpl, false, false);
const Times = root.lookupType("Times");
const Sum = root.lookupType("Sum");

app.listen(3000, () => {
  console.log('BaseApp app listening on port 3000!');
})

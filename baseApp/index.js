const express = require('express')
const bodyParser = require('body-parser')

const app = express()
app.use(bodyParser.json());
app.use(express.json());

app.get('/', (req, res) => {
  res.send('Hello World!')
})

app.post('/Count', (req, res) => {
  // Calculate the sum of lots of random numbers
  let sum = 0
  const timesToIterate = req.body.times
  for (var i = 1; i <= timesToIterate; i++) {
    sum += Math.round(Math.random() * 100)
  }
  res.send(`Counted to ${sum} \n`)
})

app.listen(3000, () => console.log('BaseApp app listening on port 3000!'))

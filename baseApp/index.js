const express = require('express')
const app = express()

app.get('/', (req, res) => {
  res.send('Hello World!')
})

app.get('/intensive', (req, res) => {
  // Calculate the sum of lots of random numbers
  let sum = 0
  const timesToIterate = 10000000000
  for (var i = 1; i <= timesToIterate; i++) {
   sum += Math.round(Math.random() * 100)
  }
  res.send(`Counted to ${sum} \n`)
})

app.listen(3000, () => console.log('BaseApp app listening on port 3000!'))

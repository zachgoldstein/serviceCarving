# Service Carving
<img align="right" width="200" src="https://i.imgur.com/XVG2Gd7.png">

This is demonstration repo to show how to carve away a high-performance golang service from a standard node app using protobuf and twirp, an RPC system.

Some major benefits of doing this:
  - It scales a system cleanly across many cores, more fully utilising larger machines.
  - The API boundary between the two systems isolates parts of the system and encourages better separation of concerns
  - Failures are more easy to isolate
  - You can upgrade parts of the system more easily
  - Multiple people or teams can work in parallel more effectively, as their roles/goals can be aligned around the boundaries you make between services
  - Yada Yada. Microservices can be great.

This project uses a very simple task as an example of a CPU intensive job - counting up random numbers.

This project has three main components:
  - A node server that counts up numbers. (./baseApp/index.js)
  - A golang server that counts up numbers (./carvedService/main.go)
  - A node server that uses the golang service to do the heavy lifting (./baseApp/carved.js)

There are also yml files that were used to do load testing and illustrate performance differences.

See the blog post for more details and results here: XXX.

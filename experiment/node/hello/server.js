var http = require("http");
var server = http.createServer(function(req, res) {
  console.log(req)
  res.writeHead(200, {"Content-Type": "text/plain"});
  res.write("Hello World!");
  res.end();
});

server.listen(9999, "0.0.0.0")
console.log("Server is running on 9999")

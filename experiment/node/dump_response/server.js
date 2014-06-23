var express = require("express");
var app = express();
var errorHandler = require("express-error-handler");
var bodyParser = require("body-parser");
var port = 13337;

app.use(errorHandler());
app.use(bodyParser());

////////////////////////////////////////////////

app.use(function(req, res, next) {
  console.log(res);
  res.send();
});

////////////////////////////////////////////////

app.listen(port);
console.log('Starting server on port 13337...');

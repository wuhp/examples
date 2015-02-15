var express = require("express");
var app = express();
var errorHandler = require("express-error-handler");
var bodyParser = require("body-parser");
var port = 13337;

app.use(errorHandler());
app.use(bodyParser());

///////////////////////////////////////////////

function midHandler1(req, res, next)  {
  console.log("Handler1");
  next();
}

function midHandler2(req, res, next)  {
  console.log("Handler2");
  next();
}

////////////////////////////////////////////////

app.use(function(req, res, next) {
  console.log("Midware 01");
  next();
});

app.use(function(req, res, next) {
  console.log("Midware 02");
  next();
});

/////////////////////////////////////////////////

app.all("/api", function(req, res, next) {
  console.log("Api all");
  next();
//  next("Customer error");
//  next(new Error("Customer error."));
});

app.get('/api', midHandler1, midHandler2, function(req, res) {
  console.log("Api get");
  res.send("End");
});

app.post('/api', function(req, res) {
  console.log("Api post");
  res.send("End");
});

app.put('/api', function(req, res) {
  console.log("Api put");
  res.send("End");
});

app.delete('/api', function(req, res) {
  console.log("Api delete");
  res.send("End");
});

////////////////////////////////////////////////

app.get("/test001/:id001", function(req, res, next) {
  console.log("Api get test001 " + req.params.id001);
  res.send("End");
});

////////////////////////////////////////////////

app.param("id002", function(req, res, next, id) {
  console.log("Param inventory " + id);
  req.xxx = id;
  next();
});

app.get("/test002/:id002", function(req, res) {
  console.log("Api get test002 " + req.xxx);
  res.send("End");
});

////////////////////////////////////////////////

var x = "default";
app.get("/data", function(req, res) {
  console.log(typeof(req.query));
  console.log("Query: %j", req.query);
  var data = {};
  data.value = x;
  res.json(data);
});

app.post("/data", function(req, res) {
  x = req.body.value;
  res.send("END");
});

////////////////////////////////////////////////

app.listen(port);
console.log('Starting server on port 13337...');

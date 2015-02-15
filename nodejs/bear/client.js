var http = require("http");

// This is a error script, because all requests are async
// The invoke seq is same as this script 
var ids = [];


list_fn = function(res) {
  console.log("Status %d", res.statusCode);
  var str = "";
  res.on("data", function(chunk) {
    str += chunk;
  });
  res.on("end", function() {
    var bear_set = JSON.parse(str);
    console.log("Getting %d object...", bear_set.length);
    for (var i = 0; i < bear_set.length; i++) {
      var bear = bear_set[i];
      console.log("Bear, id = %s, name = %s", bear._id, bear.name);
    }
  });
}

create_fn = function(res) {
  console.log("Status %d", res.statusCode);
  var str = "";
  res.on("data", function(chunk) {
    str += chunk;
  });
  res.on("end", function(chunk) {
    var object = JSON.parse(str);
    ids.push(object._id);
  });
}

query_fn = function(res) {
  console.log("Status %d", res.statusCode);
  var str = "";
  res.on("data", function(chunk) {
    str += chunk;
  });
  res.on("end", function(chunk) {
    var bear = JSON.parse(str);
    console.log("Bear, id = %s, name = %s", bear._id, bear.name);
  });
}

delete_fn = function(res) {
  console.log("Status %d", res.statusCode);
}

// create data
var options = {};
options.host = "localhost";
options.port = 8888;
options.path = "/api/bears";
options.method = "POST";
options.headers = {"Content-Type": "application/json"};

console.log("Create testing data1...");
var data = {};
data.name = "data1";

var req = http.request(options, create_fn);
req.write(JSON.stringify(data));
req.end();

console.log("Create testing data2...");
var data = {};
data.name = "data2";

var req = http.request(options, create_fn);
req.write(JSON.stringify(data));
req.end();

// List data
var options = {};
options.host = "localhost";
options.port = 8888;
options.path = "/api/bears";
options.method = "GET";

console.log("List all testing data...");
var req = http.request(options, list_fn);
req.end();

// Query data
var options = {};
options.host = "localhost";
options.port = 8888;
options.method = "GET";

for ( var i = 0; i < ids.length; i++ ) {
  options.path = "/api/bears/" + ids[i];
  console.log("Query testing data %d", ids[i]);
  var req = http.request(options, query_fn);
  req.end();
}

// Delete data
var options = {};
options.host = "localhost";
options.port = 8888;
options.method = "DELETE";

console.log("Delete all testing data...");
for ( var i = 0; i < ids.length; i++ ) {
  options.path = "/api/bears/" + ids[i];
  console.log("Delete testing data %d", ids[i]);
  var req = http.request(options, delete_fn);
  req.end();
}

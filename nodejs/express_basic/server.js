var http = require("http");
var express = require("express");
var mongoose = require("mongoose");
var bodyParser = require("body-parser");

var app = express();
app.set("port", process.env.PORT || 8888);


app.use(function(req, res, next) {
  console.log("%s %s", req.method, req.url);
  next();
});

app.use(function(req, res, next) {
  res.send("Hello World")
});

var router = express.Router();

router.get("/", function(req, res) {
  res.json({message: "Hooray! Welcome to our api!"});
});

router.route("/bears")
  .post(function(req, res){
    var bear = new Bear();
    bear.name = req.body.name;
    bear.save(function(err) {
      if (err) {
        res.send(err);
        return;
      }
      res.json({message: "Bear created!"});
    });
  })

  .get(function(req, res) {
    Bear.find(function(err, bears) {
      if (err) {
        res.send(err);
        return;
      }
      res.json(bears);
    });
  })
;

router.route("/bears/:bear_id")
  .get(function(req, res) {
    Bear.findById(req.params.bear_id, function(err, bear) {
      if (err) {
        res.send(err);
        return;
      }
      res.json(bear);
    });
  })

  .put(function(req, res) {
    Bear.findById(req.params.bear_id, function(err, bear) {
      if (err) {
        res.send(err);
        return;
      }
      bear.name = req.body.name;
      bear.save(function(err) {
        if (err) {
          res.send(err);
          return;
        }
        res.json({message: 'Bear updated!'});
      });
    });
  })
  .delete(function(req, res) {
    Bear.remove({_id: req.params.bear_id}, function(err, bear) {
      if (err) {
        res.send(err);
        return;
      } 
      res.json({message: "Bear deleted!"});
    });
  })
;

app.use(bodyParser());
app.use("/api", router);

http.createServer(app).listen(app.get('port'), function() {
  console.log("Express server listening on port " + app.get("port"));
});

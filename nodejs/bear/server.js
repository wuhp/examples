var express = require("express");
var bodyParser = require("body-parser");
var mongoose = require("mongoose");
var Bear = require("./app/models/bear");

var app = express();
var port = process.env.PORT || 8888;

mongoose.connect("mongodb://:@localhost:27017/example");

var router = express.Router();

// middleware to use for all requests
router.use(function(req, res, next) {
  console.log(req.method, req.url);
  next();
});

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
      res.json({_id: bear._id});
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
        res.send();
      });
    });
  })
  .delete(function(req, res) {
    Bear.remove({_id: req.params.bear_id}, function(err, bear) {
      if (err) {
        res.send(err);
        return;
      } 
      res.send();
    });
  })
;

app.use(bodyParser());
app.use("/api", router);
app.listen(port);
console.log("Server running on port " + port);

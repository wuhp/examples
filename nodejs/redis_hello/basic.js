var redis = require("redis");
var client = redis.createClient(16379, "localhost");

client.set("str001", "001", function(err, reply) {
  if (err) {
    console.log(err);
    return;
  }
  client.get("str001", function(err, reply) {
    if (err) {
      console.log(err);
      return;
    }
    console.log(reply);
    client.quit();
  });
});

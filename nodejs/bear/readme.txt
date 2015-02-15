### start
node server.js

### client
node client.js

### client from cmd
curl http://localhost:8888/api/bears
curl http://localhost:8888/api/bears -X POST -d '{"name":"b1"}' -H "Content-Type:application/json"
curl http://localhost:8888/api/bears -X POST -d '{"name":"b2"}' -H "Content-Type:application/json"
curl http://localhost:8888/api/bears/xxxxxxxxxxxx -X PUT -d '{"name":"b11"}' -H "Content-Type:application/json"
curl http://localhost:8888/api/bears/xxxxxxxxxxxx -X DELETE

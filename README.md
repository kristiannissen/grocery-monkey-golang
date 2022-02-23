# Grocery Monkey Go

```
docker run -it -p 3000:3000 -v $(pwd)/:/go/src [image id]
```

To run the server

```
root@6003e4f6a41c:/go/src# PORT=3000 go run main.go
```

## To test the jwt token

```
wget --post-data 'username=kitty' http://localhost:PORT/sign -O-
```

### Test create list

```
wget -O- --post-data='{"..."}' --header='Content-Type:application/json' --header='Authorization: Bearer ...' http://localhost:PORT/groceries
```

### Update list

```
wget --method=PUT --body-data='{"user":"612f5471-6fa1-491f-8513-1695853ce6f9","subscribers":["1","2", "3"],"id":"b3ed8650-156a-4ad1-a144-1184d21e508e","groceries":[{"name":"Milk","qty":42,"unit":"Liters","store":"Wallmart"},{"name":"Beer","qty":100,"unit":"Bottles","store":"WholeFood"}]}' http://localhost:3001/groceries/update -O- --header='Content-Type:application/json' --header='Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImtpdHR5IiwidWlkIjoiNjEyZjU0NzEtNmZhMS00OTFmLTg1MTMtMTY5NTg1M2NlNmY5IiwiZXhwIjoxNjQ1OTA3NTU1fQ.nehX3te6X5xiSofiZSZx2D9Gm3MdnYZzc7qgesI7u7c'
```

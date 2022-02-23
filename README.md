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

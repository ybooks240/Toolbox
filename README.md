# Toolbox
Toolbox 


# 单节点

```
go run main.go redis --address localhost:6379 -m standalone set username james

go run main.go redis --address localhost:6379 -m standalone get username james
```
# 哨兵

```
go run main.go redis --address 172.16.123.137:26379 --address 172.16.123.138:26379 --address 172.16.123.139:26379 -m sentinel -u redis_cncp -p wXGwskVXi2vCBSld set username james

go run main.go redis --address 172.16.123.137:26379 --address 172.16.123.138:26379 --address 172.16.123.139:26379 -m sentinel -u redis_cncp -p wXGwskVXi2vCBSld get username james
```

# 集群模式
```
go run main.go redis --address  172.16.123.131:50101 --address 172.16.123.132:50101 --address 172.16.123.133:50101 --address 172.16.123.134:50101 --address 172.16.123.135:50101 --address 172.16.123.136:50101 -m cluster -u redis_cncp -p RMKIOPZAdF9e2s7G set  username james

go run main.go redis --address  172.16.123.131:50101 --address 172.16.123.132:50101 --address 172.16.123.133:50101 --address 172.16.123.134:50101 --address 172.16.123.135:50101 --address 172.16.123.136:50101 -m cluster -u redis_cncp -p RMKIOPZAdF9e2s7G get  username 
```



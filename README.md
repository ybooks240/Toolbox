# Toolbox
Toolbox 


# 单节点

```
go run main.go redis --address localhost:6379 -m standalone set username james
go run main.go redis --address localhost:6379 -m standalone get username james
```
# 哨兵

```
go run main.go redis --address 172.16.123.139:26379 -m sentinel -u redis_cncp -p wXGwskVXi2vCBSld set username james
go run main.go redis --address 172.16.123.139:26379 -m sentinel -u redis_cncp -p wXGwskVXi2vCBSld get username james
```
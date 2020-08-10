##Etcd 的镜像拉取
`$ docker pull quay.io/coreos/etcd`

##编辑etcd-compose.yml,推荐养成用compose管理container的习惯

`$ cat etcd-compose.yml`
#
`
version: '3'
services:
  etcd-node1:
    image: "quay.io/coreos/etcd"
    container_name: "etcd-node1"
    ports:
      - "12379:2379"
      - "12380:2380"
    command: 'etcd -name etcd-node1 -advertise-client-urls http://0.0.0.0:2379 -listen-client-urls http://0.0.0.0:2379 -listen-peer-urls http://0.0.0.0:2380 -initial-cluster-token etcd-cluster -initial-cluster "etcd-node1=http://etcd-node1:2380,etcd-node2=http://etcd-node2:2380,etcd-node3=http://etcd-node3:2380" -initial-cluster-state new'
    networks:
      - "etcd"
  etcd-node2:
    image: "quay.io/coreos/etcd"
    container_name: "etcd-node2"
    ports:
      - "22379:2379"
      - "22380:2380"
    command: 'etcd -name etcd-node2 -advertise-client-urls http://0.0.0.0:2379 -listen-client-urls http://0.0.0.0:2379 -listen-peer-urls http://0.0.0.0:2380 -initial-cluster-token etcd-cluster -initial-cluster "etcd-node1=http://etcd-node1:2380,etcd-node2=http://etcd-node2:2380,etcd-node3=http://etcd-node3:2380" -initial-cluster-state new'
    networks:
      - "etcd"
  etcd-node3:
    image: "quay.io/coreos/etcd"
    container_name: "etcd-node3"
    ports:
      - "32379:2379"
      - "32380:2380"
    command: 'etcd -name etcd-node3 -advertise-client-urls http://0.0.0.0:2379 -listen-client-urls http://0.0.0.0:2379 -listen-peer-urls http://0.0.0.0:2380 -initial-cluster-token etcd-cluster -initial-cluster "etcd-node1=http://etcd-node1:2380,etcd-node2=http://etcd-node2:2380,etcd-node3=http://etcd-node3:2380" -initial-cluster-state new'
    networks:
      - "etcd"
networks:
  etcd:
`

##3.启动docker-compose
`$ docker-compose -f etcd-compose.yml up -d`
##4.检查容器状态查询节点成员
`
$ docker-compose -f etcd-compose.yml ps                                                                                                                                        
    Name                 Command               State                        Ports
 ------------------------------------------------------------------------------------------------------
 etcd-node1   etcd -name etcd-node1 -adv ...   Up      0.0.0.0:12379->2379/tcp, 0.0.0.0:12380->2380/tcp
 etcd-node2   etcd -name etcd-node2 -adv ...   Up      0.0.0.0:22379->2379/tcp, 0.0.0.0:22380->2380/tcp
 etcd-node3   etcd -name etcd-node3 -adv ...   Up      0.0.0.0:32379->2379/tcp, 0.0.0.0:32380->2380/tcp
`
##### 查询每台的成员，结果是一致的
##### curl http://127.0.0.1:12379/v2/members | json_pp
##### curl http://127.0.0.1:22379/v2/members | json_pp
##### curl http://127.0.0.1:32379/v2/members | json_pp

`
{
 "members" : [
    {
       "id" : "5b926f852fa1811",
       "peerURLs" : [
          "http://etcd-node1:2380"
       ],
       "clientURLs" : [
          "http://0.0.0.0:2379"
       ],
       "name" : "etcd-node1"
    },
    {
       "peerURLs" : [
          "http://etcd-node2:2380"
       ],
       "clientURLs" : [
          "http://0.0.0.0:2379"
       ],
       "name" : "etcd-node2",
       "id" : "9b3cd975d37c44ce"
    },
    {
       "id" : "9e13ad3ed0f8a26b",
       "peerURLs" : [
          "http://etcd-node3:2380"
       ],
       "clientURLs" : [
          "http://0.0.0.0:2379"
       ],
       "name" : "etcd-node3"
    }
 ]
}
`

##5.销毁测试环境
`
$ docker-compose -f etcd-compose.yml down
 Stopping etcd-node2 ... done
 Stopping etcd-node1 ... done
 Stopping etcd-node3 ... done
 Removing etcd-node2 ... done
 Removing etcd-node1 ... done
 Removing etcd-node3 ... done
 Removing network docker-compose_etcd
 `
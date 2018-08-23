# 说明
两个p2p节点：node1和node2

node2主动添加node1，node2向node1发送whisper消息，node1消费消息，其实就类似消息队列，一个生产者，一个消费者，但是对比kafka这种消息队列在数据可靠性上还是弱很多的


whisper的概念是一端创建消息(msg)，封到信封里(Envelope)，发送出去；另一端收到信封，拆开信封，读取消息；但是这之间信封是加密过的

# 执行
- cd node1
- go run main.go
- cd node2
- go run main.go

## node1输出
```bash
INFO [08-23|18:15:40.958] Starting P2P networking
INFO [08-23|18:15:40.961] UDP listener up                          self=enode://fdf241bcdd60ad1ab2e53518b43bf585d14f016c273595b65e4daea86151e2f839fa7a99ca1e188ae8d2263c42227b51512a8687da2f686ebf59559500e9dde3@[::]:8000
INFO [08-23|18:15:40.961] Node                                     info="&{ID:fdf241bcdd60ad1ab2e53518b43bf585d14f016c273595b65e4daea86151e2f839fa7a99ca1e188ae8d2263c42227b51512a8687da2f686ebf59559500e9dde3 Name: Enode:enode://fdf241bcdd60ad1ab2e53518b43bf585d14f016c273595b65e4daea86151e2f839fa7a99ca1e188ae8d2263c42227b51512a8687da2f686ebf59559500e9dde3@[::]:8000 IP::: Ports:{Discovery:8000 Listener:8000} ListenAddr:[::]:8000 Protocols:map[shh:map[version:6.0 maxMessageSize:1048576 minimumPoW:0]]}"
INFO [08-23|18:15:40.961] RLPx listener up                         self=enode://fdf241bcdd60ad1ab2e53518b43bf585d14f016c273595b65e4daea86151e2f839fa7a99ca1e188ae8d2263c42227b51512a8687da2f686ebf59559500e9dde3@[::]:8000
TRACE[08-23|18:15:40.961] New dial task                            task="discovery lookup"
INFO [08-23|18:15:40.961] started whisper v.6.0
TRACE[08-23|18:15:40.962] Dial task done                           task="discovery lookup"
TRACE[08-23|18:15:40.962] New dial task                            task="discovery lookup"
INFO [08-23|18:15:41.966] filter.Retrieve
INFO [08-23|18:15:42.969] filter.Retrieve
INFO [08-23|18:15:43.970] filter.Retrieve
TRACE[08-23|18:15:44.782] Accepted connection                      addr=127.0.0.1:64111
DEBUG[08-23|18:15:44.785] Adding p2p peer                          name= addr=127.0.0.1:64111 peers=1
TRACE[08-23|18:15:44.785] connection set up                        id=50095ad7bd27b0c9 addr=127.0.0.1:64111 conn=inbound inbound=true
TRACE[08-23|18:15:44.785] Starting protocol shh/6                  id=50095ad7bd27b0c9 conn=inbound
TRACE[08-23|18:15:44.786] start                                    peer="[80 9 90 215 189 39 176 201 158 103 58 144 177 248 24 180 8 222 213 103 42 197 120 207 128 121 147 51 190 55 28 25 12 2 161 193 107 146 189 208 137 76 87 87 197 196 152 122 254 148 34 219 178 42 248 176 11 65 148 61 176 102 173 208]"
TRACE[08-23|18:15:44.968] Dial task done                           task="discovery lookup"
TRACE[08-23|18:15:44.968] New dial task                            task="discovery lookup"
INFO [08-23|18:15:44.972] filter.Retrieve
INFO [08-23|18:15:44.972] print peer info                          id=50095ad7bd27b0c9 name="Peer 50095ad7bd27b0c9 127.0.0.1:64111"
TRACE[08-23|18:15:45.785] processing message: decrypted            hash=0x7b89de81a30489fc02ae9ec713d9e1c449761d2cb4d09b9e87c4598f9cc47527
INFO [08-23|18:15:45.972] filter.Retrieve
INFO [08-23|18:15:45.972] recvd msg                                payload="This is whisper test message from node2"
TRACE[08-23|18:15:46.790] processing message: decrypted            hash=0x4b24018a0e1fbb6b0e4992bb0f3d36ec9b66496597a719e0b3bf6535b5b66fe7
INFO [08-23|18:15:46.974] filter.Retrieve
INFO [08-23|18:15:46.974] recvd msg                                payload="This is whisper test message from node2"
TRACE[08-23|18:15:47.790] processing message: decrypted            hash=0xa355fd78e985d44508f1fcd7dce7721a84b152d01ded735bcea87c334a9ddc47
INFO [08-23|18:15:47.978] filter.Retrieve
INFO [08-23|18:15:47.978] recvd msg                                payload="This is whisper test message from node2"
TRACE[08-23|18:15:48.791] processing message: decrypted            hash=0xbdac588a1b25a7360c5cc8143d764d09494dc59da910f49822dc7e28161833c2
TRACE[08-23|18:15:48.968] Dial task done                           task="discovery lookup"
TRACE[08-23|18:15:48.968] New dial task                            task="discovery lookup"
INFO [08-23|18:15:48.981] filter.Retrieve
INFO [08-23|18:15:48.982] recvd msg                                payload="This is whisper test message from node2"
TRACE[08-23|18:15:49.794] processing message: decrypted            hash=0xb227a7ac8ba57504cafcd5b027c61b6fd49aebfe17282b8b06ba139929a4791b
INFO [08-23|18:15:49.986] filter.Retrieve
INFO [08-23|18:15:49.986] recvd msg                                payload="This is whisper test message from node2"
TRACE[08-23|18:15:50.800] processing message: decrypted            hash=0x9c6eac7a189368a7bd47028c5ef6508be3f35e4a0c696258256e09c01d431378
INFO [08-23|18:15:50.991] filter.Retrieve
INFO [08-23|18:15:50.991] recvd msg                                payload="This is whisper test message from node2"
TRACE[08-23|18:15:51.801] processing message: decrypted            hash=0x353aee2c36977b88f4e8e5fbb9a5be1ad49b31428fea43c79b1bca2554ae249c
INFO [08-23|18:15:51.995] filter.Retrieve
INFO [08-23|18:15:51.995] recvd msg                                payload="This is whisper test message from node2"
TRACE[08-23|18:15:52.805] processing message: decrypted            hash=0x0cccb9b64c8cd36cfdcde694478a59a939cb3a44dd59e87ae0ab5344e6051412
TRACE[08-23|18:15:52.969] Dial task done                           task="discovery lookup"
TRACE[08-23|18:15:52.969] New dial task                            task="discovery lookup"
INFO [08-23|18:15:52.999] filter.Retrieve
INFO [08-23|18:15:52.999] recvd msg                                payload="This is whisper test message from node2"
```


## node2输出
```bash
INFO [08-23|18:19:22.341] Starting P2P networking
INFO [08-23|18:19:22.344] UDP listener up                          self=enode://50095ad7bd27b0c99e673a90b1f818b408ded5672ac578cf80799333be371c190c02a1c16b92bdd0894c5757c5c4987afe9422dbb22af8b00b41943db066add0@[::]:8001
INFO [08-23|18:19:22.344] Node                                     info="&{ID:50095ad7bd27b0c99e673a90b1f818b408ded5672ac578cf80799333be371c190c02a1c16b92bdd0894c5757c5c4987afe9422dbb22af8b00b41943db066add0 Name: Enode:enode://50095ad7bd27b0c99e673a90b1f818b408ded5672ac578cf80799333be371c190c02a1c16b92bdd0894c5757c5c4987afe9422dbb22af8b00b41943db066add0@[::]:8001 IP::: Ports:{Discovery:8001 Listener:8001} ListenAddr:[::]:8001 Protocols:map[shh:map[version:6.0 maxMessageSize:1048576 minimumPoW:0.2]]}"
INFO [08-23|18:19:22.344] started whisper v.6.0
INFO [08-23|18:19:22.344] RLPx listener up                         self=enode://50095ad7bd27b0c99e673a90b1f818b408ded5672ac578cf80799333be371c190c02a1c16b92bdd0894c5757c5c4987afe9422dbb22af8b00b41943db066add0@[::]:8001
INFO [08-23|18:19:23.345] print peer info                          id=fdf241bcdd60ad1a name="Peer fdf241bcdd60ad1a 127.0.0.1:8000"
INFO [08-23|18:19:23.345] sent message done                        to="Peer fdf241bcdd60ad1a 127.0.0.1:8000"
INFO [08-23|18:19:24.345] print peer info                          id=fdf241bcdd60ad1a name="Peer fdf241bcdd60ad1a 127.0.0.1:8000"
INFO [08-23|18:19:24.346] sent message done                        to="Peer fdf241bcdd60ad1a 127.0.0.1:8000"
INFO [08-23|18:19:25.346] print peer info                          id=fdf241bcdd60ad1a name="Peer fdf241bcdd60ad1a 127.0.0.1:8000"
INFO [08-23|18:19:25.346] sent message done                        to="Peer fdf241bcdd60ad1a 127.0.0.1:8000"
```

### 目录结构 1 
```
├── Bug
│   ├── README.md
│   ├── d
│   │   └── main.go
│   └── pustToSessionbug
│       ├── README.md
│       └── copy_test.go
├── DiveIntoGo
│   ├── Concurrency
│   │   ├── README.md
│   │   ├── advance-concurrency
│   │   │   ├── README.md
│   │   │   ├── fetch.go
│   │   │   ├── rss_test.go
│   │   │   └── sub.go
│   │   ├── channel-antomay
│   │   │   ├── README.md
│   │   │   ├── leaky_test.go
│   │   │   ├── non_blocking_test.go
│   │   │   ├── ping-pong_test.go
│   │   │   ├── produce_channel_test.go
│   │   │   ├── sender_test.go
│   │   │   ├── test_channel_test.go
│   │   │   └── test_close_test.go
│   │   ├── communication
│   │   │   ├── README.md
│   │   │   └── main.go
│   │   ├── concurrency-in-go
│   │   │   ├── README.md
│   │   │   ├── chapter2
│   │   │   │   ├── dead-lock
│   │   │   │   │   └── dead_test.go
│   │   │   │   ├── live-lock
│   │   │   │   │   └── live_lock_test.go
│   │   │   │   └── starvation
│   │   │   │       └── stavation_test.go
│   │   │   ├── chapter3
│   │   │   │   ├── avalible-same
│   │   │   │   │   └── space_test.go
│   │   │   │   ├── channel
│   │   │   │   │   └── channel_test.go
│   │   │   │   ├── cond
│   │   │   │   │   └── cond_test.go
│   │   │   │   ├── context-switch
│   │   │   │   │   └── ben_test.go
│   │   │   │   ├── join-point
│   │   │   │   │   └── join_point_test.go
│   │   │   │   ├── memory_consumer
│   │   │   │   │   └── mem_consumer_test.go
│   │   │   │   ├── mutex
│   │   │   │   │   └── mutext_test.go
│   │   │   │   ├── once
│   │   │   │   │   └── once_test.go
│   │   │   │   ├── pool_test.go
│   │   │   │   ├── select\ 
│   │   │   │   │   └── select_test.go
│   │   │   │   └── test-wait-group
│   │   │   │       └── test-group_test.go
│   │   │   ├── chapter4
│   │   │   │   ├── confinement
│   │   │   │   │   └── confiement_test.go
│   │   │   │   ├── errorHandler
│   │   │   │   │   └── error_test.go
│   │   │   │   ├── generator
│   │   │   │   │   └── generator_test.go
│   │   │   │   ├── pipeline
│   │   │   │   │   └── pipeline_test.go
│   │   │   │   ├── preventLeak
│   │   │   │   │   └── prevent_test.go
│   │   │   │   └── repeat
│   │   │   │       └── repeat_test.go
│   │   │   ├── chapter5
│   │   │   │   ├── bridge
│   │   │   │   │   └── bridge_test.go
│   │   │   │   ├── doWork
│   │   │   │   │   └── dowork_test.go
│   │   │   │   ├── errorHandler
│   │   │   │   │   └── error_test.go
│   │   │   │   ├── fan
│   │   │   │   │   └── fan_test.go
│   │   │   │   ├── heartbeat
│   │   │   │   │   └── heartbeat_test.go
│   │   │   │   ├── orDone
│   │   │   │   │   └── or_done_test.go
│   │   │   │   ├── ratelimit
│   │   │   │   │   └── main.go
│   │   │   │   ├── ratelimit1
│   │   │   │   │   └── ratelimit1_test.go
│   │   │   │   ├── replicate
│   │   │   │   │   └── replicate_test.go
│   │   │   │   └── tee
│   │   │   │       └── tee_test.go
│   │   │   └── chapter6
│   │   │       └── chapter6_test.go
│   │   ├── errorgroup
│   │   └── forkjoinmodel
│   │       └── README.md
│   ├── CreckingSourceCode
│   │   ├── README.md
│   │   ├── compiler
│   │   │   └── doc.md
│   │   └── runtime
│   │       └── hacking.md
│   ├── GoFoundationBook
│   │   ├── README.md
│   │   ├── chapter2-float
│   │   │   ├── README.md
│   │   │   ├── add_test.go
│   │   │   └── main.go
│   │   ├── chapter3-type-infer
│   │   │   └── README.md
│   │   ├── chapter5-string
│   │   │   ├── README.md
│   │   │   └── main.go
│   │   ├── chpater1-compiler
│   │   │   ├── README.md
│   │   │   ├── escape_test.go
│   │   │   ├── inline_test.go
│   │   │   ├── main.go
│   │   │   ├── syntax_test.go
│   │   │   ├── token_test.go
│   │   │   └── variable_capture_test.go
│   │   ├── doc
│   │   │   ├── golang-devise-and-implmentation.md
│   │   │   └── note.md
│   │   ├── map
│   │   │   └── README.md
│   │   └── slice
│   │       └── README.md
│   ├── MockStandarLib
│   │   ├── function
│   │   │   ├── README.md
│   │   │   └── main.go
│   │   ├── go_start
│   │   │   ├── README.md
│   │   │   └── how_does_go
│   │   │       └── main.go
│   │   ├── http
│   │   │   ├── README.md
│   │   │   ├── STATE.md
│   │   │   └── server.go
│   │   ├── interface
│   │   │   ├── Doc.md
│   │   │   └── README.md
│   │   ├── msyql
│   │   │   ├── NOTE.md
│   │   │   ├── README.md
│   │   │   ├── conn_model.go
│   │   │   ├── db.go
│   │   │   ├── driver
│   │   │   │   ├── conn_interface.go
│   │   │   │   └── driver_model.go
│   │   │   ├── driver.go
│   │   │   ├── example
│   │   │   │   └── example_test.go
│   │   │   ├── pitfall.md
│   │   │   └── sql.go
│   │   ├── reflect
│   │   ├── scheduler
│   │   │   └── README.md
│   │   ├── suite
│   │   │   └── README.md
│   │   ├── sync
│   │   │   ├── README.md
│   │   │   ├── doc.md
│   │   │   ├── sync_test.go
│   │   │   └── syncpool.png
│   │   ├── time
│   │   │   ├── README.md
│   │   │   └── time.go
│   │   └── unsafe
│   │       ├── README.md
│   │       ├── convert_test.go
│   │       ├── doc.md
│   │       ├── main.go
│   │       ├── pointer_test.go
│   │       └── unsafe_test.go
│   ├── README.md
│   ├── datastructure
│   │   ├── channel.go
│   │   ├── goroutine.go
│   │   └── map.go
│   ├── feature
│   │   ├── BEC
│   │   │   ├── README.md
│   │   │   ├── example1
│   │   │   │   ├── main
│   │   │   │   └── main.go
│   │   │   ├── example2
│   │   │   │   ├── example2
│   │   │   │   └── example2.go
│   │   │   ├── example3
│   │   │   │   ├── example3
│   │   │   │   └── example3.go
│   │   │   ├── example4
│   │   │   │   └── example4.go
│   │   │   └── example5
│   │   │       └── example5.go
│   │   ├── README.md
│   │   ├── idflags
│   │   │   └── README.md
│   │   ├── memory
│   │   │   ├── NOTE.md
│   │   │   ├── README.md
│   │   │   ├── different-struct-memory
│   │   │   │   └── example_test.go
│   │   │   └── �\206\205�\230对�\220�\216\237�\220\206.md
│   │   └── profiling
│   │       └── README.md
│   ├── go-class
│   │   ├── Golang-Training
│   │   │   ├── README.md
│   │   │   └── REFERENCE.md
│   │   ├── chapter3-concurrency
│   │   │   ├── README.md
│   │   │   ├── api_desigin.go
│   │   │   ├── http_handler_test.go
│   │   │   └── track.go
│   │   └── chapter4-Engineering
│   │       └── README.md
│   ├── question
│   │   ├── QandA.md
│   │   ├── README.md
│   │   ├── host_connecting_test.go
│   │   ├── http_test.go
│   │   ├── select_test.go
│   │   ├── slice_test.go
│   │   └── unsafepoint_test.go
│   ├── the-go-programming-language
│   │   ├── chapter7_test.go
│   │   ├── chapter8_test.go
│   │   ├── chatper6_test.go
│   │   └── chatpter5_test.go
│   └── whathappen
│       ├── README.md
│       ├── defer_test.go
│       ├── function_test.go
│       ├── interface_test.go
│       ├── main.go
│       ├── main_test.go
│       ├── map_test.go
│       ├── slice_test.go
│       ├── string_test.go
│       ├── syntax_test.go
│       └── type_test.go
├── FrameWork
│   ├── BenchMark
│   │   ├── README.md
│   │   └── Tutorial.md
│   ├── archieve
│   │   ├── ES
│   │   │   ├── README.md
│   │   │   └── example
│   │   │       └── example_teat.go
│   │   ├── EmbeddedResources
│   │   │   ├── README.md
│   │   │   ├── gen.go
│   │   │   └── resources
│   │   ├── TextSimilarity
│   │   │   └── consine
│   │   │       ├── README.md
│   │   │       ├── main.go
│   │   │       └── sin_test.go
│   │   ├── badger
│   │   │   └── README.md
│   │   ├── binddata
│   │   │   └── GoGenerate
│   │   │       ├── README.md
│   │   │       ├── gopher.go
│   │   │       ├── gopher.y
│   │   │       └── y.output
│   │   ├── casbin
│   │   │   └── README.md
│   │   ├── go-query
│   │   │   ├── README.md
│   │   │   ├── chromp
│   │   │   │   └── main.go
│   │   │   └── spider_test.go
│   │   ├── logging
│   │   │   └── README.md
│   │   ├── mail
│   │   │   ├── README.md
│   │   │   ├── gmail.go
│   │   │   └── main.go
│   │   ├── neo4j
│   │   │   └── quick-start.md
│   │   └── tokenizer-nlp
│   │       ├── README.md
│   │       ├── jieba
│   │       │   ├── jieba_test.go
│   │       │   ├── stop_word.dict.utf8
│   │       │   └── user.dict.utf8
│   │       ├── sego
│   │       │   └── sego_test.go
│   │       └── vocabulary
│   │           └── voc_test.go
│   ├── ci-lint
│   │   └── README.md
│   ├── cron
│   │   └── README.md
│   ├── etcd
│   │   ├── NOTE.md
│   │   ├── README.md
│   │   ├── Raft
│   │   │   ├── Notes.md
│   │   │   ├── README.md
│   │   │   ├── doc.go
│   │   │   ├── interface.go
│   │   │   ├── model.go
│   │   │   ├── raft.go
│   │   │   ├── state_transfer.go
│   │   │   ├── step.go
│   │   │   └── unstable.go
│   │   ├── errors.md
│   │   └── etcd-guide.md
│   ├── jaeger
│   │   ├── README.md
│   │   ├── basic_struct
│   │   │   └── client
│   │   │       ├── allocator.go
│   │   │       ├── basic.go
│   │   │       ├── header_config.go
│   │   │       ├── metric.go
│   │   │       ├── oberser.go
│   │   │       ├── reporter.go
│   │   │       ├── sampler.go
│   │   │       ├── span.go
│   │   │       ├── tracer.go
│   │   │       └── transport.go
│   │   └── doc.md
│   ├── juno
│   │   ├── README.md
│   │   └── jupiter.md
│   ├── jupiter
│   │   ├── README.md
│   │   └── hello_world.md
│   ├── mirco
│   │   ├── README.md
│   │   └── go-chasis
│   │       ├── README.md
│   │       ├── doc.md
│   │       └── talk.md
│   ├── promethus
│   │   ├── README.md
│   │   ├── go-kit
│   │   │   ├── README.md
│   │   │   ├── factory.go
│   │   │   ├── promethues.go
│   │   │   ├── promethues_test.go
│   │   │   └── util.go
│   │   ├── jaeger-lib
│   │   │   └── README.md
│   │   ├── stat
│   │   │   ├── README.md
│   │   │   ├── prom.go
│   │   │   └── prom_test.go
│   │   └── zero
│   │       ├── README.md
│   │       ├── counter.go
│   │       └── mdoel.go
│   ├── proto
│   │   └── �\227段�\233��\226��\232\204�\235\221.md
│   ├── registry
│   │   └── README.md
│   ├── rpcx
│   │   └── README.md
│   ├── sqlx
│   │   └── README.md
│   └── testfiy
│       ├── README.md
│       └── suite
│           └── suite_test.go
├── Log.md
├── Middleware
│   ├── WebTech
│   │   ├── HTTP
│   │   │   ├── RateLimit
│   │   │   │   ├── README.md
│   │   │   │   ├── REAME.md
│   │   │   │   ├── allratelimiter.go
│   │   │   │   ├── leakybuckUber
│   │   │   │   │   └── README.md
│   │   │   │   ├── leakybucket
│   │   │   │   │   ├── README.md
│   │   │   │   │   ├── leakybucket.go
│   │   │   │   │   └── leakybucket_test.go
│   │   │   │   ├── throttler
│   │   │   │   │   ├── README.md
│   │   │   │   │   ├── remote
│   │   │   │   │   │   ├── options.go
│   │   │   │   │   │   ├── options_test.go
│   │   │   │   │   │   ├── throttler.go
│   │   │   │   │   │   └── throttler_test.go
│   │   │   │   │   └── throttler.go
│   │   │   │   └── tokenbucket
│   │   │   │       ├── README.md
│   │   │   │       └── tokentbucket.go
│   │   │   ├── http2
│   │   │   │   └── README.md
│   │   │   ├── httpclient
│   │   │   │   ├── README.md
│   │   │   │   └── httpclient.go
│   │   │   ├── httpclientMiddelWare
│   │   │   └── loghttpRequest
│   │   │       └── main.go
│   │   ├── LoadBalanceUsage
│   │   │   ├── README.md
│   │   │   ├── loadbalance
│   │   │   │   ├── README.md
│   │   │   │   └── main.go
│   │   │   └── loadbalanceGoIm
│   │   │       ├── main.go
│   │   │       └── main_test.go
│   │   └── decorate
│   │       └── README.md
│   ├── cache-genernal
│   │   ├── Cache-Reference-repo.md
│   │   ├── LRU-Cache
│   │   │   └── README.md
│   │   ├── NOTE.md
│   │   ├── README.md
│   │   ├── asyncache
│   │   │   └── README.md
│   │   ├── big-cache
│   │   │   ├── README.md
│   │   │   ├── bytequeue.go
│   │   │   ├── interface.go
│   │   │   └── model.go
│   │   ├── cache
│   │   │   ├── README.md
│   │   │   ├── cache.go
│   │   │   ├── const.go
│   │   │   ├── lru.go
│   │   │   ├── safemap.go
│   │   │   ├── stat.go
│   │   │   ├── take.go
│   │   │   ├── timingwheel.go
│   │   │   └── unstable.go
│   │   ├── go-cache
│   │   │   ├── READMD.md
│   │   │   └── item.go
│   │   ├── groupcache
│   │   │   ├── README.md
│   │   │   ├── groupcache.go
│   │   │   ├── lru.go
│   │   │   └── singleflight.go
│   │   ├── hash
│   │   │   ├── README.md
│   │   │   └── go-zero-hash
│   │   │       └── hash.go
│   │   ├── in-memroy-cache
│   │   │   └── memory_test.go
│   │   └── storeCache
│   │       └── interface.go
│   ├── client-go
│   │   └── README.md
│   ├── common
│   │   ├── README.md
│   │   └── db
│   │       └── README.md
│   ├── config-center
│   │   └── README.md
│   ├── cronjob
│   │   └── README.md
│   ├── discovery
│   │   ├── README.md
│   │   ├── app.go
│   │   ├── apps.go
│   │   ├── client
│   │   │   ├── builder.go
│   │   │   ├── chore.go
│   │   │   ├── const.go
│   │   │   ├── discovery.go
│   │   │   ├── instance.go
│   │   │   ├── interface.go
│   │   │   ├── name_test.go
│   │   │   ├── resolve.go
│   │   │   └── zone.go
│   │   ├── config.go
│   │   ├── conn.go
│   │   ├── discovery.go
│   │   ├── discovery_inner.go
│   │   ├── doc
│   │   │   ├── NOTE.md
│   │   │   ├── README.md
│   │   │   ├── ROADMAP.md
│   │   │   └── concepts.md
│   │   ├── guard.go
│   │   ├── host.go
│   │   ├── instance.go
│   │   ├── model.go
│   │   ├── node.go
│   │   ├── nodes.go
│   │   ├── register.go
│   │   └── scheduler.go
│   ├── grpc-go
│   │   ├── Learning.md
│   │   ├── NOTE.md
│   │   ├── README.md
│   │   ├── doc
│   │   │   ├── NOTE.md
│   │   │   ├── feature.md
│   │   │   ├── load_balancing.md
│   │   │   └── name_resolver.md
│   │   ├── doc-to-read
│   │   │   └── roadmap.md
│   │   ├── error.md
│   │   ├── faq.md
│   │   ├── grpc
│   │   │   ├── README.md
│   │   │   ├── go-proxy
│   │   │   │   ├── README.md
│   │   │   │   ├── basic_model.go
│   │   │   │   ├── codec.go
│   │   │   │   ├── handler.go
│   │   │   │   ├── proxy_test.go
│   │   │   │   ├── register.go
│   │   │   │   └── streamdirector.go
│   │   │   └── reverseProxy
│   │   │       └── README.md
│   │   ├── name_resovler
│   │   │   ├── DOC.md
│   │   │   └── name.go
│   │   ├── pick_wrapper_cnt.go
│   │   ├── picker_wrapper.go
│   │   └── simple-rpc
│   │       └── README
│   ├── im
│   │   ├── MQ
│   │   │   ├── JobQueue
│   │   │   │   ├── QuickTutorial.md
│   │   │   │   ├── README.md
│   │   │   │   ├── buildRedisJobQueue
│   │   │   │   │   ├── README.md
│   │   │   │   │   ├── consumer.go
│   │   │   │   │   ├── example_test.go
│   │   │   │   │   ├── integration_test.go
│   │   │   │   │   ├── package.go
│   │   │   │   │   ├── queue.go
│   │   │   │   │   └── �\226\221�\203\221.md
│   │   │   │   ├── redis
│   │   │   │   │   └── main.go
│   │   │   │   └── simpleJobQueue
│   │   │   │       ├── simple
│   │   │   │       │   └── main.go
│   │   │   │       ├── simple2
│   │   │   │       │   └── main.go
│   │   │   │       └── standalone
│   │   │   │           └── main.go
│   │   │   ├── README.md
│   │   │   ├── abc_test.go
│   │   │   ├── gim
│   │   │   │   └── README.md
│   │   │   ├── goim
│   │   │   │   └── README.md
│   │   │   └── nsq
│   │   │       ├── QuickStart.md
│   │   │       └── nsqd
│   │   │           ├── README.md
│   │   │           └── nsqd.go
│   │   ├── README.md
│   │   └── goim
│   │       └── README.md
│   ├── pkg
│   │   ├── AccessControl
│   │   │   ├── README.md
│   │   │   └── grafana-access-control
│   │   │       └── README.md
│   │   ├── BitMap
│   │   │   └── README.md
│   │   ├── CircleArray
│   │   │   └── README.md
│   │   ├── collection
│   │   │   └── README.md
│   │   ├── go-rountine-pool
│   │   │   ├── README.md
│   │   │   └── workpool
│   │   │       └── README.md
│   │   ├── integerx
│   │   ├── meta-info
│   │   │   └── README.md
│   │   ├── notify
│   │   │   └── README.md
│   │   ├── process
│   │   │   └── README.md
│   │   ├── stringx
│   │   └── timex
│   ├── redis
│   │   ├── README.md
│   │   └── redisgo
│   │       ├── README.md
│   │       ├── conn.go
│   │       ├── idellist.go
│   │       ├── interface.go
│   │       └── poolConn.go
│   └── register
│       ├── NOTE.md
│       ├── consul
│       │   ├── NOTE.md
│       │   ├── config.go
│       │   ├── model.go
│       │   ├── proxy.go
│       │   ├── register.go
│       │   ├── resolver.go
│       │   └── watch.go
│       ├── lb
│       │   ├── NOTE.md
│       │   ├── banlancer.go
│       │   ├── base.go
│       │   ├── random.go
│       │   └── selector.go
│       └── proxy
│           ├── proxy.go
│           ├── proxy_test.go
│           └── server.go
├── Pattern-Design
│   ├── CircuitBreak
│   │   ├── README.md
│   │   ├── go-zero-breaker
│   │   │   ├── README.md
│   │   │   └── googlebreak.go
│   │   └── go-zero-rollingwindow
│   │       ├── README.md
│   │       ├── bucket.go
│   │       └── window.go
│   ├── ErrorRetry
│   ├── EventBus
│   │   ├── Event-tracking
│   │   │   ├── README.md
│   │   │   └── event.go
│   │   ├── README.md
│   │   ├── in-process-bus
│   │   │   └── README.md
│   │   ├── pub-sub
│   │   │   └── pubsub.go
│   │   ├── pub-sub-lib
│   │   │   ├── README.md
│   │   │   ├── exported.go
│   │   │   ├── model.go
│   │   │   └── unexported.go
│   │   └── simpleEventBus
│   │       ├── main.go
│   │       ├── readme
│   │       └── simpleEventBus_test.go
│   ├── MapReduce-All
│   │   ├── README.md
│   │   └── mapreduce
│   │       ├── 824-mrinput.txt
│   │       ├── README.md
│   │       ├── common.go
│   │       ├── mapreduce.go
│   │       ├── master.go
│   │       ├── test_test.go
│   │       └── worker.go
│   ├── fan-in&fan-out
│   │   └── README.md
│   ├── iterator
│   │   ├── README.md
│   │   └── columnIterator
│   ├── proirtyqueue
│   │   └── README.md
│   ├── queue
│   │   └── README.md
│   ├── redis-book
│   │   ├── SkipList
│   │   │   ├── README.md
│   │   │   └── main_test.go
│   │   └── bitmap
│   │       ├── README.md
│   │       └── main.go
│   ├── short-url
│   │   └── README.md
│   ├── shut-down
│   │   └── README.md
│   ├── singleton
│   │   ├── README.md
│   │   └── example
│   │       └── main.go
│   ├── spider
│   │   └── goquery
│   │       └── README.md
│   └── timewheel
│       ├── README.md
│       ├── safemap.go
│       ├── timewheel.go
│       └── timewheel_test.go
├── Pending
│   ├── Mydocker
│   │   └── READEME.md
│   ├── csapp
│   │   └── chatper1
│   │       └── main_test.go
│   └── utilmate-go
│       └── README.md
├── README.md
├── READMELOG.md
├── RoadMap.md
├── Snippet
│   ├── Log.md
│   ├── buffer-channel.md
│   ├── channel.md
│   ├── interview-collection.md
│   ├── interview-q1.md
│   ├── interview-qa-2.md
│   ├── roadmap.md
│   ├── watch.md
│   └── workpool.md
├── StandardLib
│   ├── README.md
│   ├── benchmark
│   │   └── README.md
│   ├── binary
│   │   └── README.md
│   ├── build
│   │   ├── README.md
│   │   └── main.go
│   ├── cgo
│   │   ├── READNE.md
│   │   ├── cgo-book
│   │   │   ├── example1
│   │   │   │   └── main.go
│   │   │   ├── example2
│   │   │   │   └── main.go
│   │   │   ├── example3
│   │   │   │   ├── hello.c
│   │   │   │   └── main.go
│   │   │   └── example4
│   │   │       ├── hello.c
│   │   │       ├── hello.cpp
│   │   │       └── hello.h
│   │   ├── cgo-document
│   │   │   ├── example1
│   │   │   └── example2
│   │   │       └── main.go
│   │   ├── idflag
│   │   │   ├── README.md
│   │   │   ├── build-version
│   │   │   └── main.go
│   │   ├── porting-go
│   │   │   ├── README.md
│   │   │   ├── hello.c
│   │   │   ├── hello.go
│   │   │   ├── helloc
│   │   │   ├── hellogo
│   │   │   ├── main.go
│   │   │   └── net
│   │   │       ├── main.go
│   │   │       └── net
│   │   ├── print
│   │   │   └── main.go
│   │   └── simple-cgo
│   │       └── main.go
│   ├── channel
│   │   ├── README.md
│   │   ├── REF.md
│   │   └── chan_test.go
│   ├── chore
│   │   ├── Escape
│   │   │   └── README.md
│   │   ├── IO
│   │   │   ├── FileCopy
│   │   │   │   ├── README.md
│   │   │   │   └── main.go
│   │   │   ├── InteractiveShell
│   │   │   │   ├── README.md
│   │   │   │   └── main.go
│   │   │   ├── ListenKeyboard
│   │   │   │   ├── README.md
│   │   │   │   └── main.go
│   │   │   └── README.md
│   │   ├── JSON
│   │   │   ├── BenchMark
│   │   │   │   └── ben_test.go
│   │   │   └── README.md
│   │   ├── bootstrap
│   │   ├── byte
│   │   │   └── bufferbyte
│   │   │       └── README.md
│   │   ├── config
│   │   │   └── README.md
│   │   ├── flag
│   │   │   ├── README.md
│   │   │   ├── Usage.md
│   │   │   ├── example2_test.go
│   │   │   ├── example_test.go
│   │   │   └── flag.go
│   │   ├── fmt
│   │   │   └── readme.md
│   │   ├── generic
│   │   │   └── README.md
│   │   ├── ioutil
│   │   │   └── tmpDir
│   │   │       └── tmp_test.go
│   │   ├── make-new
│   │   │   └── README.md
│   │   ├── marshall
│   │   │   └── README.md
│   │   ├── mysql
│   │   │   └── README.md
│   │   ├── pprof
│   │   │   └── README.md
│   │   ├── rand
│   │   │   └── shuffle
│   │   │       └── main.go
│   │   ├── recover
│   │   │   └── README.md
│   │   ├── string
│   │   │   ├── README.md
│   │   │   └── main.go
│   │   ├── stringutil
│   │   │   ├── main.go
│   │   │   └── main_test.go
│   │   └── syscall
│   │       └── README.md
│   ├── context
│   │   ├── README.md
│   │   ├── Summary.md
│   │   ├── basic.go
│   │   ├── cancel_test.go
│   │   ├── ctx
│   │   ├── example
│   │   │   └── main.go
│   │   ├── example2
│   │   │   └── main.go
│   │   ├── example3
│   │   │   └── main.go
│   │   ├── example4
│   │   │   └── main.go
│   │   ├── simultaneous_test.go
│   │   └── usage
│   │       ├── deadline.go
│   │       ├── ignorectx.go
│   │       └── value.go
│   ├── defer
│   │   ├── README.md
│   │   └── pitfall
│   │       └── test.go
│   ├── errors
│   │   ├── README.md
│   │   ├── RESOURCE.md
│   │   ├── docker
│   │   │   ├── README.md
│   │   │   ├── code.go
│   │   │   ├── errors.go
│   │   │   └── register.go
│   │   ├── error.go
│   │   ├── error_better_example
│   │   │   └── error_example_1_test.go
│   │   ├── error_better_example2
│   │   │   └── exmaple_test.go
│   │   ├── error_example1
│   │   │   └── example_test.go
│   │   ├── error_test.go
│   │   ├── etcd
│   │   │   └── errors.go
│   │   ├── handleError
│   │   │   ├── err_test.go
│   │   │   └── errors_test.go
│   │   ├── lib
│   │   ├── mockPkgErrors
│   │   │   ├── README.md
│   │   │   ├── exported.go
│   │   │   ├── exported_test.go
│   │   │   ├── frame.go
│   │   │   ├── fundamental.go
│   │   │   ├── model.go
│   │   │   ├── stack.go
│   │   │   ├── stack_trace.go
│   │   │   ├── withMessage.go
│   │   │   └── withstack.go
│   │   └── tipsAndTrick
│   │       ├── aggregationError.go
│   │       ├── bufioError.go
│   │       ├── pitfall_test.go
│   │       └── switch_error_test.go
│   ├── http
│   │   ├── README.md
│   │   ├── eof
│   │   │   ├── README.md
│   │   │   ├── eof.go
│   │   │   ├── eof_test.go
│   │   │   └── gnet_test.go
│   │   ├── golangHttp
│   │   │   ├── client.go
│   │   │   ├── client_test.go
│   │   │   ├── connectMethod.go
│   │   │   ├── cookiejar.go
│   │   │   ├── persistConn.go
│   │   │   ├── readme.md
│   │   │   ├── request.go
│   │   │   ├── response.go
│   │   │   ├── roudntrippter.go
│   │   │   ├── transport.go
│   │   │   ├── url.go
│   │   │   └── wantConn.go
│   │   └── reverseproxy
│   │       ├── README.md
│   │       └── base_test.go
│   ├── interface
│   │   ├── README.md
│   │   ├── example
│   │   │   └── main_test.go
│   │   ├── example2
│   │   │   └── base_test.go
│   │   ├── iface
│   │   │   └── main.go
│   │   ├── law_of_reflection_test.go
│   │   ├── mock-interface
│   │   │   ├── main.go
│   │   │   └── main_test.go
│   │   └── typeReadme.md
│   ├── map
│   │   ├── README.md
│   │   ├── go-zero-map
│   │   ├── map_test.go
│   │   ├── orderedMap
│   │   │   └── main.go
│   │   └── original
│   │       ├── README.md
│   │       ├── const.go
│   │       └── map.go
│   ├── sync
│   │   ├── README.md
│   │   ├── atomic
│   │   │   ├── README.md
│   │   │   ├── atomic_test.go
│   │   │   └── example_test.go
│   │   ├── map
│   │   │   ├── README.md
│   │   │   ├── bug_report.md
│   │   │   ├── deep_map.go
│   │   │   ├── interface.go
│   │   │   ├── main_test.go
│   │   │   ├── map.go
│   │   │   ├── map_test.go
│   │   │   ├── model.go
│   │   │   └── rm_map.go
│   │   ├── mutex
│   │   ├── once
│   │   │   └── README.md
│   │   ├── pool
│   │   ├── pool.md
│   │   └── sync_map
│   │       ├── README.md
│   │       └── concurrent_map.go
│   └── time
│       ├── README.md
│       ├── main_test.go
│       ├── time_test.go
│       └── timer
├── Tools.md
├── blogger.md
├── cheatsheet.md
└── devops
    ├── Linux
    │   ├── Bash&Shell
    │   │   ├── README.md
    │   │   ├── back-up.sh
    │   │   ├── chapter1
    │   │   ├── chapter2
    │   │   │   ├── README.md
    │   │   │   ├── cat.sh
    │   │   │   ├── checksum.sh
    │   │   │   ├── data.txt
    │   │   │   ├── example.txt
    │   │   │   ├── interactive.sh
    │   │   │   ├── mktmp.sh
    │   │   │   ├── rename.sh
    │   │   │   ├── sort.sh
    │   │   │   ├── split.sh
    │   │   │   ├── tr.sh
    │   │   │   ├── uniq.sh
    │   │   │   └── xargs.sh
    │   │   ├── chapter3
    │   │   │   ├── A.txt
    │   │   │   ├── B.txt
    │   │   │   ├── README.md
    │   │   │   ├── chown.sh
    │   │   │   ├── comm.sh
    │   │   │   └── dd.sh
    │   │   ├── readIFs.sh
    │   │   ├── test2.sh
    │   │   └── ttt.sh
    │   ├── LinuxCmd
    │   │   └── README.md
    │   ├── README.md
    │   ├── Roadmap.md
    │   ├── Tools-Quick-View.md
    │   ├── command
    │   │   ├── base64
    │   │   ├── cat
    │   │   │   ├── README.md
    │   │   │   └── doc.md
    │   │   ├── curl
    │   │   │   └── README.md
    │   │   ├── cut
    │   │   │   └── README.md
    │   │   ├── diff
    │   │   ├── grep
    │   │   │   ├── README.md
    │   │   │   ├── sed_data.txt
    │   │   │   ├── silent_grep.sh
    │   │   │   └── student_data.txt
    │   │   ├── iostat
    │   │   │   └── README.md
    │   │   ├── ln
    │   │   │   └── README.md
    │   │   ├── ls
    │   │   │   └── ls.sh
    │   │   ├── md5
    │   │   ├── mkfs.ext4
    │   │   ├── pidstat
    │   │   │   └── README.md
    │   │   ├── ps
    │   │   ├── rsync
    │   │   │   └── rsync.md
    │   │   ├── sar
    │   │   │   └── README.md
    │   │   ├── sed
    │   │   │   ├── README.md
    │   │   │   └── op-�\226\207件�\204�\220\206.md
    │   │   ├── sort
    │   │   ├── ssh\ 
    │   │   │   └── README.md
    │   │   ├── tr
    │   │   └── xargs
    │   │       ├── README.md
    │   │       ├── args.txt
    │   │       ├── echo.sh
    │   │       └── xargs.sh
    │   ├── gdb
    │   │   └── README.md
    │   ├── git
    │   │   └── gc.md
    │   ├── io
    │   │   ├── README.md
    │   │   ├── dd
    │   │   │   └── README.md
    │   │   ├── df
    │   │   │   └── README.md
    │   │   ├── dstat
    │   │   │   └── README.md
    │   │   ├── du
    │   │   │   └── README.md
    │   │   ├── fdisk
    │   │   │   └── README.md
    │   │   ├── file
    │   │   │   ├── README.md
    │   │   │   └── filestat.sh
    │   │   ├── free
    │   │   │   └── README.md
    │   │   ├── lsblk
    │   │   │   └── README.md
    │   │   ├── lsof
    │   │   │   └── README.md
    │   │   ├── mount
    │   │   ├── tcpdump
    │   │   └── vmstat
    │   ├── network
    │   │   ├── ifconfig
    │   │   │   └── README.md
    │   │   ├── ip
    │   │   ├── netstat
    │   │   │   └── README.md
    │   │   ├── network
    │   │   ├── ping
    │   │   ├── ss
    │   │   └── traceroute
    │   │       └── README.md
    │   ├── objdump
    │   │   ├── README.md
    │   │   ├── SimpleSection.c
    │   │   └── SimpleSection.o
    │   └── system
    │       ├── ld
    │       ├── ln
    │       │   └── README.md
    │       ├── pidstat
    │       │   └── README.md
    │       ├── ppref
    │       │   └── README.md
    │       ├── strace
    │       │   └── README.md
    │       ├── systat
    │       ├── top
    │       │   └── README.md
    │       ├── uname
    │       │   └── README.md
    │       └── uptime
    ├── ansible
    │   └── README.md
    ├── cloud
    │   ├── nat.md
    │   └── sls.md
    ├── docker
    │   ├── README.md
    │   └── docker-network.md
    ├── grafana
    │   ├── NOTE
    │   │   └── op.md
    │   ├── README.md
    │   └── grafana.md
    ├── k8s-summary
    │   ├── Debug.md
    │   ├── IMAGE\ 2021-09-20\ 17:17:42.jpg
    │   ├── Pending.md
    │   ├── Q&A.md
    │   ├── README.md
    │   ├── index.md
    │   ├── k8s-doc
    │   │   ├── debugging.md
    │   │   ├── deployment.md
    │   │   ├── download-api.md
    │   │   ├── get-shell.md
    │   │   ├── index.md
    │   │   ├── k8s-认�\201�\234��\210��\222\214Service-account.md
    │   │   ├── kubeconfig.md
    │   │   ├── stackdriver.md
    │   │   ├── �\200�\210�\230�-ingress.md
    │   │   ├── �\200�\210�\230�-stateful-set.md
    │   │   ├── 客�\210�端�\217\221�\216�-pod-�\200\232信.md
    │   │   ├── 容�\231�设计模�\217.md
    │   │   ├── �\226��\200代�\233\221�\216��\236��\236\204.md
    │   │   ├── �\221�\234�\226�\225�.md
    │   │   ├── �\212\202�\202�容�\215度.md
    │   │   ├── �\212\202�\202��\203度.md
    │   │   └── �\230�\217��\224�.md
    │   ├── k8s-in-action
    │   │   ├── README.md
    │   │   ├── client-go.md
    │   │   ├── commend.md
    │   │   ├── concept.md
    │   │   ├── configmap.md
    │   │   ├── google-k8s-practice.md
    │   │   ├── helm.md
    │   │   ├── ingress.md
    │   │   ├── k8s-dashboard.md
    │   │   ├── k8s-�\221�\234.md
    │   │   ├── log-arch.md
    │   │   ├── metric-server.md
    │   │   ├── minikube-op.md
    │   │   ├── operator.md
    │   │   ├── pod-�\224记.md
    │   │   ├── pv-and-pvc.md
    │   │   ├── reouse-limitation.md
    │   │   ├── secret.md
    │   │   ├── stateful-set.md
    │   │   ├── what-is-config-map.md
    │   │   ├── what-is-safe-context.md
    │   │   ├── �\230�\202��\215�.md
    │   │   ├── �\216��\222\210.md
    │   │   ├── �\234\215�\212��\217\221�\216�.md
    │   │   ├── �\226\221�\203\221.md
    │   │   ├── �\233\221�\216�.md
    │   │   └── �\221�\234�\226�\225�.md
    │   ├── k8s-op
    │   │   └── �\237�询pod�\207签.md
    │   ├── k8s-qa
    │   │   ├── docker-namespace.md
    │   │   ├── grpc-readiness-health.md
    │   │   ├── minikube-�\224\231误.md
    │   │   ├── node-debug.md
    │   │   └── 设置�\225\234�\203\217�\216��\213\211�\206�\222�.md
    │   ├── learning-ladder.md
    │   └── troubleshooting
    │       ├── dns.md
    │       ├── k8s-端�\217��\215��\224��\227��\230.md
    │       ├── network.md
    │       └── �\234\215�\212��\236�\200\232�\200�.md
    ├── nginx
    │   └── README.md
    ├── promethues
    │   ├── Note
    │   │   └── Op.md
    │   └── README.md
    ├── storage
    │   ├── Mysql
    │   │   ├── README.md
    │   │   ├── innodb.md
    │   │   ├── jike-note
    │   │   │   └── note.md
    │   │   ├── learning-doc.md
    │   │   ├── note
    │   │   │   ├── FAQ.md
    │   │   │   ├── book.md
    │   │   │   ├── case-study.md
    │   │   │   ├── index.md
    │   │   │   ├── lock.md
    │   │   │   ├── metrics.md
    │   │   │   ├── mvcc.md
    │   │   │   ├── orderby.md
    │   │   │   ├── rank.md
    │   │   │   ├── read-write.md
    │   │   │   ├── replication.md
    │   │   │   ├── sql.md
    │   │   │   ├── unread.md
    │   │   │   └── update.md
    │   │   └── note.md
    │   ├── README.md
    │   ├── lsm.md
    │   ├── redis
    │   └── why-b+.md
    └── �\200\203�\201
        ├── aws.md
        └── cka.md
```


### 目录结构 1 
```
├── BEC
│   ├── README.md
│   ├── example1
│   │   ├── main
│   │   └── main.go
│   ├── example2
│   │   ├── example2
│   │   └── example2.go
│   ├── example3
│   │   ├── example3
│   │   └── example3.go
│   ├── example4
│   │   └── example4.go
│   └── example5
│       └── example5.go
├── Bug
│   └── pustToSessionbug
│       ├── README.md
│       └── copy_test.go
├── Concurrency
│   ├── README.md
│   ├── channel
│   │   └── test_close.go
│   ├── concurrency-in-go
│   │   ├── README.md
│   │   ├── chapter2
│   │   │   ├── dead-lock
│   │   │   │   └── dead_test.go
│   │   │   ├── live-lock
│   │   │   │   └── live_lock_test.go
│   │   │   └── starvation
│   │   │       └── stavation_test.go
│   │   ├── chapter3
│   │   │   ├── avalible-same
│   │   │   │   └── space_test.go
│   │   │   ├── channel
│   │   │   │   └── channel_test.go
│   │   │   ├── cond
│   │   │   │   └── cond_test.go
│   │   │   ├── context-switch
│   │   │   │   └── ben_test.go
│   │   │   ├── join-point
│   │   │   │   └── join_point_test.go
│   │   │   ├── memory_consumer
│   │   │   │   └── mem_consumer_test.go
│   │   │   ├── mutex
│   │   │   │   └── mutext_test.go
│   │   │   ├── once
│   │   │   │   └── once_test.go
│   │   │   ├── pool_test.go
│   │   │   ├── select\ 
│   │   │   │   └── select_test.go
│   │   │   └── test-wait-group
│   │   │       └── test-group_test.go
│   │   ├── chapter4
│   │   │   ├── confinement
│   │   │   │   └── confiement_test.go
│   │   │   ├── errorHandler
│   │   │   │   └── error_test.go
│   │   │   ├── generator
│   │   │   │   └── generator_test.go
│   │   │   ├── pipeline
│   │   │   │   └── pipeline_test.go
│   │   │   ├── preventLeak
│   │   │   │   └── prevent_test.go
│   │   │   └── repeat
│   │   │       └── repeat_test.go
│   │   ├── chapter5
│   │   │   ├── bridge
│   │   │   │   └── bridge_test.go
│   │   │   ├── doWork
│   │   │   │   └── dowork_test.go
│   │   │   ├── errorHandler
│   │   │   │   └── error_test.go
│   │   │   ├── fan
│   │   │   │   └── fan_test.go
│   │   │   ├── heartbeat
│   │   │   │   └── heartbeat_test.go
│   │   │   ├── orDone
│   │   │   │   └── or_done_test.go
│   │   │   ├── ratelimit
│   │   │   │   └── main.go
│   │   │   ├── ratelimit1
│   │   │   │   └── ratelimit1_test.go
│   │   │   ├── replicate
│   │   │   │   └── replicate_test.go
│   │   │   └── tee
│   │   │       └── tee_test.go
│   │   └── chapter6
│   │       └── chapter6_test.go
│   └── forkjoinmodel
│       └── README.md
├── DataBase
│   └── badger
├── FrameWork
│   ├── EmbeddedResources
│   │   ├── README.md
│   │   ├── gen.go
│   │   └── resources
│   ├── TextSimilarity
│   │   └── consine
│   │       ├── README.md
│   │       ├── main.go
│   │       └── sin_test.go
│   ├── binddata
│   │   └── GoGenerate
│   │       ├── README.md
│   │       ├── gopher.go
│   │       ├── gopher.y
│   │       └── y.output
│   └── tokenizer-nlp
│       ├── README.md
│       ├── jieba
│       │   ├── jieba_test.go
│       │   ├── stop_word.dict.utf8
│       │   └── user.dict.utf8
│       ├── sego
│       │   └── sego_test.go
│       └── vocabulary
│           └── voc_test.go
├── GoCommand
│   └── idflags
│       └── README.md
├── GolangBasic
│   ├── JSON
│   │   ├── BenchMark
│   │   │   └── ben_test.go
│   │   └── README.md
│   ├── README.md
│   ├── benchmark
│   │   └── README.md
│   ├── bootstrap
│   ├── build
│   │   ├── README.md
│   │   ├── main.go
│   │   └── main.go.bak
│   ├── byte
│   │   └── bufferbyte
│   │       └── README.md
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
│   ├── confOpt
│   ├── context
│   │   ├── README.md
│   │   ├── basic.go
│   │   ├── cancel_test.go
│   │   ├── example
│   │   │   └── main.go
│   │   ├── example2
│   │   │   └── main.go
│   │   ├── example3
│   │   │   └── main.go
│   │   ├── example4
│   │   │   └── main.go
│   │   └── simultaneous_test.go
│   ├── defer
│   │   └── pitfall
│   │       └── test.go
│   ├── dive_into_go
│   │   ├── README.md
│   │   └── how_does_go
│   │       └── main.go
│   ├── flag
│   │   └── README.md
│   ├── interface
│   │   ├── README.md
│   │   ├── law_of_reflection_test.go
│   │   └── typeReadme.md
│   ├── ioutil
│   │   └── tmpDir
│   │       └── tmp_test.go
│   ├── map
│   │   ├── README.md
│   │   └── orderedMap
│   │       └── main.go
│   ├── rand
│   │   └── shuffle
│   │       └── main.go
│   ├── sort
│   │   └── simpleSort
│   ├── string
│   │   ├── README.md
│   │   └── main.go
│   ├── stringutil
│   │   ├── main.go
│   │   └── main_test.go
│   └── sync
│       └── README.md
├── HTTP
│   ├── RateLimit
│   │   ├── README.md
│   │   ├── REAME.md
│   │   ├── allratelimiter.go
│   │   ├── leakybuckUber
│   │   ├── leakybucket
│   │   │   ├── README.md
│   │   │   ├── leakybucket.go
│   │   │   └── leakybucket_test.go
│   │   ├── throttler
│   │   │   ├── remote
│   │   │   │   ├── options.go
│   │   │   │   ├── options_test.go
│   │   │   │   ├── throttler.go
│   │   │   │   └── throttler_test.go
│   │   │   └── throttler.go
│   │   └── tokenbucket
│   │       ├── README.md
│   │       └── tokentbucket.go
│   ├── httpclient
│   │   ├── README.md
│   │   └── httpclient.go
│   ├── httpclientMiddelWare
│   └── loghttpRequest
│       └── main.go
├── IO
│   ├── FileCopy
│   │   ├── README.md
│   │   └── main.go
│   ├── InteractiveShell
│   │   ├── README.md
│   │   └── main.go
│   ├── ListenKeyboard
│   │   ├── README.md
│   │   └── main.go
│   └── README.md
├── Linux
│   ├── Bash&Shell
│   │   ├── README.md
│   │   ├── back-up.sh
│   │   ├── chapter1
│   │   ├── chapter2
│   │   │   ├── README.md
│   │   │   ├── example.txt
│   │   │   └── xargs.sh
│   │   ├── readIFs.sh
│   │   ├── test2.sh
│   │   └── ttt.sh
│   ├── LinuxCmd
│   │   └── README.md
│   ├── README.md
│   ├── Tools-Quick-View.md
│   ├── command
│   │   ├── curl
│   │   ├── dd
│   │   ├── du
│   │   ├── free
│   │   │   └── README.md
│   │   ├── lsof
│   │   │   └── README.md
│   │   ├── ping
│   │   ├── ps
│   │   ├── systat
│   │   ├── top
│   │   │   └── README.md
│   │   ├── uptime
│   │   ├── vmstat
│   │   └── xargs
│   │       ├── README.md
│   │       ├── args.txt
│   │       ├── echo.sh
│   │       └── xargs.sh
│   └── gdb
├── MQ
│   ├── JobQueue
│   │   ├── QuickTutorial.md
│   │   ├── README.md
│   │   ├── buildRedisJobQueue
│   │   │   ├── README.md
│   │   │   ├── consumer.go
│   │   │   ├── example_test.go
│   │   │   ├── integration_test.go
│   │   │   ├── package.go
│   │   │   ├── queue.go
│   │   │   └── �\226\221�\203\221.md
│   │   ├── redis
│   │   │   └── main.go
│   │   └── simpleJobQueue
│   │       ├── simple
│   │       │   └── main.go
│   │       ├── simple2
│   │       │   └── main.go
│   │       └── standalone
│   │           └── main.go
│   ├── README.md
│   ├── abc_test.go
│   ├── gim
│   │   └── README.md
│   ├── goim
│   │   └── README.md
│   └── nsq
│       ├── QuickStart.md
│       └── nsqd
│           ├── README.md
│           └── nsqd.go
├── Pattern
│   ├── ErrorRetry
│   ├── EventBus
│   │   ├── README.md
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
│   ├── LRU
│   │   └── README.md
│   ├── fan-in&fan-out
│   │   └── README.md
│   ├── iterator
│   │   ├── README.md
│   │   └── columnIterator
│   ├── mapreduce
│   │   └── README.md
│   └── singleton
│       ├── README.md
│       └── example
│           └── main.go
├── README.md
├── RoadMap.md
├── WebTech
│   ├── LoadBalanceUsage
│   │   ├── loadbalance
│   │   │   ├── README.md
│   │   │   └── main.go
│   │   └── loadbalanceGoIm
│   │       ├── main.go
│   │       └── main_test.go
│   ├── breaker
│   │   └── README.md
│   └── decorate
│       └── README.md
└── utilmate-go
    └── README.md
```

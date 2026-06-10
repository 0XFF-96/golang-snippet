### 介绍
asynccache` fetches and updates the latest data periodically and supports expire a key if unused for a period.

from bytedance/gopkg 

### 有趣的特点
1、定时刷新和定时删除逻辑
2、全局共用的 ticker 
3、数据存储的方式 `data sync.Map` 
(只适合内存占用不多的 cache, 且没有 cache 内存自动回收机制)
4、简单的 singleflight 算法的应用
```
var (
	// 共用 ticker
	refreshTickerMap, expireTickerMap sync.Map
)
```


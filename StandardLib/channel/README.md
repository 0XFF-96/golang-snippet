### select-inside-select 的设计
一、解析一下，这种设计的好处？
```
		select {
		case ch <- res:
		default:
			select {
			case <-ch:
			default:
			}
			ch <- res
		}
	}
```


### Channel 总结
1. https://segmentfault.com/a/1190000017958702
2. https://lessisbetter.site/about/  关于 channel 的相关源码
3. https://colobu.com/2016/04/14/Golang-Channels/

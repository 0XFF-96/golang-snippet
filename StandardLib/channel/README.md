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
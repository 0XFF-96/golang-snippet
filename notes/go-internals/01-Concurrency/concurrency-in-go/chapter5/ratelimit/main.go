package ratelimit

import (
	"context"
	"log"
	"os"
	"sync"

	"source-code-reading/time/rate"
)

// page 193

func Open() *APIConnection {
	return &APIConnection{
		rateLimiter:rate.NewLimiter(rate.Limit(1),1),
	}
}

type APIConnection struct {
	// 新版的源码好想有变
	rateLimiter *rate.Limiter
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	// Pretend we do work here
	if err := a.rateLimiter.Wait(ctx);err != nil {
		return err
	}
	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	// Pretend we do work here
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}


func main(){
	defer log.Printf("Done.")

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i:=0;i < 10;i++{
		go func(){
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("connot ReadFile: %v", err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i :=0; i < 10; i ++{
		go func(){
			defer wg.Done()
			// err := apiConnection.ReadFile(context.Background())
			//if err != nil {
			//	log.Printf("cannot ResolvAddress:%v", err)
			//}
			//log.Printf("ResolveAddress")
		}()
	}
	wg.Wait()
}


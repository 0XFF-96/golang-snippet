package ratelimit1

import (
	"context"
	"log"
	"os"
	"sync"
)


// page 193

func Open() *APIConnection {
	return &APIConnection{}
}

type APIConnection struct {
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	// Pretend we do work here
	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	// Pretend we do work here
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
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot ResolvAddress:%v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}
	wg.Wait()
}

package pool

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	maxCoroutines  = 25
	pooledResource = 2
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

func TestPool(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(maxCoroutines)

	p, err := New(createConnection, pooledResource)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxCoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Shutdown program.")
	p.Close()
}

func performQueries(query int, p *Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	//fmt.Println()
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection))
}

func TestLog(t *testing.T) {
	log.Println("this is a test", 1, 2, 3, "haha")
}

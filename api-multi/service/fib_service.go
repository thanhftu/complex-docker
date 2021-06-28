package service

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/thanhftu/api-multi/domain/fib"
	"github.com/thanhftu/worker"
)

var RedisSource string

func init() {
	RedisSource = "localhost:6379"
	if v := os.Getenv("REDIS_ADDR"); v != "" {
		RedisSource = v
	}

	fmt.Println("Redis url: ", RedisSource)

}

func GetFibFromRedisWorker(index string) (int64, error) {

	val64, err := worker.WorkerRedisFib(index, RedisSource)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return val64, nil
}

func SaveFib(index string) (*fib.FibNumber, error) {

	index64, _ := strconv.ParseInt(index, 10, 64)
	val64, _ := GetFibFromRedisWorker(index)
	// date_created := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	fibnumber := &fib.FibNumber{
		Index: index64,
		Value: val64,
	}
	if err := fibnumber.SAVE(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return fibnumber, nil
}

func GetFib(index string) (*fib.FibNumber, error) {
	index64, _ := strconv.ParseInt(index, 10, 64)
	fibnumber := &fib.FibNumber{
		Index: index64,
	}
	if err := fibnumber.GET(); err != nil {
		return nil, err
	}
	return fibnumber, nil
}
func GetLatest() (*fib.FibNumber, error) {
	var fiblastest fib.FibNumber
	if err := fiblastest.GETLATEST(); err != nil {
		return nil, err
	}
	return &fiblastest, nil
}
func GetAllFib() ([]fib.FibNumber, error) {
	return fib.GETALL()
}

func DeleteFib(ID string) error {
	ID64, _ := strconv.ParseInt(ID, 10, 64)
	fibnumber := &fib.FibNumber{
		ID: ID64,
	}
	if err := fibnumber.DELETE(); err != nil {
		return err
	}
	return nil
}

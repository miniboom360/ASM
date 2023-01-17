package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	opentracing_log "github.com/opentracing/opentracing-go/log"
	"io/ioutil"
	"time"
	// exampletasks "github.com/RichardKnop/machinery/v2/example/tasks"
	"github.com/RichardKnop/machinery/v2/example/tracers"
	"github.com/RichardKnop/machinery/v2/log"
	"github.com/RichardKnop/machinery/v2/tasks"
)

const (
	test_config         = "/Users/liyang/tools/asm/ASM/backend/workers/common/config.json"
	windows_test_config = "D:\\code\\asm-demo\\backend\\workers\\common\\config.json"
	production_config   = "/config.json"
)

type MachineryConf struct {
	Redis Redis `json:"Redis"`
}
type Redis struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	Queue    int    `json:"queue"`
}

func startServer(consumerQueue string, mTasks map[string]interface{}) (*machinery.Server, error) {
	cnf := &config.Config{
		DefaultQueue:    consumerQueue,
		ResultsExpireIn: 3600,
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}

	// m, err := ReadConfig(test_config)
	m, err := ReadConfig(production_config)
	if err != nil {
		return nil, err
	}

	// Create server instance
	redisUrl := fmt.Sprintf("%s@%s", m.Redis.Password, m.Redis.Address)
	broker := redisbroker.NewGR(cnf, []string{redisUrl}, m.Redis.Queue)
	backend := redisbackend.NewGR(cnf, []string{redisUrl}, m.Redis.Queue)
	lock := eagerlock.New()
	server := machinery.NewServer(cnf, broker, backend, lock)

	return server, server.RegisterTasks(mTasks)
}

// 这个worker一次只执行一个task，其他task等等待
func MchmultipleWorker(consumerQueue string, mTasks map[string]interface{}) error {
	cleanup, err := tracers.SetupTracer(consumerQueue)
	if err != nil {
		log.FATAL.Fatalln("[Handler][Mch] Unable to instantiate a tracer:", err)
	}
	defer cleanup()

	server, err := startServer(consumerQueue, mTasks)
	if err != nil {
		return err
	}

	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	// concurrency = 0，不限制每个worker的task数量
	worker := server.NewWorker(consumerQueue, 0)

	// Here we inject some custom code for error handling,
	// start and end of task hooks, useful for metrics for example.
	errorhandler := func(err error) {
		log.ERROR.Println("[Handler][Mch] I am an error handler:", err)
	}

	pretaskhandler := func(signature *tasks.Signature) {
		log.INFO.Println("[Handler][Mch] I am a start of task handler for:", signature.Name)
	}

	posttaskhandler := func(signature *tasks.Signature) {
		log.INFO.Println("[Handler][Mch] I am an end of task handler for:", signature.Name)
	}

	worker.SetPostTaskHandler(posttaskhandler)
	worker.SetErrorHandler(errorhandler)
	worker.SetPreTaskHandler(pretaskhandler)

	return worker.Launch()
}

func MchOneWorker(consumerQueue string, mTasks map[string]interface{}) error {
	cleanup, err := tracers.SetupTracer(consumerQueue)
	if err != nil {
		log.FATAL.Fatalln("[Handler][Mch] Unable to instantiate a tracer:", err)
	}
	defer cleanup()

	server, err := startServer(consumerQueue, mTasks)
	if err != nil {
		return err
	}

	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	worker := server.NewWorker(consumerQueue, 1)

	// Here we inject some custom code for error handling,
	// start and end of task hooks, useful for metrics for example.
	errorhandler := func(err error) {
		log.ERROR.Println("[Handler][Mch] I am an error handler:", err)
	}

	pretaskhandler := func(signature *tasks.Signature) {
		log.INFO.Println("[Handler][Mch] I am a start of task handler for:", signature.Name)
	}

	posttaskhandler := func(signature *tasks.Signature) {
		log.INFO.Println("[Handler][Mch] I am an end of task handler for:", signature.Name)
	}

	worker.SetPostTaskHandler(posttaskhandler)
	worker.SetErrorHandler(errorhandler)
	worker.SetPreTaskHandler(pretaskhandler)

	return worker.Launch()
}

// results1 := machinery2.MchClient("helloworld", task1, false)
// fmt.Println(results1)
func ReadConfig(configPath string) (*MachineryConf, error) {
	m := new(MachineryConf)
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// 读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func MchClient(serviceName string, myTaskSign tasks.Signature, nowait bool) ([]byte, error) {
	cleanup, err := tracers.SetupTracer(serviceName)
	if err != nil {
		log.FATAL.Fatal("[asm] Unable to instantiate a tracer:", err)
		return nil, err
	}
	defer cleanup()
	cnf := &config.Config{
		DefaultQueue:    serviceName,
		ResultsExpireIn: 3600,
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}
	// Create server instance
	m, err := ReadConfig(test_config)
	if err != nil {
		return nil, err
	}

	// Create server instance
	redisUrl := fmt.Sprintf("%s@%s", m.Redis.Password, m.Redis.Address)
	broker := redisbroker.NewGR(cnf, []string{redisUrl}, m.Redis.Queue)
	backend := redisbackend.NewGR(cnf, []string{redisUrl}, m.Redis.Queue)
	lock := eagerlock.New()
	machineryServer := machinery.NewServer(cnf, broker, backend, lock)

	/*
	 * Lets start a span representing this run of the `send` command and
	 * set a batch id as baggage so it can travel all the way into
	 * the worker functions.
	 */
	span, ctx := opentracing.StartSpanFromContext(context.Background(), serviceName)
	defer span.Finish()

	batchID := uuid.New().String()
	span.SetBaggageItem("batch.id", batchID)
	span.LogFields(opentracing_log.String("batch.id", batchID))

	log.INFO.Println("[Machinery]Starting batch:", batchID)
	log.INFO.Println("[Machinery]Single task:")
	asyncResult, err := machineryServer.SendTaskWithContext(ctx, &myTaskSign)
	if err != nil {
		log.FATAL.Fatal("[Machinery]Could not send task: %s", err.Error())
		return nil, err
	}
	if nowait {
		log.INFO.Println("[Machinery]Send task no wait", asyncResult)
		return nil, err
	} else {
		log.INFO.Println("[Machinery]Send task waiting...")
		results, err := asyncResult.Get(time.Duration(time.Millisecond * 1000))
		if err != nil {
			log.FATAL.Fatal("[Machinery] Getting task result failed with error: ", err.Error())
		}

		return results[0].Bytes(), nil
	}
}

func TaskBytes(m []tasks.Arg) ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		log.FATAL.Fatal("[Machinery][TaskBytes] error:", err.Error())
	}
	return b, err
}

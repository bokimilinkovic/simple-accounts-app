package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adjust/rmq"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	"github.com/bokimilinkovic/simple-accounts-app/pkg/database/gorm"
	"github.com/bokimilinkovic/simple-accounts-app/pkg/database/redis"
	"github.com/bokimilinkovic/simple-accounts-app/pkg/model"
	g "gorm.io/gorm"
)

const (
	prefetchLimit = 1000
	pollDuration  = 100 * time.Millisecond
	numConsumers  = 1 // number of consumers

	reportBatchSize = 10000
	consumeDuration = time.Millisecond
	shouldLog       = false
)

var namespace = uuid.NameSpaceURL

func init() {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // config file path
	viper.AutomaticEnv()     // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
}

// Consumer represents struct that will consume all messages from message queue and store in local cache and his database.
type Consumer struct {
	name   string
	db     *g.DB
	before time.Time
	cache  map[string]model.Transaction
}

// NewConsumer creates new consumer with tag and database
func NewConsumer(tag int, db *g.DB) *Consumer {
	return &Consumer{
		name:   fmt.Sprintf("consumer%d", tag),
		db:     db,
		before: time.Now(),
		cache:  make(map[string]model.Transaction),
	}
}

// Consume starts consuming messages. Unmarshals transaction and store in it's database.
func (c *Consumer) Consume(delivery rmq.Delivery) {
	payload := delivery.Payload()

	var transaction model.Transaction
	err := json.Unmarshal([]byte(payload), &transaction)
	if err != nil {
		fmt.Printf("erorr unamrshaling: %s", err.Error())
		if ok := delivery.Reject(); !ok {
			fmt.Printf("erorr rejecting: %s", err.Error())
		}
	}

	id := uuid.New().String()
	transaction.ID = id
	transaction.Status = model.Successfull
	if err := c.db.Create(&transaction).Error; err != nil {
		log.Fatal("error creating transaction:" + err.Error())
	}

	time.Sleep(consumeDuration)

	if !delivery.Ack() {
		log.Printf("failed to ack %s: %s", payload, err)
	} else {
		log.Printf("acknowledged %s", payload)
	}

	c.cache[transaction.ID] = transaction
}

func main() {

	var dbConfig gorm.Config
	if err := viper.Sub("database").Unmarshal(&dbConfig); err != nil {
		panic(err)
	}

	db, err := gorm.CreateConnection(dbConfig)
	if err != nil {
		panic(err)
	}

	var redisConfig redis.RedisConfig
	if err := viper.Sub("redis").Unmarshal(&redisConfig); err != nil {
		panic(err)
	}

	// ctx := context.Background()

	connection := rmq.OpenConnection("consumer", "tcp", fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port), 2)
	queue := connection.OpenQueue("transactions")

	queue.StartConsuming(prefetchLimit, 3)

	for i := 0; i < numConsumers; i++ {
		name := fmt.Sprintf("consumer %d", i)
		queue.AddConsumer(name, NewConsumer(i, db))
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	defer signal.Stop(signals)

	<-signals
	go func() {
		<-signals
		os.Exit(1)
	}()
}

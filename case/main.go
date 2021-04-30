package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	kafkago "github.com/segmentio/kafka-go"
)

var (
	client string
	mode   string

	brokers   string
	topic     string
	partition int

	msgSize     int
	numMessages int

	//value []byte
)

//func newUUID() (string, error) {
//	uuid := make([]byte, 16)
//	n, err := io.ReadFull(crand.Reader, uuid)
//	if n != len(uuid) || err != nil {
//		return "", err
//	}
//	uuid[8] = uuid[8]&^0xc0 | 0x80
//	uuid[6] = uuid[6]&^0xf0 | 0x40
//	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
//}


func produceKafkaGo() {

	// this apparently produces synchronously and is very slow - is there a way to
	// produce synchronously.

	w := kafkago.NewWriter(kafkago.WriterConfig{
		Brokers: []string{brokers},

		Topic:   "sunday2",
		// match other clients:
		Balancer:     &kafkago.Hash{},
		BatchTimeout: time.Duration(100) * time.Millisecond,
		// these are low by default on this client
		QueueCapacity: 10000,
		BatchSize:     10000,
	})

	var start = time.Now()
	for j := 0; j < numMessages; j++ {
		if j%100==0{
			fmt.Println(j)
		}
		//log.Println(j)
		err := w.WriteMessages(context.Background(), kafkago.Message{Value: []byte("ing")})
		if err != nil {
			log.Printf("failed to write message: %s", err)
			os.Exit(1)
		}
	}
	w.WriteMessages(context.Background(), kafkago.Message{Value: []byte("FINISH")})
	elapsed := time.Since(start)

	w.Close()

	log.Printf("[kafka-go] msg/s: %f", (float64(numMessages) / elapsed.Seconds()))
}

func main() {

	flag.StringVar(&brokers, "brokers", "localhost:9092", "broker addresses")
	//flag.StringVar(&topic, "topic", "topic", "topic")
	flag.IntVar(&partition, "partition", 0, "partition")
	flag.IntVar(&msgSize, "msgsize", 64, "message size")
	flag.IntVar(&numMessages, "numMessages", 1000000, "number of messages")
	//flag.StringVar(&client, "client", "confluent-kafka-go", "confluent-kafka-go / sarama / kafka-go")
	flag.StringVar(&mode, "mode", "producer", "producer / consumer")
	flag.Parse()
	produceKafkaGo()
	// value = make([]byte, msgSize)
	// rand.Read(value)

	// switch client {

	// case "confluent-kafka-go":
	// 	if mode == "producer" {
	// 		produceConfluentKafkaGo()
	// 	} else {
	// 		consumeConfluentKafkaGo()
	// 	}
	// 	break
	//
	// case "kafka-go":
	// 	if mode == "producer" {
	// 		produceKafkaGo()
	// 	} else {
	// 		log.Printf("not implemented")
	// 		os.Exit(1)
	// 	}
	// 	break

	// default:
	// 	log.Printf("unknown client: %s", client)
	// }
}

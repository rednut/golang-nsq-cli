package main

import (
    "log"
    "github.com/bitly/go-nsq"
    "flag"
    "fmt"
     "github.com/rednut/nsq-cli/lib/nsqdflags"
)

var (
    msg = flag.String("msg", "", "message to publish")
)

func init() {
    flag.StringVar(msg, "m", *msg, "message to publish")
}


func main() {
    flag.Parse()
    nsqdflags.CheckFlags()


    nsqTopic   := nsqdflags.GetTopic()
    nsqAddress := nsqdflags.GetNsqdAddress();


    if "" == *msg {
        log.Fatal("ERROR: missing required 'msg' parameter");
    }



    fmt.Printf("PRODUCER: nsqd=%s, topic=%s, msg=%s\n", nsqAddress, nsqTopic, *msg)


    config := nsq.NewConfig()
    w, _ := nsq.NewProducer(nsqAddress, config)


    err := w.Publish(nsqTopic, []byte(*msg))
    if err != nil {
        log.Fatal("Could not connect ", err)
    }

    w.Stop()

    fmt.Printf("DONE\n\n")
}

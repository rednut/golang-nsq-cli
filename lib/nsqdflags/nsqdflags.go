package nsqdflags
import (
    "flag"
    "log"
    "fmt"
)

const (
    nsqHostDefault = "127.0.0.1"
    nsqPortDefault = 4150
    nsqTopicDefault = "topic_name"
    nsqChannelDefault = "ch"
)
//var (
//    NsqHost = flag.String("host", nsqHostDefault, "nsqf host address")  
//    NsqPort = flag.Int("port", nsqPortDefault, "nsqd port")
//    VerboseMode = flag.Bool("verbose", false, "verbose mode")
 //   NsqTopic = flag.String("topic", "", "topic to publish on")
//)

/*
type NsqdFlags struct {
    Host     string 
    Port     int
    Topic    string
    Verbose  bool
}
*/
var (
  Host string
  Port int
  Topic string
  Channel string
  Verbose bool
)


func init() {
    flag.StringVar(&Host, "host",  nsqHostDefault, "nsqd host / ip address")
    flag.StringVar(&Host, "h", Host, "nsqd host / ip address")


    flag.IntVar(&Port, "port", nsqPortDefault, "nsqd tcp port")
    flag.IntVar(&Port, "p",    Port,           "nsqd tcp port")


    flag.StringVar(&Topic, "topic", nsqTopicDefault, "topic to publish on")
    flag.StringVar(&Topic, "t",      Topic,           "topic to publish on")

    
    flag.StringVar(&Channel, "channel", nsqChannelDefault, "channel to consume on")
    flag.StringVar(&Channel, "c",      Channel,            "channel to consume on")


    flag.BoolVar(&Verbose, "verbose", Verbose, "enable verbose output")
    flag.BoolVar(&Verbose, "v",       Verbose, "enable verbose output")
}


func CheckFlags() error {
    
    if "" == Host {
         log.Fatal("ERROR: missing required 'host' parameter")
    }
    
    if 0 == Port {
         log.Fatal("ERROR: missing required 'port' parameter")
    }


    if "" == Topic {
        log.Fatal("ERROR: missing required 'topic' parameter")
    }

    return nil

}


func GetNsqdAddress() (string) {
    return fmt.Sprintf("%s:%d", Host, Port)
}

func GetTopic () string {

    if "" == Topic {
        log.Fatal("ERROR: missing required 'topic' parameter");
    }
    
    return Topic
}

func GetChannel () (string) {
    
    if "" == Channel {
        log.Fatal("ERROR: missing required 'channel' parameter");
    }

    return Channel
}

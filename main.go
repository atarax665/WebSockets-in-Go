package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"

	uuid "github.com/satori/go.uuid"
)

const (
	PUBLISH     = "publish"
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
)

type PubSub struct {
	Clients       map[string]*Client
	Subscriptions []Subscription
	mu            sync.Mutex
}

type Client struct {
	Id          string
	Connections *websocket.Conn
	mu          sync.Mutex
}

type Message struct {
	Action  string          `json:"action"`
	Topic   string          `json:"topic"`
	Message json.RawMessage `json:"message"`
}

type Subscription struct {
	Topic  string
	Client *Client
}

func NewPubSub() *PubSub {
	return &PubSub{
		Clients: make(map[string]*Client),
	}
}

var ps = NewPubSub()

func AddClient(client *Client) {
	ps.Clients[client.Id] = client
	// send message to client

	payload := []byte("Hello Client ID:" + client.Id)
	client.Send(payload)
}

// RemoveClient removes a client from the PubSub.
func RemoveClient(client *Client) {
	// Remove all subscriptions by this client
	var newSubscriptions []Subscription
	for _, sub := range ps.Subscriptions {
		if sub.Client.Id != client.Id {
			newSubscriptions = append(newSubscriptions, sub)
		}
	}
	ps.Subscriptions = newSubscriptions
	// Remove client from the map
	delete(ps.Clients, client.Id)
	fmt.Println("Client ", client.Id, " is disconnected, total: ", len(ps.Clients))

}

func GetSubscriptions(topic string, client *Client) []Subscription {
	var subscriptionList []Subscription

	for _, subscription := range ps.Subscriptions {
		if client != nil {
			if subscription.Client.Id == client.Id && subscription.Topic == topic {
				subscriptionList = append(subscriptionList, subscription)
			}
		} else {
			if subscription.Topic == topic {
				subscriptionList = append(subscriptionList, subscription)
			}
		}
	}

	return subscriptionList
}

func Subscribe(client *Client, topic string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	clientSubs := GetSubscriptions(topic, client)
	if len(clientSubs) > 0 {
		// Client is already subscribed to this topic
		return
	}

	newSubscription := Subscription{
		Topic:  topic,
		Client: client,
	}
	client.Send([]byte("You are subscribed to topic " + topic))

	ps.Subscriptions = append(ps.Subscriptions, newSubscription)
}

func Publish(topic string, message []byte, excludeClient *Client) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subscriptions := GetSubscriptions(topic, nil)

	for _, sub := range subscriptions {
		if excludeClient == nil || sub.Client.Id != excludeClient.Id {
			fmt.Printf("Sending to client id %s message is %s \n", sub.Client.Id, message)
			sub.Client.Send(message)
		}
	}
	excludeClient.Send([]byte("Message is published to topic " + topic))
}

func (client *Client) Send(message []byte) {
	client.mu.Lock()
	defer client.mu.Unlock()

	websocket.Message.Send(client.Connections, string(message))
}

func Unsubscribe(client *Client, topic string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	var newSubscriptions []Subscription
	for _, sub := range ps.Subscriptions {
		if sub.Client.Id != client.Id || sub.Topic != topic {
			newSubscriptions = append(newSubscriptions, sub)
		}
	}
	ps.Subscriptions = newSubscriptions
	client.Send([]byte("You are unsubscribed from topic " + topic))
}

func HandleReceiveMessage(client *Client, payload string) {
	m := Message{}

	err := json.Unmarshal([]byte(payload), &m)
	if err != nil {
		fmt.Println("This is not a correct message payload")
		client.Send([]byte("This is not a correct message payload"))
		return
	}

	switch m.Action {
	case PUBLISH:
		fmt.Println("This is a publish new message")
		Publish(m.Topic, []byte(m.Message), nil)
	case SUBSCRIBE:
		Subscribe(client, m.Topic)
		fmt.Println("New subscriber to topic", m.Topic, len(ps.Subscriptions), client.Id)
	case UNSUBSCRIBE:
		fmt.Println("Client wants to unsubscribe from the topic", m.Topic, client.Id)
		Unsubscribe(client, m.Topic)
	default:
		break
	}
}

func autoId() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

func websocketHandler(ws *websocket.Conn) {
	client := &Client{
		Id:          autoId(),
		Connections: ws,
	}

	AddClient(client)
	fmt.Println("New Client is connected, total: ", len(ps.Clients))

	for {
		var payload string
		err := websocket.Message.Receive(ws, &payload)
		if err != nil {
			log.Println("Something went wrong", err)
			RemoveClient(client)
			return
		}

		HandleReceiveMessage(client, payload)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/ws", websocket.Handler(websocketHandler))

	fmt.Println("Server is running: http://localhost:3000/ws")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

package ws

type WebsocketHub struct {
	clients    map[string][]*WebsocketHubClient
	messages   chan WebsocketMessage
	register   chan *WebsocketHubClient
	unregister chan *WebsocketHubClient
}

type WebsocketMessage struct {
	Broadcast  bool
	ToIdentity string
	Message    []byte
}

func NewWebsocketHub() *WebsocketHub {
	return &WebsocketHub{
		clients:    make(map[string][]*WebsocketHubClient),
		messages:   make(chan WebsocketMessage),
		register:   make(chan *WebsocketHubClient),
		unregister: make(chan *WebsocketHubClient),
	}
}

func (wsh *WebsocketHub) Run() {
	for {
		select {
		case client := <-wsh.register:
			wsh.addClient(client)
		case client := <-wsh.unregister:
			wsh.removeClient(client)
		case message := <-wsh.messages:
			if message.Broadcast {
				wsh.broadcastMessage(message.Message)
			} else {
				wsh.sendMessage(message.ToIdentity, message.Message)
			}
		}
	}
}

func (wsh *WebsocketHub) Send(message WebsocketMessage) {
	wsh.messages <- message
}

func (wsh *WebsocketHub) Register(client *WebsocketHubClient) {
	wsh.register <- client
}

func (wsh *WebsocketHub) Unregister(client *WebsocketHubClient) {
	wsh.unregister <- client
}

func (wsh *WebsocketHub) addClient(client *WebsocketHubClient) {
	identity := client.GetIdentity()

	if _, ok := wsh.clients[identity]; !ok {
		wsh.clients[identity] = []*WebsocketHubClient{}
	}

	for _, tmp := range wsh.clients[identity] {
		if tmp == client {
			return
		}
	}

	wsh.clients[identity] = append(wsh.clients[identity], client)
}

func (wsh *WebsocketHub) removeClient(client *WebsocketHubClient) {
	identity := client.GetIdentity()

	if clients, ok := wsh.clients[identity]; ok {
		for index, tmp := range clients {
			if tmp == client {
				length := len(clients)
				clients[length-1], clients[index] = clients[index], clients[length-1]
				wsh.clients[identity] = clients[:length-1]
				return
			}
		}
	}
}

func (wsh *WebsocketHub) sendMessage(identity string, message []byte) {
	clients, ok := wsh.clients[identity]
	if !ok {
		return
	}

	for _, client := range clients {
		client.Send(message)
	}
}

func (wsh *WebsocketHub) broadcastMessage(message []byte) {
	for identity := range wsh.clients {
		wsh.sendMessage(identity, message)
	}
}

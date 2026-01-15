package utils


func removeClient(clients *[]Client, clientInfo Client) {
	for i, c := range *clients {
		if c.name == clientInfo.name {
			*clients = append((*clients)[:i], (*clients)[i+1:]...)
			return
		}
	}
}

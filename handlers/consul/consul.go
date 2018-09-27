package consul

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

type Inventory struct {
	Nodes struct {
		Hosts []string `json:"hosts"`
		Vars  struct {
			Consul string `json:"consul"`
		} `json:"vars"`
	} `json:"replaceme"`
	Meta struct {
		Hostvars struct {
		} `json:"hostvars"`
	} `json:"_meta"`
}

func checkServers(server string) bool {
	var allowed = false
	// We allow localhost by default
	if server == "127.0.0.1" || server == "localhost" {
		allowed = true
	}
	s := strings.Split(os.Getenv("ALLOWED_SERVERS"), ",")
	for _, b := range s {
		if b == server {
			allowed = true
		}
	}

	return (allowed)

}

/*** GenInventory does a query on consul ***/
func GenInventory(c *gin.Context) {
	consulServer := c.Param("server")
	inventoryName := c.Param("inventoryname")
	if checkServers(consulServer) {
		client, err := api.NewClient(&api.Config{Address: fmt.Sprintf("%s:8500", consulServer)})
		if err != nil {
			panic(err)
		}

		cg := client.Health()
		service, _, err := cg.Service("consul", "", true, &api.QueryOptions{AllowStale: false})
		if err != nil {
			panic(err)
		}

		b := Inventory{}

		for _, n := range service {
			b.Nodes.Hosts = append(b.Nodes.Hosts, n.Node.Address)
		}

		m, _ := json.Marshal(b)
		var a interface{}
		json.Unmarshal(m, &a)
		z := a.(map[string]interface{})
		z[inventoryName] = z["replaceme"]
		delete(z, "replaceme")
		c.JSON(200, z)
	} else {
		c.JSON(403, gin.H{
			"message": "Not allowed access",
		})
	}
}

/*** GenNodes does a query on consul ***/
func GenNodes(c *gin.Context) {
	consulServer := c.Param("server")
	inventoryName := c.Param("inventoryname")
	if checkServers(consulServer) {
		client, err := api.NewClient(api.DefaultConfig())
		if err != nil {
			panic(err)
		}

		cg := client.Catalog()
		ch := client.Health()
		nodes, _, err := cg.Nodes(&api.QueryOptions{AllowStale: false})
		if err != nil {
			panic(err)
		}

		b := Inventory{}
		b.Nodes.Vars.Consul = "pong"
		for _, n := range nodes {
			hc, _, err := ch.Node(n.Node, &api.QueryOptions{})
			if err != nil {
				panic(err)
			}

			for _, h := range hc {
				if h.Name == "Serf Health Status" && h.Status == "passing" {
					b.Nodes.Hosts = append(b.Nodes.Hosts, n.Address)
				}
			}

		}

		m, _ := json.Marshal(b)
		var a interface{}
		json.Unmarshal(m, &a)
		z := a.(map[string]interface{})
		z[inventoryName] = z["replaceme"]
		delete(z, "replaceme")
		c.JSON(200, z)
	} else {
		c.JSON(403, gin.H{
			"message": "Not allowed access",
		})
	}
}

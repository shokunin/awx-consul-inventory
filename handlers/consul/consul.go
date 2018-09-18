package consul

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

func checkServers(server string) bool {
	var allowed = false
	s := strings.Split(os.Getenv("ALLOWED_SERVERS"), ",")
	for _, b := range s {
		if b == server {
			allowed = true
		}
	}

	return (allowed)

}

func GenInventory(c *gin.Context) {
	consulServer := c.Param("server")
	if checkServers(consulServer) {
		fmt.Println(consulServer)
		client, err := api.NewClient(api.DefaultConfig())
		if err != nil {
			panic(err)
		}

		cg := client.Health()
		service, _, err := cg.Service("consul", "", true, &api.QueryOptions{AllowStale: false})
		if err != nil {
			panic(err)
		}

		for _, n := range service {
			fmt.Println(n.Node.Address)
		}

		c.JSON(200, gin.H{
			"message": "OK",
		})
	} else {
		c.JSON(403, gin.H{
			"message": "Not allowed access",
		})
	}
}

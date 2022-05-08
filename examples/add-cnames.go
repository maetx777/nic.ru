package examples

import (
	"fmt"
	api "github.com/maetx777/nic.ru-golang-client/client"
)

func AddCnames() {
	config := &api.Config{
		Credentials: &api.Credentials{
			OAuth2: &api.OAuth2Creds{
				ClientID: "qwe",
				SecretID: "xyz",
			},
			Username: "foo",
			Password: "bar",
		},
		ZoneName:       "example.com",
		DnsServiceName: "EXAMPLE",
		CachePath:      "/tmp/.nic.ru.token",
	}
	client := api.NewClient(config)
	var cnames = []string{
		`foo`,
		`bar`,
	}
	if response, err := client.AddCnames(cnames, `habr.com.`, 600); err != nil {
		fmt.Printf(`Add record error: %s`, err.Error())
		return
	} else {
		for _, record := range response.Data.Zone[0].Rr {
			fmt.Printf(`Added record: %s CNAME %s`, record.Name, record.Cname.Name)
		}
	}
	if _, err := client.CommitZone(); err != nil {
		fmt.Printf(`Commit error: %s`, err.Error())
	} else {
		fmt.Printf(`Zone committed`)
	}
}
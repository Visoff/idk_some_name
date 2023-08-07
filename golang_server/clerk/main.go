package clerk

import "github.com/clerkinc/clerk-sdk-go/clerk"

var Client clerk.Client

func Init() error {
	client, err := clerk.NewClient("sk_test_jH96ZtpF40tew2fbYT5vC6cSt90h4yljZjkdVobGIC")
	if err != nil {
		return err
	}
	Client = client
	return nil
}

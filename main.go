package main

import (
	"fmt"
)

type ResolverRequest struct {
	Resolver string
	Context  interface{}
}
type FeatureQueryInput struct {
	AppID       string
	Environment string
}
type UpdateAppInput struct {
	AppID    string
	AppName  string
	UpdateBy string
}

func handler(request ResolverRequest) interface{} {
	var response interface{}
	switch request.Resolver {
	case "query.getApplication":
		if appID, ok := request.Context.(string); ok {
			response = getApplication(appID)
		}
	case "query.listFeatures":
		if input, ok := request.Context.(FeatureQueryInput); ok {
			response = listFeatures(input)
		}
	case "mutation.updateApplication":
		if input, ok := request.Context.(UpdateAppInput); ok {
			response = updateApplication(input)
		}
	default:
		panic("Unhandler resolver")
	}
	return response
}
func main() {
	fmt.Println(handler(ResolverRequest{"query.getApplication", "facebook"}))
	fmt.Println(handler(ResolverRequest{"query.listFeatures",
		FeatureQueryInput{"facebook", "production"}}))
	fmt.Println(handler(ResolverRequest{"mutation.updateApplication",
		UpdateAppInput{"facebook", "Facebook 2019", "admin"}}))

	// Output:
	// querying application :facebook
	// list feature belong to app:facebook with environment:production
	// user:admin update application:facebook to name:Facebook 2019
}
func getApplication(appID string) string {
	return fmt.Sprintf("querying application :%s", appID)
}
func listFeatures(input FeatureQueryInput) string {
	return fmt.Sprintf("list feature belong to app:%s with environment:%s", input.AppID, input.Environment)
}
func updateApplication(input UpdateAppInput) string {
	return fmt.Sprintf("user:%s update application:%s to name:%s", input.UpdateBy, input.AppID, input.AppName)
}

package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

func main() {
	owner := "alidevhere"
	repo := "go-test-api"
	actionsCache := "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"
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: actionsCache})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// List action Cache
	list, _, _ := client.Actions.ListCaches(context.Background(), owner, repo, &github.ActionsCacheListOptions{})
	for _, s := range list.ActionsCaches {
		fmt.Printf("ID: %d, Key: %s, LastAccess: %s, Ref: %v, Created: %s, Size: %d, Version: %s\n",
			*s.ID, *s.Key, s.LastAccessedAt.String(), *s.Ref, s.CreatedAt.String(), *s.SizeInBytes, *s.Version)
	}

	//Deleting by cache Key
	// resp, err := client.Actions.DeleteCacheByKey(ctx, owner, repo, *list.ActionsCaches[0].Key, list.ActionsCaches[0].Ref)
	// print("Status: ", resp.StatusCode)
	// print(resp.Request.URL.RequestURI())
	// if err != nil {
	// 	print(err.Error())
	// }

	//===//Deleting by ID
	resp, err := client.Actions.DeleteCachesByID(ctx, owner, repo, *list.ActionsCaches[0].ID)
	print("Status: ", resp.StatusCode)
	print(resp.Request.URL.RequestURI())
	if err != nil {
		print(err.Error())
	}

	// Get Action Cache Usage
	// usage, resp, err := client.Actions.GetActionsCacheUsage(ctx, owner, repo)
	// println("Status: ", resp.StatusCode)
	// println(resp.Request.URL.RequestURI())
	// if err != nil {
	// 	print(err.Error())
	// }

	// if usage != nil {
	// 	println(*usage.ActiveCacheUsageSize)
	// 	println(*usage.ActiveCachesCount)
	// 	println(*usage.FullName)
	// }
	// org, resp, err := client.Actions.GetActionsCacheUsageForOrg(ctx, "ali-dev-org")
	// if err != nil {
	// 	println(err.Error())
	// }
	// println(resp.StatusCode)
	// println(*org.TotalCount)
	// println(*org.RepoCacheUsage[0].ActiveCacheUsageSize)
	// println(*org.RepoCacheUsage[0].ActiveCachesCount)
	// println(*org.RepoCacheUsage[0].FullName)
	// /println(*org.RepoCacheUsage[0])

	//ENTERPRISE
	// ent, resp, err := client.Actions.GetActionsCacheUsageForEnterprise(ctx, "ali-dev-org")
	// println(resp.Status)
	// if err != nil {
	// 	print(err.Error())
	// 	return
	// }
	// print(*ent.TotalActiveCacheUsageSize)
	// print(*ent.TotalActiveCachesCount)

	// l, resp, err := client.Organizations.ListAll(ctx, "", nil,&github.OrganizationsListOptions{})
	// if err != nil {
	// 	println(err.Error())
	// }
	// println(resp.Status)
	// for _, i := range l {
	// 	println(*i.ID)
	// 	println(i.GetCompany())
	// 	println(i.GetType())
	// 	println(i.String())

	// 	// println(*i.PublicRepos)

	// }

}

// github.Organization{Login:"ali-dev-org", ID:120642991, NodeID:"O_kgDOBzDdrw",
//AvatarURL:"https://avatars.githubusercontent.com/u/120642991?v=4",
// URL:"https://api.github.com/orgs/ali-dev-org", EventsURL:"https://api.github.com/orgs/ali-dev-org/events",
//HooksURL:"https://api.github.com/orgs/ali-dev-org/hooks", IssuesURL:"https://api.github.com/orgs/ali-dev-org/issues",
// MembersURL:"https://api.github.com/orgs/ali-dev-org/members{/member}",
// PublicMembersURL:"https://api.github.com/orgs/ali-dev-org/public_members{/member}",
//ReposURL:"https://api.github.com/orgs/ali-dev-org/repos"}
// 120659047

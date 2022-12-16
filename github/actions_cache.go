package github

import (
	"context"
	"fmt"
)

// ActionsCache represents a GitHub action cache.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/cache?apiVersion=2022-11-28#about-the-cache-api
type ActionsCache struct {
	ID             *int64     `json:"id,omitempty" url:"-"`
	Ref            *string    `json:"ref,omitempty" url:"ref"`
	Key            *string    `json:"key,omitempty" url:"key"`
	Version        *string    `json:"version,omitempty" url:"-"`
	LastAccessedAt *Timestamp `json:"last_accessed_at,omitempty" url:"-"`
	CreatedAt      *Timestamp `json:"created_at,omitempty" url:"-"`
	SizeInBytes    *int64     `json:"size_in_bytes,omitempty" url:"-"`
}

// ActionsCacheList represents a list of GitHub actions Cache.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/cache?apiVersion=2022-11-28#list-github-actions-caches-for-a-repository
type ActionsCacheList struct {
	TotalCount    *int64          `json:"total_count,omitempty"`
	ActionsCaches []*ActionsCache `json:"actions_caches,omitempty"`
}

// ActionsCacheUsage represents a GitHub Actions Cache Usage object.
//
//GitHub API docs: https://docs.github.com/en/rest/actions/cache?apiVersion=2022-11-28#get-github-actions-cache-usage-for-a-repository
type ActionsCacheUsage struct {
	FullName             *string `json:"full_name"`
	ActiveCacheUsageSize *int64  `json:"active_caches_size_in_bytes"`
	ActiveCachesCount    *int64  `json:"active_caches_count"`
}

// OrgRepoCacheUsageList represents a list repositories with GitHub Actions cache usage for an organization.
//
//GitHub API docs: https://docs.github.com/en/rest/actions/cache?apiVersion=2022-11-28#get-github-actions-cache-usage-for-a-repository
type OrgRepoCacheUsageList struct {
	TotalCount     *int64               `json:"total_count"`
	RepoCacheUsage []*ActionsCacheUsage `json:"repository_cache_usages"`
}

type CacheUsage struct {
	TotalActiveCacheUsageSize *int64 `json:"total_active_caches_size_in_bytes"`
	TotalActiveCachesCount    *int64 `json:"total_active_caches_count"`
}

type ActionsCacheListOptions struct {
	//The number of results per page (max 100). Default:30
	PerPage int `url:"per_page,omitempty"`
	//Page number of the results to fetch.Default:1
	Page int `url:"page,omitempty"`
	//The Git reference for the results you want to list.
	// The ref for a branch can be formatted either as refs/heads/<branch name>
	//or simply <branch name>. To reference a pull request use refs/pull/<number>/merge
	Ref string `url:"ref,omitempty"`
	Key string `url:"key,omitempty"`
	//Can be one of: "created_at", "last_accessed_at", "size_in_bytes".
	//Default: "last_accessed_at"
	Sort string `url:"sort,omitempty"`
	//Can be one of: "asc", "desc"
	//Default: desc
	Direction string `url:"direction,omitempty"`
}

// Lists the GitHub Actions caches for a repository.
// You must authenticate using an access token with the repo scope to use this endpoint.
//
// Permissions: must have the actions:read permission to use this endpoint.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/cache?apiVersion=2022-11-28#list-github-actions-caches-for-a-repository
func (s *ActionsService) ListCaches(ctx context.Context, owner, repo string, opts *ActionsCacheListOptions) (*ActionsCacheList, *Response, error) {

	u := fmt.Sprintf("repos/%v/%v/actions/caches", owner, repo)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	actionCacheList := new(ActionsCacheList)
	resp, err := s.client.Do(ctx, req, actionCacheList)
	if err != nil {
		return nil, resp, err
	}

	return actionCacheList, resp, nil
}

// Deletes one or more GitHub Actions caches for a repository, using a complete cache key.
// By default, all caches that match the provided key are deleted, but you can optionally provide
// a Git ref to restrict deletions to caches that match both the provided key and the Git ref.
// The ref for a branch can be formatted either as "refs/heads/<branch name>" or simply "<branch name>".
// To reference a pull request use "refs/pull/<number>/merge". If you don't want to use ref just pass nil in parameter.
//
// Permissions: You must authenticate using an access token with the repo scope to use this endpoint. GitHub Apps must have the actions:write permission to use this endpoint.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/cache?apiVersion=2022-11-28#delete-github-actions-caches-for-a-repository-using-a-cache-key
func (s *ActionsService) DeleteCachesByKey(ctx context.Context, owner, repo, key string, ref *string) (*Response, error) {
	u := fmt.Sprintf("repos/%v/%v/actions/caches", owner, repo)
	u, err := addOptions(u, ActionsCache{Key: &key, Ref: ref})
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Deletes a GitHub Actions cache for a repository, using a cache ID.
//
// Permissions: You must authenticate using an access token with the repo scope to use this endpoint. GitHub Apps must have the actions:write permission to use this endpoint.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/cache?apiVersion=2022-11-28#delete-a-github-actions-cache-for-a-repository-using-a-cache-id
func (s *ActionsService) DeleteCachesByID(ctx context.Context, owner, repo string, cacheId int64) (*Response, error) {
	u := fmt.Sprintf("/repos/%v/%v/actions/caches/%v", owner, repo, cacheId)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Gets GitHub Actions cache usage for a repository. The data fetched using this API is refreshed approximately every 5 minutes,
// so values returned from this endpoint may take at least 5 minutes to get updated.
//
// Permissions: Anyone with read access to the repository can use this endpoint. If the repository is private, you must use an
// access token with the repo scope. GitHub Apps must have the actions:read permission to use this endpoint.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/cache?apiVersion=2022-11-28#get-github-actions-cache-usage-for-a-repository
func (s *ActionsService) GetCacheUsageForRepo(ctx context.Context, owner, repo string) (*ActionsCacheUsage, *Response, error) {
	u := fmt.Sprintf("/repos/%v/%v/actions/cache/usage", owner, repo)

	cacheUsage := new(ActionsCacheUsage)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return cacheUsage, nil, err
	}
	res, err := s.client.Do(ctx, req, cacheUsage)
	return cacheUsage, res, err
}

// Lists repositories and their GitHub Actions cache usage for an organization. The data fetched using this API is
// refreshed approximately every 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.
//
// Permissions: You must authenticate using an access token with the read:org scope to use this endpoint.
// GitHub Apps must have the organization_admistration:read permission to use this endpoint.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/cache?apiVersion=2022-11-28#list-repositories-with-github-actions-cache-usage-for-an-organization
func (s *ActionsService) ListCacheUsageByRepoForOrg(ctx context.Context, org string, opts *ListOptions) (*OrgRepoCacheUsageList, *Response, error) {
	u := fmt.Sprintf("/orgs/%v/actions/cache/usage-by-repository", org)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	cacheUsage := new(OrgRepoCacheUsageList)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return cacheUsage, nil, err
	}
	res, err := s.client.Do(ctx, req, cacheUsage)
	return cacheUsage, res, err
}

// Gets the total GitHub Actions cache usage for an organization. The data fetched using this API is refreshed approximately every
// 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.
//
// Permissions: You must authenticate using an access token with the read:org scope to use this endpoint.
// GitHub Apps must have the organization_admistration:read permission to use this endpoint.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/cache?apiVersion=2022-11-28#get-github-actions-cache-usage-for-an-organization
func (s *ActionsService) GetCacheUsageForOrg(ctx context.Context, org string) (*CacheUsage, *Response, error) {
	u := fmt.Sprintf("/orgs/%v/actions/cache/usage", org)

	cacheUsage := new(CacheUsage)
	req, err := s.client.NewRequest("GET", u, cacheUsage)

	if err != nil {
		return cacheUsage, nil, err
	}
	res, err := s.client.Do(ctx, req, cacheUsage)
	return cacheUsage, res, err
}

// Gets the total GitHub Actions cache usage for an enterprise. The data fetched using this API is refreshed approximately every 5 minutes,
// so values returned from this endpoint may take at least 5 minutes to get updated.
//
// Permissions: You must authenticate using an access token with the "admin:enterprise" scope to use this endpoint.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/cache?apiVersion=2022-11-28#get-github-actions-cache-usage-for-an-enterprise
func (s *ActionsService) GetCacheUsageForEnterprise(ctx context.Context, enterprise string) (*CacheUsage, *Response, error) {
	u := fmt.Sprintf("/enterprises/%v/actions/cache/usage", enterprise)

	cacheUsage := new(CacheUsage)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return cacheUsage, nil, err
	}
	res, err := s.client.Do(ctx, req, cacheUsage)
	return cacheUsage, res, err
}

package regionservice

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/fegoa89/ipapi"
)

type RegionService struct {
	client *memcache.Client
	// need to add ipapi module as dependecy
}

func NewRegionService(ips ...string) *RegionService {
	mc := memcache.New(ips...)
	return &RegionService{
		client: mc,
	}
}

// TODO:: need to mock dependencies and write tests here
// the most part of this small project that need to be tested it this function
func (r *RegionService) GetRegionBasedOnIp(ip string) (string, error) {
	item, err := r.client.Get(ip)
	if err != nil && err != memcache.ErrCacheMiss {
		return "", err
	}

	// TODO::divide memcache module and ipapi module
	// in order to mock them and test the regionservice
	if err == memcache.ErrCacheMiss {
		resp, err := ipapi.FindLocation(ip)
		if err != nil {
			return "", err
		}
		err = r.client.Set(&memcache.Item{
			Key:   ip,
			Value: []byte(resp.Country),
		})
		if err != nil {
			return "", err
		}

		return resp.Country, nil
	}

	return string(item.Value), nil
}

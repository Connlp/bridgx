package alibaba

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/galaxy-future/BridgX/pkg/cloud"
)

func (p *AlibabaCloud) CreateServer(region, loadBalancerId string, servers []cloud.Server) error {
	request := slb.CreateAddBackendServersRequest()
	request.LoadBalancerId = loadBalancerId
	request.RegionId = region
	serversJson, err := json.Marshal(servers)
	if err != nil {
		return err
	}
	request.BackendServers = string(serversJson)

	_, err = p.slbClient.AddBackendServers(request)
	return err
}

func (p *AlibabaCloud) RemoveServer(region, loadBalancerId string, servers []cloud.Server) (err error) {
	chunkServers := ServersChunk(servers, _batchSize)
	for _, batchServers := range chunkServers {
		serversJson, err := json.Marshal(batchServers)
		if err != nil {
			return err
		}
		request := slb.CreateRemoveBackendServersRequest()
		request.LoadBalancerId = loadBalancerId
		request.RegionId = region
		request.BackendServers = string(serversJson)
		_, err = p.slbClient.RemoveBackendServers(request)
		if err != nil {
			return err
		}
	}
	return err
}
func ServersChunk(slice []cloud.Server, size int) (chunkServers [][]cloud.Server) {
	if size >= len(slice) {
		chunkServers = append(chunkServers, slice)
		return
	}
	end := size
	for i := 0; i <= (len(slice) - size); i += size {
		chunkServers = append(chunkServers, slice[i:end])
		end += size
	}
	return
}

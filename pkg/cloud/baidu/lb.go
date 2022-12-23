package baidu

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/services/blb"
	"github.com/galaxy-future/BridgX/pkg/cloud"
)

func (b *BaiduCloud) CreateServer(region, loadBalancerId string, servers []cloud.Server) error {
	bsms := []blb.BackendServerModel{}
	for _, v := range servers {
		w, _ := strconv.Atoi(v.Weight)
		bsm := blb.BackendServerModel{
			InstanceId: v.ServerId,
			Weight:     w,
		}
		bsms = append(bsms, bsm)
	}
	args := &blb.AddBackendServersArgs{
		// 配置后端服务器的列表及权重
		BackendServerList: bsms,
	}
	err := b.blbClient.AddBackendServers(loadBalancerId, args)
	if err != nil {
		return err
	}
	return nil
}

func (b *BaiduCloud) RemoveServer(region, loadBalancerId string, servers []cloud.Server) (err error) {
	ins := []string{}
	for _, v := range servers {
		ins = append(ins, v.ServerId)
	}
	args := &blb.RemoveBackendServersArgs{
		// 要从后端服务器列表中释放的实例列表
		BackendServerList: ins,
	}
	err = b.blbClient.RemoveBackendServers(loadBalancerId, args)
	if err != nil {
		return err
	}
	return
}

package tencent

import (
	"errors"
	"fmt"

	"github.com/galaxy-future/BridgX/pkg/cloud"
	clb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
)

func (p *TencentCloud) CreateLoadBalancer(req cloud.CreateLoadBalancerRequest) (cloud.CreateLoadBalancerResponse, error) {
	var (
		addressTypeInternet = "INTERNAL"
	)
	request := clb.NewCreateLoadBalancerRequest()
	request.LoadBalancerName = &req.LoadBalancerName
	request.LoadBalancerType = &addressTypeInternet

	response, err := p.clbClient.CreateLoadBalancer(request)
	if err != nil {
		return cloud.CreateLoadBalancerResponse{}, err
	}
	if len(response.Response.LoadBalancerIds) == 0 {
		if response.Response.DealName != nil {
			// TODO:
			// 存在某些场景，如创建出现延迟时，此字段可能返回为空；
			// 此时可以根据接口返回的RequestId或DealName参数，通过DescribeTaskStatus接口查询创建的资源ID。
		}
		return cloud.CreateLoadBalancerResponse{},
			fmt.Errorf("no load balancer created, requestId: %s", *response.Response.RequestId)
	}
	return cloud.CreateLoadBalancerResponse{
		LoadBalancerId: *response.Response.LoadBalancerIds[0],
	}, nil
}

func (p *TencentCloud) CreateListener(req cloud.CreateListenerRequest) error {
	if req.Protocol == "" {
		return errors.New("protocol empty")
	}
	var protocol = string(req.Protocol)
	if len(req.PortList) == 0 {
		return errors.New("port empty")
	}

	request := clb.NewCreateListenerRequest()
	request.LoadBalancerId = &req.LoadBalancerId
	request.Protocol = &protocol

	for _, port := range req.PortList {
		var temPort = int64(port)
		request.Ports = append(request.Ports, &temPort)
	}

	response, err := p.clbClient.CreateListener(request)
	if err != nil {
		return err
	}

	if len(response.Response.ListenerIds) == 0 {
		return fmt.Errorf("no listener created, requestId: %s", *response.Response.RequestId)
	}

	return nil
}

func (p *TencentCloud) RegisterBackendServer(req cloud.RegisterBackendServerRequest) error {
	if len(req.BackendServerList) == 0 {
		return errors.New("target backend server empty")
	}

	request := clb.NewRegisterTargetsRequest()
	request.LoadBalancerId = &req.LoadBalancerId
	request.ListenerId = &req.ListenerId

	for _, server := range req.BackendServerList {
		var (
			temPort   = int64(server.Port)
			temWeight = int64(server.Weight)
		)

		request.Targets = append(request.Targets, &clb.Target{
			Port:       &temPort,
			InstanceId: &server.ServerId,
			Weight:     &temWeight,
		})
	}

	_, err := p.clbClient.RegisterTargets(request)
	return err
}

func (p *TencentCloud) DeregisterBackendServer(req cloud.DeregisterBackendServerRequest) error {
	if len(req.BackendServerList) == 0 {
		return errors.New("target backend server empty")
	}

	request := clb.NewDeregisterTargetsRequest()
	request.LoadBalancerId = &req.LoadBalancerId
	request.ListenerId = &req.ListenerId

	for _, server := range req.BackendServerList {
		var (
			temPort   = int64(server.Port)
			temWeight = int64(server.Weight)
		)

		request.Targets = append(request.Targets, &clb.Target{
			Port:       &temPort,
			InstanceId: &server.ServerId,
			Weight:     &temWeight,
		})
	}

	_, err := p.clbClient.DeregisterTargets(request)
	return err
}

func (p *TencentCloud) UpdateBackendServer(req cloud.UpdateBackendServerRequest) error {
	if len(req.BackendServerList) == 0 {
		return errors.New("target backend server empty")
	}

	request := clb.NewBatchModifyTargetWeightRequest()
	request.LoadBalancerId = &req.LoadBalancerId

	for _, server := range req.BackendServerList {
		if server.ServerId == "" {
			return errors.New("ServerId empty")
		}

		var (
			temWeight = int64(server.Weight)
			temPort   = int64(server.Port)
			temTarget = []*clb.Target{
				&clb.Target{
					Port: &temPort,
				},
			}
		)

		request.ModifyList = append(request.ModifyList, &clb.RsWeightRule{
			ListenerId: &server.ServerId,
			Targets:    temTarget,
			Weight:     &temWeight,
		})
	}

	_, err := p.clbClient.BatchModifyTargetWeight(request)
	return err
}

func (p *TencentCloud) StartLoadBalancerListener(req cloud.StartLoadBalancerListenerRequest) error {
	return errors.New("do not use this api")
}

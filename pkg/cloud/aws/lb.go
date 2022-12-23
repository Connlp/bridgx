package aws

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/spf13/cast"
)

func (p *AWSCloud) CreateServer(region, loadBalancerId string, servers []cloud.Server) (err error) {
	targets := make([]*elbv2.TargetDescription, 0)
	for _, v := range servers {
		addr := strings.Split(v.ServerIp, ":")
		if len(addr) != 2 {
			return fmt.Errorf("invalid ServerIp %v", v.ServerIp)
		}

		targets = append(targets, &elbv2.TargetDescription{
			Id:   aws.String(v.ServerId),
			Port: aws.Int64(cast.ToInt64(addr[1])),
		})
	}
	input := &elbv2.RegisterTargetsInput{
		TargetGroupArn: aws.String(loadBalancerId),
		Targets:        targets,
	}

	_, err = p.elbClient.RegisterTargets(input)
	if err != nil {
		return err
	}
	return nil
}

func (p *AWSCloud) RemoveServer(region, loadBalancerId string, servers []cloud.Server) (err error) {
	targets := make([]*elbv2.TargetDescription, 0)
	for _, v := range servers {
		targets = append(targets, &elbv2.TargetDescription{
			Id: aws.String(v.ServerId),
		})
	}
	input := &elbv2.DeregisterTargetsInput{
		TargetGroupArn: aws.String(loadBalancerId),
		Targets:        targets,
	}

	_, err = p.elbClient.DeregisterTargets(input)
	if err != nil {
		return err
	}
	return nil
}

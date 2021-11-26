package helper

import (
	"time"

	"github.com/galaxy-future/BridgX/cmd/api/response"
	"github.com/galaxy-future/BridgX/internal/model"
	"github.com/galaxy-future/BridgX/internal/service"
	"github.com/spf13/cast"
)

func ConvertToClusterThumbList(clusters []model.Cluster, countMap map[string]int64) []response.ClusterThumb {
	res := make([]response.ClusterThumb, 0)
	for _, cluster := range clusters {
		c := response.ClusterThumb{
			ClusterId:     cast.ToString(cluster.Id),
			ClusterName:   cluster.ClusterName,
			InstanceCount: countMap[cluster.ClusterName],
			Provider:      cluster.Provider,
			Account:       cluster.AccountKey,
			CreateAt:      cluster.CreateAt.String(),
			CreateBy:      cluster.CreateBy,
		}
		res = append(res, c)
	}
	return res
}

func ConvertToInstanceStat(instanceType service.InstanceTypeByZone, count int64) response.InstanceStatResponse {
	return response.InstanceStatResponse{
		InstanceTypeDesc: instanceType.GetDesc(),
		InstanceCount:    count,
	}
}

func ConvertToTaskThumbList(tasks []model.Task) []response.TaskThumb {
	res := make([]response.TaskThumb, 0)
	for _, task := range tasks {
		finishTime := time.Now()
		if task.FinishTime != nil {
			finishTime = *task.FinishTime
		}
		t := response.TaskThumb{
			TaskId:      cast.ToString(task.Id),
			TaskName:    cast.ToString(task.TaskName),
			TaskAction:  task.TaskAction,
			Status:      task.Status,
			ClusterName: task.TaskFilter,
			CreateAt:    getStringTime(task.CreateAt),
			ExecuteTime: int(finishTime.Sub(*task.CreateAt).Seconds()),
			FinishAt:    getStringTime(task.FinishTime),
		}
		res = append(res, t)
	}
	return res
}

func getStringTime(time *time.Time) string {
	if time == nil {
		return ""
	}
	return time.String()
}

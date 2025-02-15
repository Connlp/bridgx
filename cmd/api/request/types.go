package request

import (
	"github.com/galaxy-future/BridgX/internal/model"
	"github.com/galaxy-future/BridgX/internal/service"
	"github.com/galaxy-future/BridgX/internal/types"
)

type TagRequest struct {
	ClusterName string            `json:"cluster_name" binding:"required,min=1"`
	Tags        map[string]string `json:"tags" binding:"required"`
	ValueDesc   map[string]string `json:"value_desc"`
}

type GetTagsRequest struct {
	ClusterName string `json:"cluster_name" binding:"required,min=1" form:"cluster_name"`
	TagKey      string `json:"tag_key" form:"tag_key"`
	PageNumber  int    `json:"page_number" form:"page_number"`
	PageSize    int    `json:"page_size" form:"page_size"`
}

type SetExpectInstanceCountRequest struct {
	ClusterName string `json:"cluster_name"`
	ExpectCount int    `json:"expect_count"`
}

type ExpandClusterRequest struct {
	TaskName    string `json:"task_name"`
	ClusterName string `json:"cluster_name" binding:"required"`
	Count       int    `json:"count" binding:"required,min=1,max=10000"`
}

type ShrinkClusterRequest struct {
	TaskName    string   `json:"task_name"`
	ClusterName string   `json:"cluster_name" binding:"required"`
	IPs         []string `json:"ips"`
	Count       int      `json:"count" binding:"required,min=1,max=10000"`
}

type ShrinkAllInstancesRequest struct {
	TaskName    string `json:"task_name"`
	ClusterName string `json:"cluster_name" binding:"required"`
}

type CreateVpcRequest struct {
	Provider  string `json:"provider"`
	RegionId  string `json:"region_id"`
	VpcName   string `json:"vpc_name"`
	CidrBlock string `json:"cidr_block"`
	AK        string `json:"ak"`
}

func (c *CreateVpcRequest) Check() bool {
	return c.Provider != "" && c.RegionId != "" && c.VpcName != "" && c.AK != ""
}

type DescribeVpcRequest struct {
	Provider   string `form:"provider" binding:"required"`
	RegionId   string `form:"region_id" binding:"required"`
	VpcName    string `form:"vpc_name"`
	AccountKey string `form:"account_key" binding:"required"`
}

type CreateSwitchRequest struct {
	SwitchName string `json:"switch_name"`
	RegionId   string `json:"region_id"`
	VpcId      string `json:"vpc_id"`
	CidrBlock  string `json:"cidr_block"`
	GatewayIp  string `json:"gateway_ip"`
	ZoneId     string `json:"zone_id"`
	AK         string `json:"account_key" binding:"required"`
}

func (c *CreateSwitchRequest) Check() bool {
	return c.SwitchName != "" && c.VpcId != "" && c.CidrBlock != "" && c.ZoneId != ""
}

type CreateSecurityGroupRequest struct {
	AK                string `json:"account_key" binding:"required"`
	VpcId             string `json:"vpc_id"`
	RegionId          string `json:"region_id"`
	SecurityGroupName string `json:"security_group_name"`
	SecurityGroupType string `json:"security_group_type"`
}

func (c *CreateSecurityGroupRequest) Check() bool {
	return c.SecurityGroupName != "" && c.RegionId != "" && c.VpcId != ""
}

type AddSecurityGroupRuleRequest struct {
	AK              string              `json:"account_key" binding:"required"`
	VpcId           string              `json:"vpc_id"`
	RegionId        string              `json:"region_id"`
	SecurityGroupId string              `json:"security_group_id"`
	Rules           []service.GroupRule `json:"rules"`
}

func (c *AddSecurityGroupRuleRequest) Check() bool {
	return c.VpcId != "" && c.RegionId != "" && c.SecurityGroupId != "" &&
		len(c.Rules) > 0
}

type CreateSecurityGroupWithRuleRequest struct {
	AK                string              `json:"account_key" binding:"required"`
	VpcId             string              `json:"vpc_id"`
	RegionId          string              `json:"region_id"`
	SecurityGroupName string              `json:"security_group_name"`
	SecurityGroupType string              `json:"security_group_type"`
	Rules             []service.GroupRule `json:"rules"`
}

func (c *CreateSecurityGroupWithRuleRequest) Check() bool {
	return c.SecurityGroupName != "" && c.RegionId != "" && c.VpcId != ""
}

type CreateNetworkRequest struct {
	Provider          string              `json:"provider" binding:"required,mustIn=cloud"`
	RegionId          string              `json:"region_id" binding:"required"`
	CidrBlock         string              `json:"cidr_block" binding:"required"`
	VpcName           string              `json:"vpc_name" binding:"required"`
	ZoneId            string              `json:"zone_id" binding:"required"`
	SwitchCidrBlock   string              `json:"switch_cidr_block" binding:"required"`
	GatewayIp         string              `json:"gateway_ip" binding:"required"`
	SwitchName        string              `json:"switch_name" binding:"required"`
	SecurityGroupName string              `json:"security_group_name" binding:"required"`
	SecurityGroupType string              `json:"security_group_type"`
	AK                string              `json:"ak" binding:"required"`
	Rules             []service.GroupRule `json:"rules"`
}

type SyncNetworkRequest struct {
	Provider   string `json:"provider" binding:"required"`
	RegionId   string `json:"region_id" binding:"required"`
	AccountKey string `json:"account_key" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type CreateCloudAccountRequest struct {
	AccountName   string `json:"account_name"`
	Provider      string `json:"provider" binding:"required,mustIn=cloud"`
	AccountKey    string `json:"account_key" binding:"required"`
	AccountSecret string `json:"account_secret" binding:"required"`
}

type EditCloudAccountRequest struct {
	AccountId   string `json:"account_id"`
	AccountName string `json:"account_name"`
	Provider    string `json:"provider"`
}

type EditOrgRequest struct {
	OrgId   int64  `json:"org_id" binding:"required"`
	OrgName string `json:"org_name" binding:"required"`
}

type CreateUserRequest struct {
	UserName string `json:"username" binding:"min=1,max=32"`
	Password string `json:"password" binding:"min=1,max=32"`
	UserType string `json:"user_type" binding:"required"`
}

type ModifyAdminPasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ModifyUsernameRequest struct {
	UserId      string `json:"user_id"`
	NewUsername string `json:"new_username"`
}

type ModifyUserTypeRequest struct {
	UserIds  []int64 `json:"user_ids" binding:"min=1"`
	UserType string  `json:"user_type" binding:"required"`
}

type UserStatusRequest struct {
	UserNames []string `json:"usernames"`
	Action    string   `json:"action"`
}

type CreateOrgRequest struct {
	OrgName  string `json:"org_name"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type ListClusterByTagsRequest struct {
	Tags       map[string]string `json:"tags" binding:"required"`
	PageNumber int               `json:"page_number"`
	PageSize   int               `json:"page_size"`
}

type SyncInstanceExpireTimeRequest struct {
	ClusterName string `json:"cluster_name" binding:"required"`
}

type CustomPublicCloudClusterRequest struct {
	ClusterName  string                        `json:"name" binding:"required"`
	ClusterDesc  string                        `json:"desc"`
	Provider     string                        `json:"provider" binding:"required,mustIn=cloud"`
	AccountKey   string                        `json:"account_key"`
	InstanceList []model.CustomClusterInstance `json:"instance_list" binding:"required,min=1"`
}

type CustomPrivateCloudClusterRequest struct {
	ClusterName  string                        `json:"name" binding:"required"`
	ClusterDesc  string                        `json:"desc"`
	InstanceList []model.CustomClusterInstance `json:"instance_list" binding:"required,min=1"`
}

type CheckInstanceConnectableRequest struct {
	InstanceList []model.CustomClusterInstance `json:"instance_list"`
}
type CreateKeyPairRequest struct {
	AK          string `json:"account_key" binding:"required"`
	Provider    string `json:"provider" binding:"required"`
	RegionId    string `json:"region_id" binding:"required"`
	KeyPairName string `json:"key_pair_name" binding:"required"`
}
type ImportKeyPairRequest struct {
	AK          string `json:"account_key" binding:"required"`
	Provider    string `json:"provider" binding:"required"`
	RegionId    string `json:"region_id" binding:"required"`
	KeyPairName string `json:"key_pair_name" binding:"required"`
	PrivateKey  string `json:"private_key" binding:"required"`
	PublicKey   string `json:"public_key" binding:"required"`
}

type ListKeyPairRequest struct {
	AK       string `form:"account_key" binding:"required"`
	Provider string `form:"provider" binding:"required"`
	RegionId string `form:"region_id" binding:"required"`
	types.PageReq
}
type ClusterAuthRequest struct {
	ClusterName string `form:"cluster_name" binding:"required"`
}

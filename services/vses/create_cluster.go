/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type CreateCluster struct {
	ClusterName string `json:"clusterName"`
	ElasticSearchVersionCode string `json:"elasticSearchVersionCode"`
	KibanaHttpPort string `json:"kibanaHttpPort"`
	KibanaUserName string `json:"kibanaUserName"`
	KibanaUserPassword string `json:"kibanaUserPassword"`
	SoftwareProductCode string `json:"softwareProductCode"`
	VpcNo int32 `json:"vpcNo"`
	ManagerNodeProductCode string `json:"managerNodeProductCode"`
	ManagerNodeSubnetNo int32 `json:"managerNodeSubnetNo"`
	DataNodeProductCode string `json:"dataNodeProductCode"`
	DataNodeCount int32 `json:"dataNodeCount"`
	DataNodeSubnetNo int32 `json:"dataNodeSubnetNo"`
	DataNodeStorageSize int32 `json:"dataNodeStorageSize"`
	LoginKeyName string `json:"loginKeyName"`
	IsDualManager bool `json:"isDualManager"`
	IsMasterOnlyNodeActivated bool `json:"isMasterOnlyNodeActivated"`
	MasterNodeSubnetNo int32 `json:"masterNodeSubnetNo,omitempty"`
	MasterNodeProductCode string `json:"masterNodeProductCode,omitempty"`
	MasterNodeCount int32 `json:"masterNodeCount,omitempty"`
}

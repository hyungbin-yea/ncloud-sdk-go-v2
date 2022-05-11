/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type VpcSubnetListResponseDetailVo struct {
	IsPublic bool `json:"isPublic,omitempty"`
	Permission string `json:"permission,omitempty"`
	StatusCode string `json:"statusCode,omitempty"`
	Subnet string `json:"subnet,omitempty"`
	SubnetName string `json:"subnetName,omitempty"`
	SubnetNo int32 `json:"subnetNo,omitempty"`
	VpcName string `json:"vpcName,omitempty"`
	VpcNo int32 `json:"vpcNo,omitempty"`
	ZoneName string `json:"zoneName,omitempty"`
	ZoneNo int32 `json:"zoneNo,omitempty"`
}

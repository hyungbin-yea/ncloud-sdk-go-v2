/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.apigw.ntruss.com/vpc/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type CreateNetworkAclDenyAllowGroupRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo"`

	// 네트워크ACL허용차단그룹이름
NetworkAclDenyAllowGroupName *string `json:"networkAclDenyAllowGroupName,omitempty"`

	// 네트워크ACL허용차단그룹설명
NetworkAclDenyAllowGroupDescription *string `json:"networkAclDenyAllowGroupDescription,omitempty"`
}

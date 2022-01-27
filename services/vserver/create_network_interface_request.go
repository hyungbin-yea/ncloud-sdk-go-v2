/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type CreateNetworkInterfaceRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// ACG번호리스트
AccessControlGroupNoList []*string `json:"accessControlGroupNoList"`

	// 네트워크인터페이스설명
NetworkInterfaceDescription *string `json:"networkInterfaceDescription,omitempty"`

	// 네트워크인터페이스이름
NetworkInterfaceName *string `json:"networkInterfaceName,omitempty"`

	// 서브넷번호
SubnetNo *string `json:"subnetNo"`

	// VPC번호
VpcNo *string `json:"vpcNo"`

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`

	// IP주소
Ip *string `json:"ip,omitempty"`

	// 보조IP리스트
SecondaryIpList []*string `json:"secondaryIpList,omitempty"`

	// 보조IP자동할당개수
SecondaryIpCount *int32 `json:"secondaryIpCount,omitempty"`
}

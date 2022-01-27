/*
 * vloadbalancer
 *
 * VPC Load Balancer 관련 API<br/>https://ncloud.apigw.ntruss.com/vloadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vloadbalancer

type DeleteLoadBalancerInstancesRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 로드밸런서인스턴스번호리스트
LoadBalancerInstanceNoList []*string `json:"loadBalancerInstanceNoList"`

	// 공인IP반납여부
ReturnPublicIpTogether *bool `json:"returnPublicIpTogether,omitempty"`
}

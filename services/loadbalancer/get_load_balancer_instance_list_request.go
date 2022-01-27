/*
 * loadbalancer
 *
 * <br/>https://ncloud.apigw.ntruss.com/loadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

type GetLoadBalancerInstanceListRequest struct {

	// 로드밸런서명
LoadBalancerName *string `json:"loadBalancerName,omitempty"`

	// 로드밸런서인스턴스번호리스트
LoadBalancerInstanceNoList []*string `json:"loadBalancerInstanceNoList,omitempty"`

	// 네트워크 구분코드
NetworkUsageTypeCode *string `json:"networkUsageTypeCode,omitempty"`

	// 리전번호
RegionNo *string `json:"regionNo,omitempty"`

	// ZONE번호
ZoneNo *string `json:"zoneNo,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// 소팅대상
SortedBy *string `json:"sortedBy,omitempty"`

	// 소팅순서
SortingOrder *string `json:"sortingOrder,omitempty"`
}

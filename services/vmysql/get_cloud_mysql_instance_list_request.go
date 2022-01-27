/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type GetCloudMysqlInstanceListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo,omitempty"`

	// Subnet번호
SubnetNo *string `json:"subnetNo,omitempty"`

	// CloudMysql서비스이름
CloudMysqlServiceName *string `json:"cloudMysqlServiceName,omitempty"`

	// CloudMysql인스턴스번호리스트
CloudMysqlInstanceNoList []*string `json:"cloudMysqlInstanceNoList,omitempty"`

	// CloudMysql서버이름
CloudMysqlServerName *string `json:"cloudMysqlServerName,omitempty"`

	// CloudMysql서버인스턴스번호리스트
CloudMysqlServerInstanceNoList []*string `json:"cloudMysqlServerInstanceNoList,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`
}

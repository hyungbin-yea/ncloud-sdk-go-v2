/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type PortForwardingConfiguration struct {

	// 포트포워딩설정번호
PortForwardingConfigurationNo *string `json:"portForwardingConfigurationNo,omitempty"`

	// 포트포워딩공인IP
PortForwardingPublicIp *string `json:"portForwardingPublicIp,omitempty"`

	// 서버인스턴스번호리스트
ServerInstanceNoList []*string `json:"serverInstanceNoList,omitempty"`
}

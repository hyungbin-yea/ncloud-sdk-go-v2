/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type SetMemberServerImageSharingPermissionRequest struct {

	// 회원서버이미지인스턴스번호
MemberServerImageInstanceNo *string `json:"memberServerImageInstanceNo"`

	// 대상로그인ID리스트
TargetLoginIdList []*string `json:"targetLoginIdList,omitempty"`
}

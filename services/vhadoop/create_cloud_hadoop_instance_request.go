/*
 * vhadoop
 *
 * <br/>https://ncloud.apigw.ntruss.com/vhadoop/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vhadoop

type CreateCloudHadoopInstanceRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo"`

	// Cloud Hadoop 클러스터 이름
CloudHadoopClusterName *string `json:"cloudHadoopClusterName"`

	// Cloud Hadoop이미지상품코드
CloudHadoopImageProductCode *string `json:"cloudHadoopImageProductCode,omitempty"`

	// Cloud Hadoop 클러스터 유형 코드
CloudHadoopClusterTypeCode *string `json:"cloudHadoopClusterTypeCode"`

	// Cloud Hadoop Add-On 코드 리스트
CloudHadoopAddOnCodeList []*string `json:"cloudHadoopAddOnCodeList,omitempty"`

	// 클러스터 관리자 계정
CloudHadoopAdminUserName *string `json:"cloudHadoopAdminUserName"`

	// 클러스터 관리자 계정 패스워드
CloudHadoopAdminUserPassword *string `json:"cloudHadoopAdminUserPassword"`

	// 인증키 이름
LoginKeyName *string `json:"loginKeyName,omitempty"`

	// 버킷 이름
BucketName *string `json:"bucketName"`

	// Cloud Hadoop 엣지노드 상품 코드
EdgeNodeProductCode *string `json:"edgeNodeProductCode,omitempty"`

	// 엣지노드의 Subnet 번호
EdgeNodeSubnetNo *string `json:"edgeNodeSubnetNo"`

	// Cloud Hadoop 마스터노드 상품 코드
MasterNodeProductCode *string `json:"masterNodeProductCode,omitempty"`

	// 마스터노드의 Subnet 번호
MasterNodeSubnetNo *string `json:"masterNodeSubnetNo"`

	// 마스터노드의 데이터 스토리지 타입 코드
MasterNodeDataStorageTypeCode *string `json:"masterNodeDataStorageTypeCode,omitempty"`

	// 마스터노드의 데이터 스토리지 크기
MasterNodeDataStorageSize *int32 `json:"masterNodeDataStorageSize"`

	// Cloud Hadoop 작업자노드 상품 코드
WorkerNodeProductCode *string `json:"workerNodeProductCode,omitempty"`

	// 작업자노드 개수
WorkerNodeCount *int32 `json:"workerNodeCount,omitempty"`

	// 작업자노드의 Subnet 번호
WorkerNodeSubnetNo *string `json:"workerNodeSubnetNo"`

	// 작업자노드의 데이터 스토리지 타입 코드
WorkerNodeDataStorageTypeCode *string `json:"workerNodeDataStorageTypeCode,omitempty"`

	// 작업자노드의 데이터 스토리지 크기
WorkerNodeDataStorageSize *int32 `json:"workerNodeDataStorageSize"`

	// 커버로스 인증 구성 여부
UseKdc *string `json:"useKdc,omitempty"`

	// KDC의 Realm 정보
KdcRealm *string `json:"kdcRealm,omitempty"`

	// KDC admin 계정의 패스워드
KdcPassword *string `json:"kdcPassword,omitempty"`

	// Cloud Hadoop 부트스트랩 스크립트 사용 여부
UseBootstrapScript *string `json:"useBootstrapScript,omitempty"`

	// Cloud Hadoop 부트스트랩 스크립트 사용 여부
BootstrapScript *string `json:"bootstrapScript,omitempty"`
}

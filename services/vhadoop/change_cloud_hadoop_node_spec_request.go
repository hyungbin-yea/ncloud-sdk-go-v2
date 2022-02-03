/*
 * vhadoop
 *
 * <br/>https://ncloud.apigw.ntruss.com/vhadoop/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vhadoop

type ChangeCloudHadoopNodeSpecRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// Cloud Hadoop 인스턴스번호
CloudHadoopInstanceNo *string `json:"cloudHadoopInstanceNo"`

	// Cloud Hadoop 작업자 노드 상품 개수
WorkerNodeProductCode *string `json:"workerNodeProductCode"`

	// Cloud Hadoop 엣지 노드 상품 개수
EdgeNodeProductCode *string `json:"edgeNodeProductCode"`

	// Cloud Hadoop 마스터 노드 상품 개수
MasterNodeProductCode *string `json:"masterNodeProductCode"`
}

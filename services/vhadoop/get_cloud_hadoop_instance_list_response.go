/*
 * vhadoop
 *
 * <br/>https://ncloud.apigw.ntruss.com/vhadoop/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vhadoop

type GetCloudHadoopInstanceListResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

	// CloudHadoop인스턴스리스트
CloudHadoopInstanceList *CloudHadoopInstanceList `json:"cloudHadoopInstanceList,omitempty"`
}

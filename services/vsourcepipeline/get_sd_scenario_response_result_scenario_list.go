/*
 * vsourcepipeline
 *
 * <br/>https://vpcsourcepipeline.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vsourcepipeline

type GetSdScenarioResponseResultScenarioList struct {
	Id *int32 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Source *GetSdScenarioResponseResultSource `json:"source,omitempty"`
}

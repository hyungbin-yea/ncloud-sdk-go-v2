/*
 * sourcepipeline
 *
 * <br/>https://sourcepipeline.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sourcepipeline

type GetSdScenarioResponseResultSourceTarget struct {
	ProjectName *string `json:"projectName,omitempty"`

	Manifest []*string `json:"manifest,omitempty"`

	File *string `json:"file,omitempty"`
}

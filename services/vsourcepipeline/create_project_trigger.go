/*
 * vsourcepipeline
 *
 * <br/>https://vpcsourcepipeline.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vsourcepipeline

type CreateProjectTrigger struct {
	Setting *bool `json:"setting,omitempty"`

	Sourcecommit []*GetScTargetInfo `json:"sourcecommit,omitempty"`
}

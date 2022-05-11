/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type GetOsMonitoringData struct {
	TimeStart int32 `json:"timeStart"`
	TimeEnd int32 `json:"timeEnd"`
	Metric string `json:"metric"`
	ComputeInstanceNo string `json:"computeInstanceNo"`
	Interval string `json:"interval,omitempty"`
}

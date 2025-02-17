# Go API client for vsourcepipeline

    <br/>https://vpcsourcepipeline.apigw.ntruss.com/api/v1

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2022-05-25T13:56:57Z
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.NcpGoForNcloudClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
"./vsourcepipeline"
```

## Documentation for API Endpoints

All URIs are relative to *https://vpcsourcepipeline.apigw.ntruss.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*V1Api* | [**GetProjects**](docs/V1Api.md#GetProjects) | **Get** /project | 
*V1Api* | [**CreateProject**](docs/V1Api.md#CreateProject) | **Post** /project | 
*V1Api* | [**DeleteProject**](docs/V1Api.md#DeleteProject) | **Delete** /project/{projectId} | 
*V1Api* | [**StartProject**](docs/V1Api.md#StartProject) | **Post** /project/{projectId}/do | 
*V1Api* | [**GetProject**](docs/V1Api.md#GetProject) | **Get** /project/{projectId} | 
*V1Api* | [**GetProjectHistories**](docs/V1Api.md#GetProjectHistories) | **Get** /project/{projectId}/history | 
*V1Api* | [**CancelProject**](docs/V1Api.md#CancelProject) | **Post** /project/{projectId}/history/{historyId}/cancel | 
*V1Api* | [**GetProjectHistory**](docs/V1Api.md#GetProjectHistory) | **Get** /project/{projectId}/history/{historyId} | 
*V1Api* | [**ChangeProject**](docs/V1Api.md#ChangeProject) | **Patch** /project/{projectId} | 
*V1Api* | [**GetSourcebuildProjects**](docs/V1Api.md#GetSourcebuildProjects) | **Get** /sourcebuild/project | 
*V1Api* | [**GetSourcecommitRepositories**](docs/V1Api.md#GetSourcecommitRepositories) | **Get** /sourcecommit/repository | 
*V1Api* | [**GetSourcecommitRepositoryBranches**](docs/V1Api.md#GetSourcecommitRepositoryBranches) | **Get** /sourcecommit/repository/{repositoryName}/branch | 
*V1Api* | [**GetSourcedeployProjects**](docs/V1Api.md#GetSourcedeployProjects) | **Get** /sourcedeploy/project | 
*V1Api* | [**GetSourcedeployProjectStages**](docs/V1Api.md#GetSourcedeployProjectStages) | **Get** /sourcedeploy/project/{projectId}/stage | 
*V1Api* | [**GetSourcedeployProjectScenarios**](docs/V1Api.md#GetSourcedeployProjectScenarios) | **Get** /sourcedeploy/project/{projectId}/stage/{stageId}/scenario | 


## Documentation For Models

 - [CancelPipelineResponse](docs/CancelPipelineResponse.md)
 - [ChangeProject](docs/ChangeProject.md)
 - [ChangeProjectReponse](docs/ChangeProjectReponse.md)
 - [CreateProject](docs/CreateProject.md)
 - [CreateProjectConfig](docs/CreateProjectConfig.md)
 - [CreateProjectConfigTarget](docs/CreateProjectConfigTarget.md)
 - [CreateProjectConfigTargetInfo](docs/CreateProjectConfigTargetInfo.md)
 - [CreateProjectResponse](docs/CreateProjectResponse.md)
 - [CreateProjectTasks](docs/CreateProjectTasks.md)
 - [CreateProjectTrigger](docs/CreateProjectTrigger.md)
 - [DeleteProjectResponse](docs/DeleteProjectResponse.md)
 - [GetHistoryDetailResponse](docs/GetHistoryDetailResponse.md)
 - [GetHistoryDetailResponseConfig](docs/GetHistoryDetailResponseConfig.md)
 - [GetHistoryDetailResponseConfigTarget](docs/GetHistoryDetailResponseConfigTarget.md)
 - [GetHistoryDetailResponseHistory](docs/GetHistoryDetailResponseHistory.md)
 - [GetHistoryDetailResponseTasks](docs/GetHistoryDetailResponseTasks.md)
 - [GetHistoryListResponse](docs/GetHistoryListResponse.md)
 - [GetHistoryListResponseHistoryList](docs/GetHistoryListResponseHistoryList.md)
 - [GetProjectDetailResponse](docs/GetProjectDetailResponse.md)
 - [GetProjectDetailResponseConfig](docs/GetProjectDetailResponseConfig.md)
 - [GetProjectDetailResponseConfigTarget](docs/GetProjectDetailResponseConfigTarget.md)
 - [GetProjectDetailResponseConfigTargetInfo](docs/GetProjectDetailResponseConfigTargetInfo.md)
 - [GetProjectDetailResponseConfigTargetInfoWorkspace](docs/GetProjectDetailResponseConfigTargetInfoWorkspace.md)
 - [GetProjectDetailResponseTasks](docs/GetProjectDetailResponseTasks.md)
 - [GetProjectListResponse](docs/GetProjectListResponse.md)
 - [GetProjectTrigger](docs/GetProjectTrigger.md)
 - [GetSbProjectResponse](docs/GetSbProjectResponse.md)
 - [GetSbProjectResponseProjectList](docs/GetSbProjectResponseProjectList.md)
 - [GetSbProjectResponseSource](docs/GetSbProjectResponseSource.md)
 - [GetScBranchResponse](docs/GetScBranchResponse.md)
 - [GetScBranchResponseBranchList](docs/GetScBranchResponseBranchList.md)
 - [GetScRepositoryResposne](docs/GetScRepositoryResposne.md)
 - [GetScTargetInfo](docs/GetScTargetInfo.md)
 - [GetSdProjectResponse](docs/GetSdProjectResponse.md)
 - [GetSdScenarioResponse](docs/GetSdScenarioResponse.md)
 - [GetSdScenarioResponseResultScenarioList](docs/GetSdScenarioResponseResultScenarioList.md)
 - [GetSdScenarioResponseResultSource](docs/GetSdScenarioResponseResultSource.md)
 - [GetSdScenarioResponseResultSourceTarget](docs/GetSdScenarioResponseResultSourceTarget.md)
 - [GetSdStageResponse](docs/GetSdStageResponse.md)
 - [StartPipelineResponse](docs/StartPipelineResponse.md)


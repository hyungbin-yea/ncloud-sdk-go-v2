/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type Product struct {

	// 상품코드
ProductCode *string `json:"productCode,omitempty"`

	// 상품이름
ProductName *string `json:"productName,omitempty"`

	// 상품유형
ProductType *CommonCode `json:"productType,omitempty"`

	// 상품설명
ProductDescription *string `json:"productDescription,omitempty"`

	// 인프라자원유형
InfraResourceType *CommonCode `json:"infraResourceType,omitempty"`

	// 인프라자원상세유형
InfraResourceDetailType *CommonCode `json:"infraResourceDetailType,omitempty"`

	// CPU개수
CpuCount *int32 `json:"cpuCount,omitempty"`

	// 메모리사이즈
MemorySize *int64 `json:"memorySize,omitempty"`

	// 기본블록스토리지사이즈
BaseBlockStorageSize *int64 `json:"baseBlockStorageSize,omitempty"`

	// 플랫폼유형
PlatformType *CommonCode `json:"platformType,omitempty"`

	// OS정보
OsInformation *string `json:"osInformation,omitempty"`

	// 디스크유형
DiskType *CommonCode `json:"diskType,omitempty"`

	// DB유형코드
DbKindCode *string `json:"dbKindCode,omitempty"`

	// 추가블록스토리지사이즈
AddBlockStorageSize *int64 `json:"addBlockStorageSize,omitempty"`

	// 세대코드
GenerationCode *string `json:"generationCode,omitempty"`
}

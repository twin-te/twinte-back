package restv3

import openapi_types "github.com/oapi-codegen/runtime/types"

func toApiUUID[T ~[16]byte](id T) openapi_types.UUID {
	return openapi_types.UUID(id)
}

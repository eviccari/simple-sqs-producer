package dtos

// SuccessResponse - Describe SuccessResponse message with generated protocolID
type SuccessResponse struct {
	ProtocolID string `json:"protocolID"`
}

// NewSuccessResponse - Create new SuccessResponse instance
func NewSuccessResponse(protocolID string) *SuccessResponse {
	return &SuccessResponse{
		ProtocolID: protocolID,
	}
}

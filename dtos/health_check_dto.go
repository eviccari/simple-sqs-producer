package dtos

// HealthCheckDTO - Describe HealthCheckDTO struct
type HealthCheckDTO struct {
	Message string `json:"message"`
}

// NewHealthCheckDTO - Create new HealthCheckDTO instance with default success message
func NewHealthCheckDTO() *HealthCheckDTO {
	return &HealthCheckDTO{
		Message: "check system status: OK",
	}
}

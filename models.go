package bigpanda

// IncidentTag -
type IncidentTag struct {
	ID          		string `json:"id"`
	Name        		string `json:"name"`
	Description 		string `json:"description"`
	Type        		string `json:"type"`
	Active:     		bool `json:"active"`
	Deleted:    		bool `json:"deleted"`
	CanManualInput:     bool `json:"canManualInput"`
	CreatedAt:      	int64 `json:"created_at"`
	CreatedBy:        	string `json:"created_by"`
	UpdatedAt:      	int64 `json:"updated_at"`
	UpdatedBy:        	string `json:"updated_by"`
}

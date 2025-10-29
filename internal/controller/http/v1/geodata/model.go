package geodata

// ProcessGeodataInput is input model of process user's geodata request.
type ProcessGeodataInput struct {
	UserID    int64   `json:"user_id"`   // User ID.
	Latitude  float64 `json:"latitude"`  // Latitude of user's coordinates.
	Longitude float64 `json:"longitude"` // Longitude of user's coordinates.
} // @name ProcessGeodataInput

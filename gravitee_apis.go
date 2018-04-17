package gravitee

import (
	"fmt"
)

type ApiInfo struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	Version         string        `json:"version"`
	Description     string        `json:"description"`
	Visibility      string        `json:"visibility"`
	State           string        `json:"state"`
	Views           []string      `json:"views"`
	Labels          []string      `json:"labels"`
	Manageable      bool          `json:"manageable"`
	NumberOfRatings int           `json:"numberOfRatings"`
	CreatedAt       int           `json:"created_at"`
	UpdatedAdd      int           `json:"updated_at"`
	Owner           UserReference `json:"owner"`
	PictureURL      string        `json:"picture_url"`
	ContextPath     string        `json:"context_path"`
}

func (ai ApiInfo) String() string {
	return fmt.Sprintf("%v (%v)", ai.Name, ai.ID)
}

// GetAllAPIs retrieves a list of all APIs registered in Gravitee.
func (s *GraviteeSession) GetAllAPIs() ([]ApiInfo, error) {
	var result *[]ApiInfo

	err := s.getForEntity(&result, "apis")

	if err != nil {
		return nil, err
	}

	return *result, nil
}

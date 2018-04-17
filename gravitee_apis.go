package gravitee

import (
	"fmt"
)

// ApiState is an enumeration of possible *State* to be used for an API.
type ApiState string

const (
	ApiState_Started ApiState = "started"
	ApiState_Stopped ApiState = "stopped"
)

// ApiVisibility is an enumeration of possible *Visibility* to be used for an API.
type ApiVisibility string

const (
	ApiVisibility_Private ApiVisibility = "private"
	ApiVisibility_Public  ApiVisibility = "public"
)

type ApiInfo struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	Version         string        `json:"version"`
	Description     string        `json:"description"`
	Visibility      ApiVisibility `json:"visibility"`
	State           ApiState      `json:"state"`
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
	return fmt.Sprintf("%v (%v, %v)", ai.Name, ai.ID, ai.State)
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

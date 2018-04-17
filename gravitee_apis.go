package gravitee

import (
	"fmt"
	"strings"
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

// ApiMetadataFormat is an enumeration of possible *Format* to be used for an API metdata entry.
type ApiMetadataFormat string

const (
	ApiMetadataFormat_String  ApiMetadataFormat = "string"
	ApiMetadataFormat_Numeric ApiMetadataFormat = "numeric"
	ApiMetadataFormat_Boolean ApiMetadataFormat = "boolean"
	ApiMetadataFormat_Date    ApiMetadataFormat = "date"
	ApiMetadataFormat_Mail    ApiMetadataFormat = "mail"
	ApiMetadataFormat_URL     ApiMetadataFormat = "url"
)

type ApiMetadata struct {
	Key          string            `json:"key"`
	Name         string            `json:"name"`
	Format       ApiMetadataFormat `json:"format"`
	LocalValue   string            `json:"value"`
	DefaultValue string            `json:"defaultValue"`
	ApiID        string            `json:"apiId"`
}

func (ai ApiMetadata) Value() string {
	if ai.LocalValue != "" {
		return ai.LocalValue
	}

	return ai.DefaultValue
}

func (ai ApiMetadata) String() string {
	return fmt.Sprintf("%v = %v [%v]", ai.Name, ai.Value(), ai.Format)
}

type ApiDetailsEndpoint struct {
	Name     string `json:"name"`
	Target   string `json:"target"`
	Weight   int    `json:"weight"`
	IsBackup bool   `json:"backup"`
	Type     string `json:"type"`

	Http struct {
		ConnectTimeoutMS         int  `json:"connectTimeout"`
		IdleTimeoutMS            int  `json:"idleTimeout"`
		ReadTimeoutMS            int  `json:"readTimeout"`
		KeepAlive                bool `json:"keepAlive"`
		Pipelining               bool `json:"pipelining"`
		MaxConcurrentConnections int  `json:"maxConcurrentConnections"`
		UseCompression           bool `json:"useCompression"`
		FollowRedirects          bool `json:"followRedirects"`
	} `json:"http"`
}

func (ai ApiDetailsEndpoint) String() string {
	return fmt.Sprintf("%v (%v, %v)", ai.Name, ai.Target, ai.Type)
}

type ApiDetails struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Version     string        `json:"version"`
	Description string        `json:"description"`
	Visibility  ApiVisibility `json:"visibility"`
	State       ApiState      `json:"state"`
	Tags        []string      `json:"tags"`
	Labels      []string      `json:"labels"`
	//Paths       []string      `json:"paths"`
	CreatedAt   int           `json:"created_at"`
	UpdatedAdd  int           `json:"updated_at"`
	DeployedAt  int           `json:"deployed_at"`
	Owner       UserReference `json:"owner"`
	PictureURL  string        `json:"picture_url"`
	ContextPath string        `json:"context_path"`
	Proxy       struct {
		ContextPath      string               `json:"context_path"`
		StripContextPath bool                 `json:"strip_context_path"`
		LoggingMode      string               `json:"loggingMode"`
		Endpoints        []ApiDetailsEndpoint `json:"endpoints"`

		LoadBalancing struct {
			Type string `json:"type"`
		} `json:"load_balancing"`
	} `json:"proxy"`
}

func (ai ApiDetails) String() string {
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

// GetAPIsByLabel retrieves a list of all APIs registered in Gravitee.
func (s *GraviteeSession) GetAPIsByLabel(label string) ([]ApiInfo, error) {
	result, err := s.GetAllAPIs()
	if err != nil {
		return nil, err
	}

	filtered := make([]ApiInfo, 0)

	for _, ai := range result {
		for _, lbl := range ai.Labels {
			if strings.EqualFold(lbl, label) {
				filtered = append(filtered, ai)
				break
			}
		}
	}

	return filtered, nil
}

// GetAPI retrieves details on an API registered in Gravitee.
func (s *GraviteeSession) GetAPI(id string) (*ApiDetails, error) {
	var result *ApiDetails

	err := s.getForEntity(&result, "apis", id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAPIMetadata retrieves the metadata on an API registered in Gravitee.
func (s *GraviteeSession) GetAPIMetadata(id string) ([]ApiMetadata, error) {
	var result *[]ApiMetadata

	err := s.getForEntity(&result, "apis", id, "metadata")
	if err != nil {
		return nil, err
	}

	return *result, nil
}

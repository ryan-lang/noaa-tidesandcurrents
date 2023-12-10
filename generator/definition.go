package generator

import (
	"os"
	"sort"

	"gopkg.in/yaml.v3"
)

type (
	ClientDefinition struct {
		DataAPI     DataApiDefinition     `yaml:"data_api"`
		MetadataAPI MetadataApiDefinition `yaml:"metadata_api"`
	}

	DataApiDefinition struct {
		Products []ProductDefinition
		Model    []ModelDefinition
	}

	MetadataApiDefinition struct {
		Model            []ModelDefinition
		StationResources []StationResourceDefinition
	}

	ProductDefinition struct {
		ProductID           string `yaml:"product_id"`
		Name                string `yaml:"name"`
		RequestType         string `yaml:"request"`
		ResponseType        string `yaml:"response"`
		TestStation         string `yaml:"test_station"`
		TestHistoricalRange bool   `yaml:"test_historical_range"`
	}

	ModelDefinition struct {
		Name            string
		IsResponse      bool
		CustomUnmarshal bool
		Fields          []FieldDefinition
	}

	FieldDefinition struct {
		Name      string `yaml:"name"`
		Type      string `yaml:"type"`
		UrlParam  string `yaml:"url"`
		JsonParam string `yaml:"json"`
		JsonType  string `yaml:"json_type"`
		Required  bool   `yaml:"required"`
		Default   string `yaml:"default"`
	}

	StationResourceDefinition struct {
		ResourceID   string   `yaml:"resource_id"`
		Name         string   `yaml:"name"`
		RequestType  string   `yaml:"request"`
		ResponseType string   `yaml:"response"`
		Availability []string `yaml:"availability"`
		TestStation  string   `yaml:"test_station"`
	}
)

func ParseClientDefinition(path string) (*ClientDefinition, error) {

	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var def ClientDefinition
	err = yaml.Unmarshal(body, &def)
	if err != nil {
		return nil, err
	}

	return &def, nil
}

func (m *ModelDefinition) UnmarshalYAML(node *yaml.Node) error {
	var tmp struct {
		IsResponse      bool                       `yaml:"is_response"`
		CustomUnmarshal bool                       `yaml:"custom_unmarshal"`
		Fields          map[string]FieldDefinition `yaml:"fields"`
	}
	if err := node.Decode(&tmp); err != nil {
		return err
	}

	for name, field := range tmp.Fields {
		field.Name = name
		m.Fields = append(m.Fields, field)
	}

	// sort the fields by name
	sort.Slice(m.Fields, func(i, j int) bool {
		return m.Fields[i].Name < m.Fields[j].Name
	})

	m.IsResponse = tmp.IsResponse
	m.CustomUnmarshal = tmp.CustomUnmarshal

	return nil
}

func (m *DataApiDefinition) UnmarshalYAML(node *yaml.Node) error {
	var tmp struct {
		Products []ProductDefinition        `yaml:"products"`
		Model    map[string]ModelDefinition `yaml:"model"`
	}
	if err := node.Decode(&tmp); err != nil {
		return err
	}

	for name, model := range tmp.Model {
		model.Name = name
		m.Model = append(m.Model, model)
	}

	// sort the models by name
	sort.Slice(m.Model, func(i, j int) bool {
		return m.Model[i].Name < m.Model[j].Name
	})

	m.Products = tmp.Products

	return nil
}

func (m *MetadataApiDefinition) UnmarshalYAML(node *yaml.Node) error {
	var tmp struct {
		Model            map[string]ModelDefinition  `yaml:"model"`
		StationResources []StationResourceDefinition `yaml:"station_resources"`
	}
	if err := node.Decode(&tmp); err != nil {
		return err
	}

	for name, model := range tmp.Model {
		model.Name = name
		m.Model = append(m.Model, model)
	}

	// sort the models by name
	sort.Slice(m.Model, func(i, j int) bool {
		return m.Model[i].Name < m.Model[j].Name
	})

	m.StationResources = tmp.StationResources

	return nil
}

func (d *ClientDefinition) Validate() error {

	// TODO we can't have any method name conflicts across apis

	return nil
}

func (d *FieldDefinition) GetUrlParam() string {
	if d.UrlParam != "" {
		return d.UrlParam
	}

	return ToSnakeCase(d.Name)
}

func (d *FieldDefinition) GetJsonParam() string {
	if d.JsonParam != "" {
		return d.JsonParam
	}

	return ToSnakeCase(d.Name)
}

func (d *FieldDefinition) GetJsonType() string {
	if d.JsonType != "" {
		return d.JsonType
	}

	return d.Type
}

func (d *DataApiDefinition) GetModel(name string) ModelDefinition {
	for _, model := range d.Model {
		if model.Name == name {
			return model
		}
	}
	return ModelDefinition{}
}

func (d *MetadataApiDefinition) GetModel(name string) ModelDefinition {
	for _, model := range d.Model {
		if model.Name == name {
			return model
		}
	}
	return ModelDefinition{}
}

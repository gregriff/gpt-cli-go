package openai

import (
	"fmt"
	"strings"

	"github.com/gregriff/ducky/models"
)

type OpenAIModelConfig struct {
	models.BaseModelConfig
	SupportsTemperature *bool
	Reasoning           *bool
}

// GetOpenAIModelConfigs returns a map of OpenAI model names to properties about those models
var OpenAIModelConfigs = map[string]OpenAIModelConfig{
	"o3": {
		BaseModelConfig: models.BaseModelConfig{
			Id:           "o3",
			PromptCost:   10. / 1_000_000,
			ResponseCost: 40. / 1_000_000,
		},
		Reasoning:           models.BoolPtr(true),
		SupportsTemperature: models.BoolPtr(false),
	},
	"o4-mini": {
		BaseModelConfig: models.BaseModelConfig{
			Id:           "o4-mini",
			PromptCost:   1.1 / 1_000_000,
			ResponseCost: 4.4 / 1_000_000,
		},
		Reasoning:           models.BoolPtr(true),
		SupportsTemperature: models.BoolPtr(false),
	},
	"4o-mini": {
		BaseModelConfig: models.BaseModelConfig{
			Id:           "gpt-4o-mini",
			PromptCost:   .15 / 1_000_000,
			ResponseCost: .075 / 1_000_000,
		},
	},
	"4o": {
		BaseModelConfig: models.BaseModelConfig{
			Id:           "gpt-4o",
			PromptCost:   2.5 / 1_000_000,
			ResponseCost: 10. / 1_000_000,
		},
	},
	"4.1": {
		BaseModelConfig: models.BaseModelConfig{
			Id:           "gpt-4.1",
			PromptCost:   2. / 1_000_000,
			ResponseCost: 8. / 1_000_000,
		},
	},
	"4.1-mini": {
		BaseModelConfig: models.BaseModelConfig{
			Id:           "gpt-4.1-mini",
			PromptCost:   .4 / 1_000_000,
			ResponseCost: 1.6 / 1_000_000,
		},
	},
	"4.1-nano": {
		BaseModelConfig: models.BaseModelConfig{
			Id:           "gpt-4.1-nano",
			PromptCost:   .1 / 1_000_000,
			ResponseCost: .4 / 1_000_000,
		},
	},
}

// ValidateModelName validates that a modelName is one of our supported models. If so, it returns the modelId
func ValidateModelName(modelName string) error {
	if _, exists := OpenAIModelConfigs[modelName]; !exists {
		var validNames []string
		for name := range OpenAIModelConfigs {
			validNames = append(validNames, name)
		}
		err := fmt.Errorf("invalid model name '%s'. Valid options: %s", modelName, strings.Join(validNames, ", "))
		return err
	}
	return nil
}

// GetValidModelNames returns the keys of OpenAIModelConfigs, our supported OpenAI models
func GetValidModelNames() []string {
	var names []string
	for name := range OpenAIModelConfigs {
		names = append(names, name)
	}
	return names
}

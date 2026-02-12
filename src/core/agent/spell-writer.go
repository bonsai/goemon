package agent

import (
	"encoding/json"
	"fmt"
	"goemon/src/core/bridge"
	"goemon/src/core/db"
	"path/filepath"
)

type SpellWriter struct {
	DB *db.VectorDB
}

func NewSpellWriter(vdb *db.VectorDB) *SpellWriter {
	return &SpellWriter{
		DB: vdb,
	}
}

func (sw *SpellWriter) CastSpell(seed string) (string, error) {
	fragments, err := sw.DB.GetRandomFragments(3)
	if err != nil {
		return "", fmt.Errorf("failed to steal fragments: %v", err)
	}

	prompt := fmt.Sprintf("Seed: %s\nFragments: %v\nGenerate a new spell.", seed, fragments)
	fmt.Printf("Requesting GLM with prompt: %s\n", prompt)

	fragmentsJSON, _ := json.Marshal(fragments)
	scriptPath := filepath.Join("src", "core", "agent", "python", "glm_bridge.py")

	resultJSON, err := bridge.CallPythonBridge(scriptPath, seed, string(fragmentsJSON))
	if err != nil {
		return "", fmt.Errorf("GLM bridge failed: %v", err)
	}

	var response struct {
		Status string `json:"status"`
		Spell  string `json:"spell"`
	}
	if err := json.Unmarshal([]byte(resultJSON), &response); err != nil {
		return "", fmt.Errorf("failed to parse GLM response: %v", err)
	}

	return response.Spell, nil
}

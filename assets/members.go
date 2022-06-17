package assets

import (
	_ "embed"
	"encoding/json"
	"log"
)

type MembersDefinition struct {
	LegLeft0  []float32
	LegLeft1  []float32
	LegRight0 []float32
	LegRight1 []float32
	ArmLeft0  []float32
	ArmLeft1  []float32
	ArmRight0 []float32
	ArmRight1 []float32
}

func (md *MembersDefinition) AppendUniforms(m map[string]interface{}) map[string]interface{} {
	m["LegLeft0"] = md.LegLeft0
	m["LegLeft1"] = md.LegLeft1
	m["LegRight0"] = md.LegRight0
	m["LegRight1"] = md.LegRight1
	m["ArmLeft0"] = md.ArmLeft0
	m["ArmLeft1"] = md.ArmLeft1
	m["ArmRight0"] = md.ArmRight0
	m["ArmRight1"] = md.ArmRight1

	return m
}

var (
	//go:embed members/falling.json
	positionFallingSrc []byte
	PositionFalling    *MembersDefinition
	//go:embed members/sliding.json
	positionSlidingSrc []byte
	PositionSliding    *MembersDefinition
	//go:embed members/dashing.json
	positionDashingSrc []byte
	PositionDashing    *MembersDefinition
	//go:embed members/bouncing.json
	positionBouncingSrc []byte
	PositionBouncing    *MembersDefinition
)

func init() {
	PositionFalling = &MembersDefinition{}
	if err := json.Unmarshal(positionFallingSrc, PositionFalling); err != nil {
		log.Fatal(err)
	}

	PositionSliding = &MembersDefinition{}
	if err := json.Unmarshal(positionSlidingSrc, PositionSliding); err != nil {
		log.Fatal(err)
	}

	PositionDashing = &MembersDefinition{}
	if err := json.Unmarshal(positionDashingSrc, PositionDashing); err != nil {
		log.Fatal(err)
	}

	PositionBouncing = &MembersDefinition{}
	if err := json.Unmarshal(positionBouncingSrc, PositionBouncing); err != nil {
		log.Fatal(err)
	}
}

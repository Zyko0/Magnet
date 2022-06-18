package assets

import (
	_ "embed"
	"encoding/json"
	"log"
)

type Bones struct {
	LegLeft0  []float32
	LegLeft1  []float32
	LegRight0 []float32
	LegRight1 []float32
	ArmLeft0  []float32
	ArmLeft1  []float32
	ArmRight0 []float32
	ArmRight1 []float32
}

func (b *Bones) AppendUniforms(uniforms map[string]interface{}) map[string]interface{} {
	uniforms["LegLeft0"] = b.LegLeft0
	uniforms["LegLeft1"] = b.LegLeft1
	uniforms["LegRight0"] = b.LegRight0
	uniforms["LegRight1"] = b.LegRight1
	uniforms["ArmLeft0"] = b.ArmLeft0
	uniforms["ArmLeft1"] = b.ArmLeft1
	uniforms["ArmRight0"] = b.ArmRight0
	uniforms["ArmRight1"] = b.ArmRight1

	return uniforms
}

var (
	//go:embed bones/falling.json
	bonesFallingDataSrc []byte
	bonesFallingData    *Bones
	//go:embed bones/sliding.json
	bonesSlidingDataSrc []byte
	bonesSlidingData    *Bones
	//go:embed bones/dashing.json
	bonesDashingDataSrc []byte
	bonesDashingData    *Bones
	//go:embed bones/bouncing.json
	bonesBouncingDataSrc []byte
	bonesBouncingData    *Bones

	boneSets []*Bones
)

type BoneSet byte

const (
	BoneSetFalling BoneSet = iota
	BoneSetSliding
	BoneSetDashing
	BoneSetBouncing
)

func (b BoneSet) GetBones() *Bones {
	return boneSets[b]
}

func init() {
	bonesFallingData = &Bones{}
	if err := json.Unmarshal(bonesFallingDataSrc, bonesFallingData); err != nil {
		log.Fatal(err)
	}

	bonesSlidingData = &Bones{}
	if err := json.Unmarshal(bonesSlidingDataSrc, bonesSlidingData); err != nil {
		log.Fatal(err)
	}

	bonesDashingData = &Bones{}
	if err := json.Unmarshal(bonesDashingDataSrc, bonesDashingData); err != nil {
		log.Fatal(err)
	}

	bonesBouncingData = &Bones{}
	if err := json.Unmarshal(bonesBouncingDataSrc, bonesBouncingData); err != nil {
		log.Fatal(err)
	}

	boneSets = []*Bones{
		BoneSetFalling:  bonesFallingData,
		BoneSetSliding:  bonesSlidingData,
		BoneSetDashing:  bonesDashingData,
		BoneSetBouncing: bonesBouncingData,
	}
}

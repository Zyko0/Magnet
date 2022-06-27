package core

import (
	"github.com/Zyko0/Magnet/assets"
	"github.com/Zyko0/Magnet/logic"
)

type Difficulty struct {
	s string

	portalObstacleEnabled bool
	electricalRingEnabled bool
	availableShapeIndices []int

	ringAdvanceSpeed                 float32
	obstacleDeathSpawnZInterval      float32
	obstaclePortalSpawnTicksInterval uint64
}

func (d *Difficulty) String() string {
	return d.s
}

var (
	difficultyEasy = &Difficulty{
		s: "Ez",

		portalObstacleEnabled: false,
		electricalRingEnabled: false,
		availableShapeIndices: []int{
			assets.ShapeIndexLine,
			assets.ShapeIndexHalfQuad,
		},

		ringAdvanceSpeed:                 0.015,
		obstacleDeathSpawnZInterval:      1.5,
		obstaclePortalSpawnTicksInterval: logic.TPS * 5,
	}
	difficultyMedium = &Difficulty{
		s: "Medium",

		portalObstacleEnabled: true,
		electricalRingEnabled: false,
		availableShapeIndices: []int{
			assets.ShapeIndexLine,
			assets.ShapeIndexCross,
			assets.ShapeIndexHalfQuad,
		},

		ringAdvanceSpeed:                 0.02,
		obstacleDeathSpawnZInterval:      1.5,
		obstaclePortalSpawnTicksInterval: logic.TPS * 6.5,
	}
	difficultyHard = &Difficulty{
		s: "Hard",

		portalObstacleEnabled: true,
		electricalRingEnabled: true,
		availableShapeIndices: []int{
			assets.ShapeIndexLine,
			assets.ShapeIndexCross,
			assets.ShapeIndexEmptyLine,
			assets.ShapeIndexEmptyCross,
			assets.ShapeIndexHalfQuad,
		},

		ringAdvanceSpeed:                 0.025,
		obstacleDeathSpawnZInterval:      1.5,
		obstaclePortalSpawnTicksInterval: logic.TPS * 8.5,
	}
	difficultyExtreme = &Difficulty{
		s: "Extreme",

		portalObstacleEnabled: true,
		electricalRingEnabled: true,
		availableShapeIndices: []int{
			assets.ShapeIndexCross,
			assets.ShapeIndexEmptyLine,
			assets.ShapeIndexEmptyCross,
			assets.ShapeIndexHalfQuad,
			assets.ShapeIndexPizzaSlice,
		},

		ringAdvanceSpeed:                 0.03,
		obstacleDeathSpawnZInterval:      1.5,
		obstaclePortalSpawnTicksInterval: logic.TPS * 10,
	}
)

func getDifficulty(ticks uint64) *Difficulty {
	const (
		duration = 15 * logic.TPS

		easyRange1                 = duration
		mediumRange0, mediumRange1 = easyRange1, easyRange1 + duration
		hardRange0, hardRange1     = mediumRange1, mediumRange1 + duration
	)

	switch {
	case ticks <= easyRange1:
		return difficultyEasy
	case ticks >= mediumRange0 && ticks <= mediumRange1:
		return difficultyMedium
	case ticks >= hardRange0 && ticks <= hardRange1:
		return difficultyHard
	}

	return difficultyExtreme
}

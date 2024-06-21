package domain_service

import (
	"math"
	"time"

	"github.com/oapi-codegen/runtime/types"

	"first_move/generated/db/first_move/public/model"
	"first_move/generated/oapi"
)

func ParseUserBehavior(
	userBehaviors []model.UserBehavior,
	startDate time.Time,
	endDate time.Time,
) ([]oapi.Score, error) {
	res := []oapi.Score{}

	// userBehavior must be sorted be Date in ascending order
	sum := 0
	variance := 0
	for idx, ub := range userBehaviors {
		// calc current mean and variance
		sum += int(ub.Score)
		currMean := sum / (idx + 1)
		variance += int(math.Pow(float64(ub.Score-int32(currMean)), 2))

		stdDev, zScore := float64(0), float32(0)
		if idx > 0 {
			stdDev = math.Sqrt(float64(variance / (idx + 1)))
			zScore = (float32(ub.Score) - float32(currMean)) / float32(stdDev)
		}

		if ub.Date.Compare(endDate) == 1 {
			break
		}

		if ub.Date.Compare(startDate) > -1 {
			res = append(res, oapi.Score{
				CurrentScore:      int(ub.Score),
				Date:              types.Date{Time: ub.Date},
				Mean:              currMean,
				StandardDeviation: int(stdDev),
				ZScore:            zScore,
			})
		}
	}

	return res, nil
}

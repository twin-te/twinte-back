package timetablev1conv

import timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"

func ToPBRecommendedGrade(recommendedGrade timetabledomain.RecommendedGrade) int32 {
	return int32(recommendedGrade.Int())
}

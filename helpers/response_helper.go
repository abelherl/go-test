package helpers

import (
	"github.com/abelherl/go-test/constants"
	"github.com/lib/pq"
)

func EmptyArrayIfNil(arr []any) []any {
	if arr == nil {
		return []any{}
	}
	return arr
}

func EmptyPQStringArrayIfNil(arr pq.StringArray) pq.StringArray {
	if arr == nil {
		return pq.StringArray{}
	}
	return arr
}

func IsTagValid(tag string) bool {
	switch constants.Tag(tag) {
	case
		constants.TagProject, constants.TagTutorial, constants.TagOpinion, constants.TagLife, constants.TagResearch:
		return true
	default:
		return false
	}
}

func ValidateTags(tags []string) bool {
	for _, tag := range tags {
		if !IsTagValid(tag) {
			return false
		}
	}
	return true
}

func IsTechnologyValid(tech string) bool {
	switch constants.Technology(tech) {
	case
		// Frontend
		constants.TechFlutter, constants.TechBloc, constants.TechGetx, constants.TechProvider, constants.TechFigma,
		// Backend
		constants.TechGo, constants.TechGin, constants.TechGorm, constants.TechGraphQL, constants.TechREST, constants.TechPostgres,
		// Infra / Hosting
		constants.TechFirebase, constants.TechSupabase, constants.TechRender, constants.TechVercel, constants.TechGCP, constants.TechAWS,
		// Next Gen Tech
		constants.TechSolidity, constants.TechWeb3, constants.TechThreeJS, constants.TechAR, constants.TechVR, constants.TechOpenAI:
		return true
	default:
		return false
	}
}

func ValidateTechnologies(techs []string) bool {
	for _, tech := range techs {
		if !IsTechnologyValid(tech) {
			return false
		}
	}
	return true
}

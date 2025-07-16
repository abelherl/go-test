package constants

type Tag string

const (
	TagProject  Tag = "project"
	TagTutorial Tag = "tutorial"
	TagOpinion  Tag = "opinion"
	TagLife     Tag = "life"
	TagResearch Tag = "research"
)

type Technology string

const (
	// Frontend
	TechFlutter  Technology = "flutter"
	TechBloc     Technology = "bloc"
	TechGetx     Technology = "getx"
	TechProvider Technology = "provider"
	TechFigma    Technology = "figma"

	// Backend
	TechGo       Technology = "go"
	TechGin      Technology = "gin"
	TechGorm     Technology = "gorm"
	TechGraphQL  Technology = "graphql"
	TechREST     Technology = "rest"
	TechPostgres Technology = "postgres"

	// Infra / Hosting
	TechFirebase Technology = "firebase"
	TechSupabase Technology = "supabase"
	TechRender   Technology = "render"
	TechVercel   Technology = "vercel"
	TechGCP      Technology = "gcp"
	TechAWS      Technology = "aws"

	// Next Gen Tech
	TechSolidity Technology = "solidity"
	TechWeb3     Technology = "web3"
	TechThreeJS  Technology = "threejs"
	TechAR       Technology = "ar"
	TechVR       Technology = "vr"
	TechOpenAI   Technology = "openai"
)

package constants

const (
	SCHEDULER_ROUND_ROBIN   = "round_robin"
	SCHEDULER_LOW_LATENCY   = "low_latency"
	SCHEDULER_RANDOM        = "random"
	SCHEDULER_LOW_BANDIT    = "low_bandit"
	SCHEDULER_PEEKABOO      = "peekaboo"
	SCHEDULER_ECF           = "ecf"
	SCHEDULER_BLEST         = "blest"
	SCHEDULER_DQNA          = "dqnAgent"
	SCHEDULER_FIRST_PATH    = "first_path"
	SCHEDULER_OPTIMUM_SPLIT = "optimum_split"
	SCHEDULER_SHRAVAN_ML = "shravan_ml"
	SCHEDULER_REDUNDANT = "redundant"

	// Directory Names
	DEFAULT_TRAINING_DIR        = "online_training"
	DEFAULT_CLIENT_TRAINING_DIR = "client_online_training"
	DEFAULT_MODEL_OUTPUT_DIR    = "optimal_split_model"

	// File names
	TRAINING_FILE_NAME = "train.txt"
)

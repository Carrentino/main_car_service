package models

const (
	FUEL_TYPE_AI_92   = "92"
	FUEL_TYPE_AI_95   = "95"
	FUEL_TYPE_AI_100  = "10"
	FUEL_TYPE_GAS     = "GS"
	FUEL_TYPE_DIESEL  = "DT"
	FUEL_TYPE_ELECTRO = "EL"

	DRIVE_RWD = "RWD"
	DRIVE_FWD = "FWD"
	DRIVE_AWD = "AWD"

	GEARBOX_MANUAL    = "MA"
	GEARBOX_AUTOMATIC = "AU"
	GEARBOX_ROBOT     = "AR"
	GEARBOX_CVT       = "AC"

	BODY_TYPE_SEDAN         = "SE"
	BODY_TYPE_LIFTBACK      = "LF"
	BODY_TYPE_COUPE         = "CP"
	BODY_TYPE_HATCHBACK_3   = "H3"
	BODY_TYPE_HATCHBACK_5   = "H5"
	BODY_TYPE_STATION_WAGON = "SW"
	BODY_TYPE_SUV_3         = "S3"
	BODY_TYPE_SUV_5         = "S5"
	BODY_TYPE_MINIVAN       = "MV"
	BODY_TYPE_PICKUP        = "PC"
	BODY_TYPE_LIMOUSINE     = "LM"
	BODY_TYPE_VAN           = "VN"
	BODY_TYPE_CABRIOLET     = "CB"
)

var (
	FUEL_TYPE_CHOICES = []string{
		FUEL_TYPE_AI_92, FUEL_TYPE_AI_95, FUEL_TYPE_AI_100,
		FUEL_TYPE_GAS, FUEL_TYPE_DIESEL, FUEL_TYPE_ELECTRO,
	}

	DRIVE_CHOICES = []string{
		DRIVE_RWD, DRIVE_FWD, DRIVE_AWD,
	}

	GEARBOX_CHOICES = []string{
		GEARBOX_MANUAL, GEARBOX_AUTOMATIC, GEARBOX_ROBOT, GEARBOX_CVT,
	}

	BODY_TYPE_CHOICES = []string{
		BODY_TYPE_SEDAN, BODY_TYPE_LIFTBACK, BODY_TYPE_COUPE, BODY_TYPE_HATCHBACK_3,
		BODY_TYPE_HATCHBACK_5, BODY_TYPE_STATION_WAGON, BODY_TYPE_SUV_3,
		BODY_TYPE_SUV_5, BODY_TYPE_MINIVAN, BODY_TYPE_PICKUP, BODY_TYPE_LIMOUSINE,
		BODY_TYPE_VAN, BODY_TYPE_CABRIOLET,
	}
)
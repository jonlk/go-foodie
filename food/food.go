package food

type Food struct {
	FdcId           int
	Description     string
	PublicationDate string
	DataType        string
	FoodClass       string
	FoodComponents  []FoodComponent
	FoodAttributes  []FoodAttribute
	FoodNutrients   []FoodNutrient
	FoodPortions    []FoodPortion
	FoodCategory    FoodCategory
}

type FoodNutrient struct {
	Id                     int
	Amount                 float32
	DataPoints             int
	Type                   string
	Nutrient               Nutrient
	FoodNutrientDerivation FoodNutrientDerivation
}

type Nutrient struct {
	Id       int
	Number   string
	Name     string
	Rank     int
	UnitName string
}

type FoodNutrientDerivation struct {
	Id          int
	Code        string
	Description string
}

// TODO: Figure out what these do
type FoodAttribute struct{}
type FoodComponent struct{}

type FoodCategory struct {
	Id          int
	Code        string
	Description string
}

type FoodPortion struct {
	Id          int
	DataPoints  int
	GramWeight  float32
	Amount      float32
	Modifier    string
	MeasureUnit MeasureUnit
}

type MeasureUnit struct {
	Id           int
	Name         string
	Abbreviation string
}

type FoodDisplayModel struct {
	Name     string
	Amount   float32
	UnitName string
}

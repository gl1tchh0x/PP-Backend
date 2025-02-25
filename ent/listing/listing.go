// Code generated by ent, DO NOT EDIT.

package listing

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the listing type in the database.
	Label = "listing"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldZipCode holds the string denoting the zip_code field in the database.
	FieldZipCode = "zip_code"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldBedroom holds the string denoting the bedroom field in the database.
	FieldBedroom = "bedroom"
	// FieldBathroom holds the string denoting the bathroom field in the database.
	FieldBathroom = "bathroom"
	// FieldGarage holds the string denoting the garage field in the database.
	FieldGarage = "garage"
	// FieldSqft holds the string denoting the sqft field in the database.
	FieldSqft = "sqft"
	// FieldTypeOfProperty holds the string denoting the type_of_property field in the database.
	FieldTypeOfProperty = "type_of_property"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldLotSize holds the string denoting the lot_size field in the database.
	FieldLotSize = "lot_size"
	// FieldPool holds the string denoting the pool field in the database.
	FieldPool = "pool"
	// FieldYearBuilt holds the string denoting the year_built field in the database.
	FieldYearBuilt = "year_built"
	// FieldMedia holds the string denoting the media field in the database.
	FieldMedia = "media"
	// FieldRealtorID holds the string denoting the realtor_id field in the database.
	FieldRealtorID = "realtor_id"
	// EdgeRealtor holds the string denoting the realtor edge name in mutations.
	EdgeRealtor = "realtor"
	// Table holds the table name of the listing in the database.
	Table = "listings"
	// RealtorTable is the table that holds the realtor relation/edge.
	RealtorTable = "listings"
	// RealtorInverseTable is the table name for the Realtor entity.
	// It exists in this package in order to avoid circular dependency with the "realtor" package.
	RealtorInverseTable = "realtors"
	// RealtorColumn is the table column denoting the realtor relation/edge.
	RealtorColumn = "realtor_id"
)

// Columns holds all SQL columns for listing fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldTitle,
	FieldAddress,
	FieldCity,
	FieldState,
	FieldZipCode,
	FieldDescription,
	FieldPrice,
	FieldBedroom,
	FieldBathroom,
	FieldGarage,
	FieldSqft,
	FieldTypeOfProperty,
	FieldStatus,
	FieldLotSize,
	FieldPool,
	FieldYearBuilt,
	FieldMedia,
	FieldRealtorID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// CityValidator is a validator for the "city" field. It is called by the builders before save.
	CityValidator func(string) error
	// StateValidator is a validator for the "state" field. It is called by the builders before save.
	StateValidator func(string) error
	// ZipCodeValidator is a validator for the "zip_code" field. It is called by the builders before save.
	ZipCodeValidator func(string) error
	// BedroomValidator is a validator for the "bedroom" field. It is called by the builders before save.
	BedroomValidator func(int) error
	// BathroomValidator is a validator for the "bathroom" field. It is called by the builders before save.
	BathroomValidator func(float64) error
	// GarageValidator is a validator for the "garage" field. It is called by the builders before save.
	GarageValidator func(int) error
	// SqftValidator is a validator for the "sqft" field. It is called by the builders before save.
	SqftValidator func(int) error
	// LotSizeValidator is a validator for the "lot_size" field. It is called by the builders before save.
	LotSizeValidator func(int) error
	// YearBuiltValidator is a validator for the "year_built" field. It is called by the builders before save.
	YearBuiltValidator func(int) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// TypeOfProperty defines the type for the "type_of_property" enum field.
type TypeOfProperty string

// TypeOfPropertyHouse is the default value of the TypeOfProperty enum.
const DefaultTypeOfProperty = TypeOfPropertyHouse

// TypeOfProperty values.
const (
	TypeOfPropertyHouse     TypeOfProperty = "house"
	TypeOfPropertyApartment TypeOfProperty = "apartment"
	TypeOfPropertyCondo     TypeOfProperty = "condo"
	TypeOfPropertyTownhouse TypeOfProperty = "townhouse"
)

func (top TypeOfProperty) String() string {
	return string(top)
}

// TypeOfPropertyValidator is a validator for the "type_of_property" field enum values. It is called by the builders before save.
func TypeOfPropertyValidator(top TypeOfProperty) error {
	switch top {
	case TypeOfPropertyHouse, TypeOfPropertyApartment, TypeOfPropertyCondo, TypeOfPropertyTownhouse:
		return nil
	default:
		return fmt.Errorf("listing: invalid enum value for type_of_property field: %q", top)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// StatusDRAFT is the default value of the Status enum.
const DefaultStatus = StatusDRAFT

// Status values.
const (
	StatusDRAFT     Status = "DRAFT"
	StatusPUBLISHED Status = "PUBLISHED"
	StatusARCHIVED  Status = "ARCHIVED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusDRAFT, StatusPUBLISHED, StatusARCHIVED:
		return nil
	default:
		return fmt.Errorf("listing: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Listing queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByZipCode orders the results by the zip_code field.
func ByZipCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldZipCode, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByBedroom orders the results by the bedroom field.
func ByBedroom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBedroom, opts...).ToFunc()
}

// ByBathroom orders the results by the bathroom field.
func ByBathroom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBathroom, opts...).ToFunc()
}

// ByGarage orders the results by the garage field.
func ByGarage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGarage, opts...).ToFunc()
}

// BySqft orders the results by the sqft field.
func BySqft(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSqft, opts...).ToFunc()
}

// ByTypeOfProperty orders the results by the type_of_property field.
func ByTypeOfProperty(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTypeOfProperty, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByLotSize orders the results by the lot_size field.
func ByLotSize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLotSize, opts...).ToFunc()
}

// ByPool orders the results by the pool field.
func ByPool(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPool, opts...).ToFunc()
}

// ByYearBuilt orders the results by the year_built field.
func ByYearBuilt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYearBuilt, opts...).ToFunc()
}

// ByRealtorID orders the results by the realtor_id field.
func ByRealtorID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRealtorID, opts...).ToFunc()
}

// ByRealtorField orders the results by realtor field.
func ByRealtorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRealtorStep(), sql.OrderByField(field, opts...))
	}
}
func newRealtorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RealtorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RealtorTable, RealtorColumn),
	)
}

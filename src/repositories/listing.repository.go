package repositories

import (
	"errors"

	"backend.com/go-backend/src/config"
	"backend.com/go-backend/src/models"
	"github.com/google/uuid"
)

// CreateListingInput represents the data structure for creating a new real estate listing.
// It contains all necessary fields to create a property listing in the system.
type CreateListingInput struct {
	// Title is the display name of the listing
	Title string `json:"title"`
	// Address is the street address of the property
	Address string `json:"address"`
	// City where the property is located
	City string `json:"city"`
	// State where the property is located
	State string `json:"state"`
	// ZipCode of the property location
	ZipCode string `json:"zip_code"`
	// Description provides detailed information about the property
	Description string `json:"description"`
	// Price of the property (stored as string to handle various currency formats)
	Price string `json:"price"`
	// Bedroom indicates the number of bedrooms
	Bedroom int `json:"bedroom"`
	// Bathroom indicates the number of bathrooms (supports half baths with float)
	Bathroom float32 `json:"bathroom"`
	// Garage indicates the number of garage spaces
	Garage int `json:"garage"`
	// Sqft represents the total square footage of the property
	Sqft int64 `json:"sqft"`
	// TypeOfProperty indicates the category (e.g., single-family, condo, etc.)
	TypeOfProperty string `json:"type_of_property"`
	// LotSize represents the total land area in square feet
	LotSize int64 `json:"lot_size"`
	// Pool indicates if the property has a pool
	Pool bool `json:"pool"`
	// YearBuilt indicates when the property was constructed
	YearBuilt string `json:"year_built"`

	// PhotoMain is the primary display image URL
	PhotoMain string `json:"photo_main"`
	// Photo1 through Photo5 are additional property image URLs
	Photo1 string `json:"photo_1"`
	Photo2 string `json:"photo_2"`
	Photo3 string `json:"photo_3"`
	Photo4 string `json:"photo_4"`
	Photo5 string `json:"photo_5"`

	// IsPublished determines if the listing is visible in search results
	IsPublished bool `json:"is_published"`
	// RealtorID is the unique identifier of the realtor managing this listing
	RealtorID uuid.UUID `json:"realtor_id"`
}

// CreateListingRepository creates a new listing record in the database.
// It first checks if a listing with the given title or address already exists.
// If such a listing exists, it returns an error.
// If not, it creates a new listing record with the provided data.
// The operation is performed within a transaction to ensure atomicity.
//
// Parameters:
//   - data: CreateListingInput containing the details of the listing to be created.
//
// Returns:
//   - error: An error if the listing already exists or if the creation fails, otherwise nil.
func CreateListingRepository(data CreateListingInput) error {
	var existingListing models.Listing
	if err := config.DB.Where("title = ? OR address = ?", data.Title, data.Address).First(&existingListing).Error; err == nil {
		return errors.New("listing with the given title or address already exists")

	}

	// Create a new listing
	listing := models.Listing{
		Title:          data.Title,
		Address:        data.Address,
		City:           data.City,
		State:          data.State,
		ZipCode:        data.ZipCode,
		Description:    data.Description,
		Price:          data.Price,
		Bedroom:        data.Bedroom,
		Bathroom:       data.Bathroom,
		Garage:         data.Garage,
		Sqft:           data.Sqft,
		TypeOfProperty: data.TypeOfProperty,
		LotSize:        data.LotSize,
		Pool:           data.Pool,
		YearBuilt:      data.YearBuilt,
		PhotoMain:      data.PhotoMain,
		Photo1:         data.Photo1,
		Photo2:         data.Photo2,
		Photo3:         data.Photo3,
		Photo4:         data.Photo4,
		Photo5:         data.Photo5,
		IsPublished:    data.IsPublished,
		RealtorID:      data.RealtorID,
	}

	// Start a new transaction
	tx := config.DB.Begin()
	if err := tx.Create(&listing).Error; err != nil {
		return errors.New("failed to create listing")
	}
	tx.Commit()
	return nil
}

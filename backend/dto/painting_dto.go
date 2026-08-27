package dto

type CreatePaintingRequest struct {
	Title       string  `json:"title" binding:"required,min=2,max=200"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"description" binding:"required"`

	Category string `json:"category" binding:"required"`
	Medium   string `json:"medium" binding:"required"`

	Width  float64 `json:"width" binding:"required,gt=0"`
	Height float64 `json:"height" binding:"required,gt=0"`
	Unit   string  `json:"unit" binding:"required"`

	Featured bool   `json:"featured"`
	Status   string `json:"status" binding:"omitempty,oneof=AVAILABLE SOLD"`
	Tags     string `json:"tags"`

	ProcessVideoURL *string `json:"process_video_url,omitempty"`
}

type UpdatePaintingRequest struct {
	Title       string  `json:"title" binding:"required,min=2,max=200"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"description" binding:"required"`

	Category string `json:"category" binding:"required"`
	Medium   string `json:"medium" binding:"required"`

	Width  float64 `json:"width" binding:"required,gt=0"`
	Height float64 `json:"height" binding:"required,gt=0"`
	Unit   string  `json:"unit" binding:"required"`

	Featured bool   `json:"featured"`
	Status   string `json:"status" binding:"omitempty,oneof=AVAILABLE SOLD"`
	Tags     string `json:"tags"`

	ProcessVideoURL *string `json:"process_video_url,omitempty"`
}

type PaintingResponse struct {
	ID uint `json:"id"`

	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`

	Category string `json:"category"`
	Medium   string `json:"medium"`

	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Unit   string  `json:"unit"`

	Featured bool   `json:"featured"`
	Status   string `json:"status"`
	Tags     string `json:"tags"`

	ProcessVideoURL *string `json:"process_video_url,omitempty"`

	Views uint `json:"views"`

	Images []PaintingImageResponse `json:"images"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateFeaturedRequest struct {
	Featured bool `json:"featured"`
}

type UpdatePaintingStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=AVAILABLE SOLD"`
}

type PaintingImageResponse struct {
	ID        uint   `json:"id"`
	ImageURL  string `json:"image_url"`
	CreatedAt string `json:"created_at"`
}

type AdminPaintingResponse struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title"`
	Price           float64 `json:"price"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	Medium          string  `json:"medium"`
	Width           float64 `json:"width"`
	Height          float64 `json:"height"`
	Unit            string  `json:"unit"`
	Featured        bool    `json:"featured"`
	Status          string  `json:"status"`
	ProcessVideoURL *string `json:"process_video_url,omitempty"`
	Views           uint    `json:"views"`
	Tags            string  `json:"tags"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`

	// Sales information
	Sold            bool    `json:"sold"`
	SoldOrderID     *uint   `json:"sold_order_id,omitempty"`
	SoldPrice       *float64 `json:"sold_price,omitempty"`
}

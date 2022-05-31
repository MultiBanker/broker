package models

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/hashicorp/go-multierror"
)

var (
	ErrInvalidURL  = errors.New("invalid url")
	ErrInvalidType = errors.New("invalid media type")
)

type MediaType string

const (
	Image MediaType = "image"
	Video MediaType = "video"
)

type Media struct {
	Type MediaType `json:"type" bson:"type"`
	URL  string    `json:"url" bson:"url"`
}

type Medias []Media

func (m Medias) Image() string {
	for i := range m {
		if m[i].Type == Image {
			return m[i].URL
		}
	}
	return ""
}

func (m Medias) Validate() error {
	var result *multierror.Error

	for _, media := range m {
		if _, err := url.ParseRequestURI(media.URL); err != nil {
			result = multierror.Append(result, fmt.Errorf("media: %w", err))
		}

		u, err := url.Parse(media.URL)
		if err != nil {
			result = multierror.Append(result, fmt.Errorf("media: %w", err))
		}

		if u.Scheme == "" || u.Host == "" {
			result = multierror.Append(result, fmt.Errorf("media: %v", ErrInvalidURL))
		}

	}

	return result.ErrorOrNil()
}

type Auto struct {
	ID    string
	SKU   string

	Title LangOptions `json:"title,omitempty" bson:"title" validate:"required"`
	Brand Brand       `json:"brand" bson:"brand"`
	Color Color       `json:"color,omitempty" bson:"color"`
	Media Medias      `json:"media,omitempty" bson:"media"`
	About LangOptions `json:"about,omitempty" bson:"about,omitempty"`
	Price Price       `json:"price" bson:"price"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
type Price struct {
	Current int `json:"current" bson:"current"`
	Old     int `json:"old" bson:"old"`
	// Minimum is the minimum possible price for this product. Used for validation of merchant prices.
	Minimum int `json:"minimum" bson:"minimum"`
}

func (p Price) Validate() error {
	var result *multierror.Error

	if p.Current == 0 {
		result = multierror.Append(result, errors.New("price is zero"))
	}

	return result.ErrorOrNil()
}

type Color struct {
	Slug  string      `json:"slug" bson:"slug"`
	Title LangOptions `json:"title" bson:"title"`
	Type  int         `json:"type" bson:"type"`
	HEX   string      `json:"hex" bson:"hex"`
}

func (c Color) Validate() error {
	var result *multierror.Error

	if c.Slug == "" {
		result = multierror.Append(result, errors.New("slug empty"))
	}

	if c.HEX == "" {
		result = multierror.Append(result, errors.New("hex empty"))
	}

	if err := c.Title.Validate(); err != nil {
		result = multierror.Append(result, fmt.Errorf("color: %w", err))
	}

	return result
}

type LangOptions struct {
	KZ string `json:"kz" bson:"kz"`
	RU string `json:"ru" bson:"ru"`
}

func (lo LangOptions) Validate() error {
	var result *multierror.Error

	if lo.KZ == "" {
		result = multierror.Append(result, errors.New("empty kz lang"))
	}
	if lo.RU == "" {
		result = multierror.Append(result, errors.New("empty ru lang"))
	}

	return result.ErrorOrNil()
}

type Brand struct {
	Title     LangOptions `json:"title" bson:"title"`
	Slug      string      `json:"slug" bson:"slug"`
	Equipment string      `json:"equipment" bson:"equipment"`
}

func (b Brand) Validate() error {
	var result *multierror.Error

	if b.Slug == "" {
		result = multierror.Append(result, errors.New("slug empty"))
	}

	if b.Equipment == "" {
		result = multierror.Append(result, errors.New("hex empty"))
	}

	if err := b.Title.Validate(); err != nil {
		result = multierror.Append(result, fmt.Errorf("brand: %w", err))
	}

	return result
}

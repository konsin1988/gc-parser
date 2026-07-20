package parser

import (
	"regexp"
	"fmt"

	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/model"
)

type productHeadingWidget struct {
    Title string `json:"title"`
}

type priceWidget struct {
	Price 					string		`json:"price"`
	CardPrice 			string		`json:"cardPrice"`
	OriginalPrice 	string		`json:"originalPrice"`
	Availability		bool 			`json:"isAvailable"`
}

type breadcrumbsWidget struct {
	Breadcrumbs []struct {
		Text string `json:"text"`
		Link string `json:"link"`
	} `json:"breadcrumbs"`
}


type brandWidget struct {
	Content struct {
		Link string `json:"link"`

		Title struct {
			Text []struct {
				Content string `json:"content"`
			} `json:"text"`
		} `json:"title"`
	} `json:"content"`
}

type imageWidget struct {
	CoverImage string `json:"coverImage"`

	Images []struct {
		Src string `json:"src"`
	} `json:"images"`
}

type reviewLinkWidget struct {
	Link string `json:"link"`
}


type sellerWidget struct {
	Header struct {
		Badge struct {
			Subscribed struct {
				Common struct {
					Action struct {
						Params struct {
							SellerID string `json:"sellerId"`
						} `json:"params"`
					} `json:"action"`
				} `json:"common"`
			} `json:"subscribed"`
		} `json:"badge"`
	} `json:"header"`

	SellerCell struct {
		Common struct {
			Action struct {
				Link string `json:"link"`
			} `json:"action"`
		} `json:"common"`
	} `json:"sellerCell"`

	TrustFactors []struct {
		Title struct {
			Text string `json:"text"`
		} `json:"title"`

		Tooltip struct {
			Subtitle []struct {
				Content string `json:"content"`
			} `json:"subtitle"`
		} `json:"tooltip"`
	} `json:"trustFactors"`
}



func ParseGoodItem (page *ozon.PageResponse) (*model.ParsedGoodItem, error) {
	result := &model.ParsedGoodItem{}
	result.Sku = page.PageInfo.AnalyticsInfo.Sku.String()

	// ------------------------------------------------------------------------ TITLE 
	var heading productHeadingWidget
	if err := ParseWidget(page, "webProductHeading-", &heading); err != nil {
	    return nil, err
	}
	result.Title = heading.Title

	// --------------------------------------------------------------- PRICE 
	var prices priceWidget
	if err := ParseWidget(page, "webPrice-", &prices); err != nil {
	    return nil, err
	}

	var err error
	result.Price, err = parsePrice(prices.Price)
	if err != nil {
		return nil, err
	}
	result.CardPrice, err = parsePrice(prices.CardPrice)
	if err != nil {
		return nil, err
	}
	result.OriginalPrice, err = parsePrice(prices.OriginalPrice)
	if err != nil {
		return nil, err
	}
	result.Availability = 	prices.Availability

	// ----------------------------------------------------------- CATEGORIES 
	var crumbs breadcrumbsWidget
	
	if err := ParseWidget(page, "breadCrumbs", &crumbs); err != nil {
		return nil, err
	}
	result.Categories = make([]model.Category, 0, len(crumbs.Breadcrumbs))
	
	for _, c := range crumbs.Breadcrumbs {
		result.Categories = append(result.Categories, model.Category{
			Name: c.Text,
			Slug: normalizeSlug(c.Link, "/category"),
		})
	}

	// ----------------------------------------------------------- BRAND 
	var brandwidget brandWidget
	if err := ParseWidget(page, "webBrand-", &brandwidget); err != nil {
		return nil, err
	}

  key, err := FindWidgetKey(page, "webBrand-")
  if err != nil {
      return nil, err
  }

  raw, ok := page.WidgetStates[key]
  if !ok {
      return nil, fmt.Errorf("widget %q not found", key)
  }
	brand := &model.Brand{}
	if len(brandwidget.Content.Title.Text) > 0 {
		brand.Title = brandwidget.Content.Title.Text[0].Content
	}

	re := regexp.MustCompile(`/brand/(.+?)/\?all_items=true`)
	match := re.FindStringSubmatch(string(raw))

	if len(match) > 1 {
		brand.Slug = match[1]
	}

	result.Brand = brand 

	// -------------------------------------------------------------- IMAGES
	var gallery imageWidget
	
	if err := ParseWidget(page, "webGallery-", &gallery); err != nil {
		return nil, err
	}
	result.Images = make([]model.Image, 0, len(gallery.Images)+1)
	seen := map[string]struct{}{}


	if gallery.CoverImage != "" {
		result.Images = append(result.Images, model.Image{
			Sku:     result.Sku,
			ImgURL:  gallery.CoverImage,
			IsCover: true,
		})
		seen[gallery.CoverImage] = struct{}{}
	}

	for _, img := range gallery.Images {
		if _, ok := seen[img.Src]; ok {
			continue
		}
		result.Images = append(result.Images, model.Image{
			Sku:     result.Sku,
			ImgURL:  img.Src,
			IsCover: false,
		})
		seen[img.Src] = struct{}{}
	}

	// ----------------------------------------------------------- REVIEW LINK
	var score reviewLinkWidget
	
	if err := ParseWidget(page, "webSingleProductScore-", &score); err != nil {
		return nil, err
	}
	result.ReviewLink = score.Link


	// ------------------------------------------------------------ SELLER 
	var seller sellerWidget
	
	if err := ParseWidget(page, "webCurrentSeller-", &seller); err != nil {
		return nil, err
	}

	result.Seller = &model.ParsedSeller{
		ID: seller.Header.Badge.Subscribed.Common.Action.Params.SellerID,
		Slug: parseSellerSlug(
			seller.SellerCell.Common.Action.Link,
		),
		OGRNIP: parseOGRNIP(seller.TrustFactors),
	}
	
	for _, factor := range seller.TrustFactors {
		if factor.Title.Text == "О магазине" &&
			len(factor.Tooltip.Subtitle) > 0 {
	
			result.Seller.Name = factor.Tooltip.Subtitle[0].Content
			break
		}
	}


	return result, nil
}

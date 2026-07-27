package parser

import (
	"path"
	"strings"

	"konsin1988/gc-agent/model"
)


// ======================================================== ParseBreadcrumbs
func trimSlug(link string) string {
    return strings.Trim(link, "/")
}

func ParseBreadcrumbs(
		crumbs []model.Category,
) ([]model.Category, *model.Brand) {

    if len(crumbs) == 0 {
        return nil, nil
    }

    lastSlug := trimSlug(crumbs[len(crumbs)-1].Slug)

    switch {

    case strings.HasPrefix(lastSlug, "brand/"):
        return parseBrandBreadcrumbs(crumbs)

    case strings.HasPrefix(lastSlug, "seller/"):
        return parseSellerBreadcrumbs(crumbs)

    case strings.HasPrefix(lastSlug, "category/"):
        return parseCategoryBreadcrumbs(crumbs)

    default:
        return nil, nil
    }
}

func parseBrandBreadcrumbs(
		crumbs  []model.Category,
) ([]model.Category, *model.Brand) {

    brand := &model.Brand{
        Title: crumbs[0].Name,
        Slug:  path.Base(trimSlug(crumbs[0].Slug)),
    }

    categories := make([]model.Category, 0, len(crumbs)-1)

    for _, c := range crumbs[1:] {
        categories = append(categories, model.Category{
            Name: c.Name,
            Slug: normalizeBrandSlug(c.Slug),
        })
    }

		return categories, brand
}

func normalizeBrandSlug(link string) string{
	slug := strings.Trim(link, "/")
	parts := strings.Split(slug, "/")
	return parts[len(parts)-1]
}


// -----------------------------------------
func parseSellerBreadcrumbs(
    crumbs []model.Category,
) ([]model.Category, *model.Brand) {

    categories := make([]model.Category, 0, len(crumbs)-1)

    for _, c := range crumbs[1:] {
        categories = append(categories, model.Category{
            Name: c.Name,
            Slug: path.Base(trimSlug(c.Slug)),
        })
    }

    return categories, nil
}


func parseCategoryBreadcrumbs(
    crumbs  []model.Category,
) ([]model.Category, *model.Brand) {

    categories := make([]model.Category, 0, len(crumbs))

    for _, c := range crumbs {
        slug := strings.TrimPrefix(
            trimSlug(c.Slug),
            "category/",
        )

        categories = append(categories, model.Category{
            Name: c.Name,
            Slug: slug,
        })
    }

    if len(categories) == 0 {
        return categories, nil
    }

    last := categories[len(categories)-1]

    if strings.Contains(last.Slug, "/") {

        brand := &model.Brand{
            Title: last.Name,
            Slug:  path.Base(last.Slug),
        }

        return categories[:len(categories)-1], brand
    }

    return categories, nil
}

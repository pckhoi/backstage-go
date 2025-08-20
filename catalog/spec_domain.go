package catalog

// DomainSpec contains the Domain standard spec fields.
//
// See: https://backstage.io/docs/features/software-catalog/descriptor-format#kind-domain
type DomainSpec struct {
	// An entity reference to the owner of the domain. This field is required.
	Owner string `json:"owner"`
	// An entity reference to another domain of which the domain is a part of.
	SubdomainOf string `json:"subdomainOf,omitempty"`
	// The type of domain.
	Type string `json:"type,omitempty"`
}

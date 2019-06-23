package square

type Address struct {
	AddressLine1                 string `json:"address_line_1"`
	AddressLine2                 string `json:"address_line_2"`
	AddressLine3                 string `json:"address_line_3"`
	Locality                     string `json:"locality"`
	Sublocality                  string `json:"sublocality"`
	Sublocality2                 string `json:"sublocality_2"`
	Sublocality3                 string `json:"sublocality_3"`
	AdministrativeDistrictLevel1 string `json:"administrative_district_level_1"`
	AdministrativeDistrictLevel2 string `json:"administrative_district_level_2"`
	AdministrativeDistrictLevel3 string `json:"administrative_district_level_3"`
	PostalCode                   string `json:"postal_code"`
	Country                      string `json:"country"`
	FirstName                    string `json:"first_name"`
	LastName                     string `json:"last_name"`
	Organization                 string `json:"organization"`
}

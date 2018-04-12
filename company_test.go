package uri

import "testing"

func TestGetCompany(t *testing.T) {
	tt := []struct {
		companyNumber string
		companyName   string
	}{
		{"05581537", "OPEN RIGHTS"},
	}

	for _, tc := range tt {
		t.Run(tc.companyNumber, func(t *testing.T) {
			company, err := GetCompany(tc.companyNumber)
			if err != nil {
				t.Fatalf("Expected to pass, but got %v", err)
			}

			if company.Name != tc.companyName {
				t.Errorf("Expected company name to be %s. but got %s", tc.companyName, company.Name)
			}

			t.Logf("%v", company)
		})
	}
}

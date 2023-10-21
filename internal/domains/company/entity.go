package company

type (
	Company struct {
		CompanyId    int64
		Industry     int64
		CompanyScale int64
		CompanyValue int64
		CompanyName  string
		CompanyDesc  string
	}
	HeadUser struct {
		HeadUserID  int64
		FistName    string
		LastName    string
		PhoneNumber string
	}

	Register struct {
		HeadUser
		Company
	}
)

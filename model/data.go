package model

type ErrorResponse struct {
	Errors []struct {
		Code   string `json:"code"`
		Status string `json:"status"`
		Id     string `json:"id"`
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Source string `json:"source"`
	} `json:"errors"`
}

type Certificates struct {
	Certificate []struct {
		Attributes struct {
			CertificateContent string `json:"certificateContent"`
			DisplayName        string `json:"displayName"`
			ExpirationDate     string `json:"expirationDate"`
			Name               string `json:"name"`
			Platform           string `json:"platform"`
			SerialNumber       string `json:"serialNumber"`
			CertificateType    string `json:"certificateType"`
		} `json:"attributes"`
		Id    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`

	PagedDocumentLinks struct {
		First string `json:"first"`
		Next  string `json:"next"`
		Self  string `json:"self"`
	} `json:"links"`

	PagingInformation struct {
		Paging struct {
			Total int `json:"total"`
			Limit int `json:"limit"`
		} `json:"paging"`
	} `json:"meta"`
}

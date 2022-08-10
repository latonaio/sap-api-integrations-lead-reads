package responses

type ToLeadItemCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                   string `json:"ObjectID"`
			ParentObjectID             string `json:"ParentObjectID"`
			LeadID                     string `json:"LeadID"`
			ID                         string `json:"ID"`
			UUID                       string `json:"UUID"`
			Description                string `json:"Description"`
			LanguageCode               string `json:"languageCode"`
			LanguageCodeText           string `json:"languageCodeText"`
			ProductID                  string `json:"ProductID"`
			ProductUUID                string `json:"ProductUUID"`
			ProductName                string `json:"ProductName"`
			ProductCategoryInternalID  string `json:"ProductCategoryInternalID"`
			Quantity                   string `json:"Quantity"`
			UnitCode                   string `json:"unitCode"`
			UnitCodeText               string `json:"unitCodeText"`
			ProductCategoryDescription string `json:"ProductCategoryDescription"`
			EntityLastChangedOn        string `json:"EntityLastChangedOn"`
			ETag                       string `json:"ETag"`
			Lead                       struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"Lead"`
		} `json:"results"`
	} `json:"d"`
}

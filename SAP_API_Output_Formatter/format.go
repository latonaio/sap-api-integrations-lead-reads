package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-lead-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToLeadCollection(raw []byte, l *logger.Logger) ([]LeadCollection, error) {
	pm := &responses.LeadCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to LeadCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	leadCollection := make([]LeadCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		leadCollection = append(leadCollection, LeadCollection{
			ObjectID:                                 data.ObjectID,
			ID:                                       data.ID,
			Name:                                     data.Name,
			NameLanguageCode:                         data.NameLanguageCode,
			AccountPartyID:                           data.AccountPartyID,
			AccountPartyUUID:                         data.AccountPartyUUID,
			AccountPartyName:                         data.AccountPartyName,
			ContactID:                                data.ContactID,
			ContactUUID:                              data.ContactUUID,
			ContactName:                              data.ContactName,
			Company:                                  data.Company,
			ContactFirstName:                         data.ContactFirstName,
			ContactMiddleName:                        data.ContactMiddleName,
			ContactLastName:                          data.ContactLastName,
			IndividualCustomerGivenName:              data.IndividualCustomerGivenName,
			IndividualCustomerMiddleName:             data.IndividualCustomerMiddleName,
			IndividualCustomerFamilyName:             data.IndividualCustomerFamilyName,
			QualificationLevelCode:                   data.QualificationLevelCode,
			UserStatusCode:                           data.UserStatusCode,
			ResultReasonCode:                         data.ResultReasonCode,
			ApprovalStatusCode:                       data.ApprovalStatusCode,
			ConsistencyStatusCode:                    data.ConsistencyStatusCode,
			OriginTypeCode:                           data.OriginTypeCode,
			PriorityCode:                             data.PriorityCode,
			StartDate:                                data.StartDate,
			EndDate:                                  data.EndDate,
			CampaignID:                               data.CampaignID,
			GroupCode:                                data.GroupCode,
			OwnerPartyID:                             data.OwnerPartyID,
			OwnerPartyUUID:                           data.OwnerPartyUUID,
			EmployeeResponsibleUUID:                  data.EmployeeResponsibleUUID,
			OwnerPartyName:                           data.OwnerPartyName,
			OwnerPartyIDSales:                        data.OwnerPartyIDSales,
			OwnerUUIDSales:                           data.OwnerUUIDSales,
			OwnerSalesName:                           data.OwnerSalesName,
			MarketingUnitPartyID:                     data.MarketingUnitPartyID,
			MarketingUnitPartyUUID:                   data.MarketingUnitPartyUUID,
			MarketingName:                            data.MarketingName,
			SalesUnitPartyID:                         data.SalesUnitPartyID,
			SalesUnitPartyUUID:                       data.SalesUnitPartyUUID,
			SalesUnitName:                            data.SalesUnitName,
			SalesOrganisationID:                      data.SalesOrganisationID,
			SalesOrganisationUUID:                    data.SalesOrganisationUUID,
			DistributionChannelCode:                  data.DistributionChannelCode,
			DivisionCode:                             data.DivisionCode,
			SalesOfficeID:                            data.SalesOfficeID,
			SalesOfficeUUID:                          data.SalesOfficeUUID,
			SalesGroupID:                             data.SalesGroupID,
			SalesGroupUUID:                           data.SalesGroupUUID,
			SurveyTotalScoreValue:                    data.SurveyTotalScoreValue,
			CreationDateTime:                         data.CreationDateTime,
			LastChangeDateTime:                       data.LastChangeDateTime,
			CreationIdentityUUID:                     data.CreationIdentityUUID,
			LastChangeIdentityUUID:                   data.LastChangeIdentityUUID,
			Note:                                     data.Note,
			SalesTerritoryID:                         data.SalesTerritoryID,
			SalesTerritoryUUID:                       data.SalesTerritoryUUID,
			SalesTerritoryName:                       data.SalesTerritoryName,
			ExpectedRevenueAmount:                    data.ExpectedRevenueAmount,
			ExpectedRevenueCurrencyCode:              data.ExpectedRevenueCurrencyCode,
			Score:                                    data.Score,
			CompanySecondName:                        data.CompanySecondName,
			CompanyThirdName:                         data.CompanyThirdName,
			CompanyFourthName:                        data.CompanyFourthName,
			AccountPostalAddressElementsHouseID:      data.AccountPostalAddressElementsHouseID,
			AccountPostalAddressElementsStreetPrefix: data.AccountPostalAddressElementsStreetPrefix,
			AccountPostalAddressElementsAdditionalStreetPrefixName: data.AccountPostalAddressElementsAdditionalStreetPrefixName,
			AccountPostalAddressElementsStreetName:                 data.AccountPostalAddressElementsStreetName,
			AccountPostalAddressElementsStreetSufix:                data.AccountPostalAddressElementsStreetSufix,
			AccountPostalAddressElementsAdditionalStreetSuffixName: data.AccountPostalAddressElementsAdditionalStreetSuffixName,
			AccountCity:                         data.AccountCity,
			AccountCountry:                      data.AccountCountry,
			AccountState:                        data.AccountState,
			AccountPostalAddressElementsPOBoxID: data.AccountPostalAddressElementsPOBoxID,
			AccountPostalAddressElementsStreetPostalCode: data.AccountPostalAddressElementsStreetPostalCode,
			AccountPostalAddressElementsCountyName:       data.AccountPostalAddressElementsCountyName,
			AccountPhone:                                 data.AccountPhone,
			AccountFax:                                   data.AccountFax,
			AccountMobile:                                data.AccountMobile,
			AccountEMail:                                 data.AccountEMail,
			AccountWebsite:                               data.AccountWebsite,
			AccountLatitudeMeasure:                       data.AccountLatitudeMeasure,
			AccountLatitudeMeasureUnitCode:               data.AccountLatitudeMeasureUnitCode,
			AccountLongitudeMeasure:                      data.AccountLongitudeMeasure,
			AccountLongitudeMeasureUnitCode:              data.AccountLongitudeMeasureUnitCode,
			AccountLegalForm:                             data.AccountLegalForm,
			OrganisationAccountABCClassificationCode:     data.OrganisationAccountABCClassificationCode,
			OrganisationAccountIndustrialSectorCode:      data.OrganisationAccountIndustrialSectorCode,
			AccountDUNS:                                  data.AccountDUNS,
			OrganisationAccountContactAllowedCode:        data.OrganisationAccountContactAllowedCode,
			AccountCorrespondenceLanguageCode:            data.AccountCorrespondenceLanguageCode,
			AccountNote:                                  data.AccountNote,
			ContactFormOfAddressCode:                     data.ContactFormOfAddressCode,
			ContactFunctionalTitleName:                   data.ContactFunctionalTitleName,
			ContactAcademicTitleCode:                     data.ContactAcademicTitleCode,
			ContactAdditionalAcademicTitleCode:           data.ContactAdditionalAcademicTitleCode,
			ContactNickName:                              data.ContactNickName,
			ContactCorrespondenceLanguageCode:            data.ContactCorrespondenceLanguageCode,
			ContactGenderCode:                            data.ContactGenderCode,
			ContactMaritalStatusCode:                     data.ContactMaritalStatusCode,
			BusinessPartnerRelationshipBusinessPartnerFunctionTypeCode:   data.BusinessPartnerRelationshipBusinessPartnerFunctionTypeCode,
			BusinessPartnerRelationshipBusinessPartnerFunctionalAreaCode: data.BusinessPartnerRelationshipBusinessPartnerFunctionalAreaCode,
			ContactDepartmentName:                                 data.ContactDepartmentName,
			BusinessPartnerRelationshipContactVIPReasonCode:       data.BusinessPartnerRelationshipContactVIPReasonCode,
			ContactAllowedCode:                                    data.ContactAllowedCode,
			BusinessPartnerRelationshipEngagementScoreNumberValue: data.BusinessPartnerRelationshipEngagementScoreNumberValue,
			ContactBuildingID:                                     data.ContactBuildingID,
			ContactFloorID:                                        data.ContactFloorID,
			ContactRoomID:                                         data.ContactRoomID,
			ContactPhone:                                          data.ContactPhone,
			ContactFacsimileFormattedNumberDescription:            data.ContactFacsimileFormattedNumberDescription,
			ContactMobile:                                         data.ContactMobile,
			ContactEMail:                                          data.ContactEMail,
			ContactEMailUsageDeniedIndicator:                      data.ContactEMailUsageDeniedIndicator,
			ContactNote:                                           data.ContactNote,
			IndividualCustomerABCClassificationCode:               data.IndividualCustomerABCClassificationCode,
			IndividualCustomerGenderCode:                          data.IndividualCustomerGenderCode,
			IndividualCustomerMaritalStatusCode:                   data.IndividualCustomerMaritalStatusCode,
			IndividualCustomerEMail:                               data.IndividualCustomerEMail,
			IndividualCustomerPhone:                               data.IndividualCustomerPhone,
			IndividualCustomerFax:                                 data.IndividualCustomerFax,
			IndividualCustomerMobile:                              data.IndividualCustomerMobile,
			IndividualCustomerAddressHouseID:                      data.IndividualCustomerAddressHouseID,
			IndividualCustomerAddressStreetName:                   data.IndividualCustomerAddressStreetName,
			IndividualCustomerAddressCity:                         data.IndividualCustomerAddressCity,
			IndividualCustomerAddressCountry:                      data.IndividualCustomerAddressCountry,
			IndividualCustomerAddressState:                        data.IndividualCustomerAddressState,
			IndividualCustomerAddressPostalCode:                   data.IndividualCustomerAddressPostalCode,
			IndividualCustomerAddressCountyName:                   data.IndividualCustomerAddressCountyName,
			IndividualCustomerNationalityCountryCode:              data.IndividualCustomerNationalityCountryCode,
			IndividualCustomerBirthDate:                           data.IndividualCustomerBirthDate,
			IndividualCustomerFormOfAddressCode:                   data.IndividualCustomerFormOfAddressCode,
			IndividualCustomerAcademicTitleCode:                   data.IndividualCustomerAcademicTitleCode,
			IndividualCustomerOccupationCode:                      data.IndividualCustomerOccupationCode,
			IndividualCustomerContactAllowedCode:                  data.IndividualCustomerContactAllowedCode,
			IndividualCustomerNonVerbalCommunicationLanguageCode:  data.IndividualCustomerNonVerbalCommunicationLanguageCode,
			IndividualCustomerInitialsName:                        data.IndividualCustomerInitialsName,
			AccountPreferredCommunicationMediumTypeCode:           data.AccountPreferredCommunicationMediumTypeCode,
			IndividualCustomerNamePrefixCode:                      data.IndividualCustomerNamePrefixCode,
			AccountLifeCycleStatusCode:                            data.AccountLifeCycleStatusCode,
			MainContactLifeCycleStatusCode:                        data.MainContactLifeCycleStatusCode,
			LifeCycleStatusCode:                                   data.LifeCycleStatusCode,
			ExternalID:                                            data.ExternalID,
			CreatedBy:                                             data.CreatedBy,
			LastChangedBy:                                         data.LastChangedBy,
			EntityLastChangedOn:                                   data.EntityLastChangedOn,
			ETag:                                                  data.ETag,
			ToLeadItemCollection:                                  data.LeadItem.Deferred.URI,
		})
	}

	return leadCollection, nil
}

func ConvertToLeadItemCollection(raw []byte, l *logger.Logger) ([]LeadItemCollection, error) {
	pm := &responses.ToLeadItemCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToLeadItemCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	leadItemCollection := make([]LeadItemCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		leadItemCollection = append(leadItemCollection, LeadItemCollection{
			ObjectID:                   data.ObjectID,
			ParentObjectID:             data.ParentObjectID,
			LeadID:                     data.LeadID,
			ID:                         data.ID,
			UUID:                       data.UUID,
			Description:                data.Description,
			LanguageCode:               data.LanguageCode,
			LanguageCodeText:           data.LanguageCodeText,
			ProductID:                  data.ProductID,
			ProductUUID:                data.ProductUUID,
			ProductName:                data.ProductName,
			ProductCategoryInternalID:  data.ProductCategoryInternalID,
			Quantity:                   data.Quantity,
			UnitCode:                   data.UnitCode,
			UnitCodeText:               data.UnitCodeText,
			ProductCategoryDescription: data.ProductCategoryDescription,
			EntityLastChangedOn:        data.EntityLastChangedOn,
			ETag:                       data.ETag,
		})
	}

	return leadItemCollection, nil
}
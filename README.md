# sap-api-integrations-lead-reads  
sap-api-integrations-lead-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API リードデータを取得するマイクロサービスです。  
sap-api-integrations-lead-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-lead-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/lead/overview  

## 動作環境
sap-api-integrations-lead-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-lead-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-lead-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/lead/overview
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-lead-reads には、次の API をコールするためのリソースが含まれています。  

* LeadCollection（リード - リード）※リードの詳細データを取得するために、ToLeadItemCollection、と合わせて利用されます。  
* ToLeadItemCollection（リード - リード明細 ※To）


## API への 値入力条件 の 初期値
sap-api-integrations-lead-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.LeadCollection.ID（ID）


## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"LeadCollection" が指定されています。    
  
```
	"api_schema": "Lead",
	"accepter": ["LeadCollection"],
	"lead_code": "25779",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "Lead",
	"accepter": ["All"],
	"lead_code": "25779",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetLead(iD, leadID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "LeadCollection":
			func() {
				c.LeadCollection(iD)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP の リードデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "ETag" は、/SAP_API_Output_Formatter/type.go 内 の Type LeadCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-lead-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-lead-reads/SAP_API_Caller.(*SAPAPICaller).LeadCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E0C6CDB1ED7A1B004E3242D49E9",
			"ID": "25782",
			"Name": "Vacation Offers group 'B'",
			"NameLanguageCode": "EN",
			"AccountPartyID": "1001562",
			"AccountPartyUUID": "00163E09-46E6-1ED4-9ACC-887764EF755C",
			"AccountPartyName": "Kixo Pvt Ltd",
			"ContactID": "1001561",
			"ContactUUID": "00163E09-46E6-1ED4-9ACC-887764F0B55C",
			"ContactName": "Duncan Williams",
			"Company": "Kixo Pvt Ltd",
			"ContactFirstName": "Duncan",
			"ContactMiddleName": "",
			"ContactLastName": "Williams",
			"IndividualCustomerGivenName": "",
			"IndividualCustomerMiddleName": "",
			"IndividualCustomerFamilyName": "",
			"QualificationLevelCode": "",
			"UserStatusCode": "03",
			"ResultReasonCode": "",
			"ApprovalStatusCode": "1",
			"ConsistencyStatusCode": "3",
			"OriginTypeCode": "003",
			"PriorityCode": "3",
			"StartDate": "2017-08-20T09:00:00+09:00",
			"EndDate": "2017-08-27T09:00:00+09:00",
			"CampaignID": "",
			"GroupCode": "",
			"OwnerPartyID": "8000000001",
			"OwnerPartyUUID": "00163E03-A070-1EE2-88BA-37FB8F5F50A9",
			"EmployeeResponsibleUUID": "00163E03-A070-1EE2-88BA-37FB8F5F50A9",
			"OwnerPartyName": "Mike Summers",
			"OwnerPartyIDSales": "",
			"OwnerUUIDSales": "",
			"OwnerSalesName": "",
			"MarketingUnitPartyID": "US1100",
			"MarketingUnitPartyUUID": "00163E03-A070-1EE2-88B8-D4A36EBC8D9B",
			"MarketingName": "Sales Unit US",
			"SalesUnitPartyID": "US1100",
			"SalesUnitPartyUUID": "00163E03-A070-1EE2-88B8-D4A36EBC8D9B",
			"SalesUnitName": "Sales Unit US",
			"SalesOrganisationID": "2000",
			"SalesOrganisationUUID": "00163E06-64D3-1EE3-B1AA-93C1DE920DDC",
			"DistributionChannelCode": "01",
			"DivisionCode": "00",
			"SalesOfficeID": "",
			"SalesOfficeUUID": "",
			"SalesGroupID": "",
			"SalesGroupUUID": "",
			"SurveyTotalScoreValue": 0,
			"CreationDateTime": "2017-08-20T17:18:32+09:00",
			"LastChangeDateTime": "2017-09-07T17:57:21+09:00",
			"CreationIdentityUUID": "00163E0C-6CDB-1EE6-B8AC-E924D02614E4",
			"LastChangeIdentityUUID": "00163E0C-84E2-1ED7-80DF-DF0660BB2577",
			"Note": "",
			"SalesTerritoryID": "2",
			"SalesTerritoryUUID": "00163E03-A070-1ED2-8B9D-F1A6706DD0A9",
			"SalesTerritoryName": "United States",
			"ExpectedRevenueAmount": "0.000000",
			"ExpectedRevenueCurrencyCode": "",
			"Score": 0,
			"CompanySecondName": "",
			"CompanyThirdName": "",
			"CompanyFourthName": "",
			"AccountPostalAddressElementsHouseID": "901",
			"AccountPostalAddressElementsStreetPrefix": "",
			"AccountPostalAddressElementsAdditionalStreetPrefixName": "",
			"AccountPostalAddressElementsStreetName": "E California Ave",
			"AccountPostalAddressElementsStreetSufix": "",
			"AccountPostalAddressElementsAdditionalStreetSuffixName": "",
			"AccountCity": "Sunnyvale",
			"AccountCountry": "US",
			"AccountState": "",
			"AccountPostalAddressElementsPOBoxID": "",
			"AccountPostalAddressElementsStreetPostalCode": "94085-4503",
			"AccountPostalAddressElementsCountyName": "",
			"AccountPhone": "+1 9738814948",
			"AccountFax": "+1 9738814730",
			"AccountMobile": "+1 738-814-999",
			"AccountEMail": "",
			"AccountWebsite": "www.kixo.com",
			"AccountLatitudeMeasure": "0.00000000000000",
			"AccountLatitudeMeasureUnitCode": "",
			"AccountLongitudeMeasure": "0.00000000000000",
			"AccountLongitudeMeasureUnitCode": "",
			"AccountLegalForm": "",
			"OrganisationAccountABCClassificationCode": "6",
			"OrganisationAccountIndustrialSectorCode": "",
			"AccountDUNS": "",
			"OrganisationAccountContactAllowedCode": "",
			"AccountCorrespondenceLanguageCode": "EN",
			"AccountNote": "",
			"ContactFormOfAddressCode": "",
			"ContactFunctionalTitleName": "VP Purchasing",
			"ContactAcademicTitleCode": "",
			"ContactAdditionalAcademicTitleCode": "",
			"ContactNickName": "",
			"ContactCorrespondenceLanguageCode": "EN",
			"ContactGenderCode": "1",
			"ContactMaritalStatusCode": "",
			"BusinessPartnerRelationshipBusinessPartnerFunctionTypeCode": "",
			"BusinessPartnerRelationshipBusinessPartnerFunctionalAreaCode": "0001",
			"ContactDepartmentName": "",
			"BusinessPartnerRelationshipContactVIPReasonCode": "",
			"ContactAllowedCode": "",
			"BusinessPartnerRelationshipEngagementScoreNumberValue": "0 ",
			"ContactBuildingID": "",
			"ContactFloorID": "",
			"ContactRoomID": "",
			"ContactPhone": "+1 650 464 7898",
			"ContactFacsimileFormattedNumberDescription": "",
			"ContactMobile": "+1 650-464-7880",
			"ContactEMail": "duncankixo@gmail.com",
			"ContactEMailUsageDeniedIndicator": "",
			"ContactNote": "",
			"IndividualCustomerABCClassificationCode": "",
			"IndividualCustomerGenderCode": "",
			"IndividualCustomerMaritalStatusCode": "",
			"IndividualCustomerEMail": "",
			"IndividualCustomerPhone": "",
			"IndividualCustomerFax": "",
			"IndividualCustomerMobile": "",
			"IndividualCustomerAddressHouseID": "",
			"IndividualCustomerAddressStreetName": "",
			"IndividualCustomerAddressCity": "",
			"IndividualCustomerAddressCountry": "",
			"IndividualCustomerAddressState": "",
			"IndividualCustomerAddressPostalCode": "",
			"IndividualCustomerAddressCountyName": "",
			"IndividualCustomerNationalityCountryCode": "",
			"IndividualCustomerBirthDate": "",
			"IndividualCustomerFormOfAddressCode": "",
			"IndividualCustomerAcademicTitleCode": "",
			"IndividualCustomerOccupationCode": "",
			"IndividualCustomerContactAllowedCode": "",
			"IndividualCustomerNonVerbalCommunicationLanguageCode": "",
			"IndividualCustomerInitialsName": "",
			"AccountPreferredCommunicationMediumTypeCode": "",
			"IndividualCustomerNamePrefixCode": "",
			"AccountLifeCycleStatusCode": "2",
			"MainContactLifeCycleStatusCode": "2",
			"LifeCycleStatusCode": "6",
			"ExternalID": "00163E38C29E1ED7A1B00497365581FF",
			"CreatedBy": "",
			"LastChangedBy": "Robert Mark",
			"EntityLastChangedOn": "2019-11-14T20:24:17+09:00",
			"ETag": "2019-11-14T20:24:17+09:00",
			"LeadItemCollection": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/LeadCollection('00163E0C6CDB1ED7A1B004E3242D49E9')/LeadItem"
		}
	],
	"time": "2022-08-10T11:41:39+09:00"
}
```
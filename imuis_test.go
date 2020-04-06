package imuis_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-imuis"
)

func TestCheckconnection(t *testing.T) {
	resp, err := client.CheckConnection(&imuis.CheckConnection{})
	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestAdminInfo(t *testing.T) {
	resp, err := client.GetAdmInfo(&imuis.GetAdmInfo{PartnerKey: &partnerKey, Omgevingscode: &omgevingsCode})
	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetStamtabelRecordsDebiteuren(t *testing.T) {
	// selectie := imuis.CData(`<NewDataSet>
	// <Table1>
	// 	<TABLE>DEB</TABLE>
	// 	<SELECTFIELDS>NR;ZKSL;NAAM;NAAM2;ADRES;PLAATS</SELECTFIELDS>
	// 	<WHEREFIELDS>NR</WHEREFIELDS>
	// 	<WHEREOPERATORS>=</WHEREOPERATORS>
	// 	<WHEREVALUES>10000</WHEREVALUES>
	// 	<MUTDATE_VA >01-01-2016</MUTDATE_VA > (Optioneel)
	// 	<ORDERBY>NR</ORDERBY>
	// 	<MAXRESULT>0</MAXRESULT>
	// 	<PAGESIZE>10000</PAGESIZE>
	// 	<SELECTPAGE>1</SELECTPAGE>
	// </Table1>
	// </NewDataSet>`)
	selectie := imuis.Selectie{
		Table:          "DEB",
		SelectFields:   "NR;ZKSL;NAAM;NAAM2;ADRES;PLAATS",
		WhereFields:    "NR",
		WhereOperators: ">",
		WhereValues:    "10000",
		MutDateVA:      "01-01-2016",
		OrderBy:        "NR",
		MaxResult:      0,
		PageSize:       10000,
		SelectPage:     1,
	}
	debiteuren := &[]imuis.Debiteur{}
	_, err := client.GetStamtabelRecords(&imuis.GetStamtabelRecords{
		PartnerKey:    &partnerKey,
		Omgevingscode: &omgevingsCode,
		SessionId:     &sessionID,
		Selectie:      &selectie,
	}, debiteuren)
	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(debiteuren, "", "  ")
	log.Println(string(b))
}

func TestGetStamtabelRecordsBoekingen(t *testing.T) {
	// selectie := imuis.CData(`<NewDataSet>
	// <Table1>
	// 	<TABLE>DEB</TABLE>
	// 	<SELECTFIELDS>NR;ZKSL;NAAM;NAAM2;ADRES;PLAATS</SELECTFIELDS>
	// 	<WHEREFIELDS>NR</WHEREFIELDS>
	// 	<WHEREOPERATORS>=</WHEREOPERATORS>
	// 	<WHEREVALUES>10000</WHEREVALUES>
	// 	<MUTDATE_VA >01-01-2016</MUTDATE_VA > (Optioneel)
	// 	<ORDERBY>NR</ORDERBY>
	// 	<MAXRESULT>0</MAXRESULT>
	// 	<PAGESIZE>10000</PAGESIZE>
	// 	<SELECTPAGE>1</SELECTPAGE>
	// </Table1>
	// </NewDataSet>`)
	selectie := imuis.Selectie{
		Table:          "BOE",
		SelectFields:   "AANT;AANT2;AANT3;BEDRBOEK;BEDRBOEKVAL;BEDRBTW;BEDRBTWVAL;BOEKSTUK;BTW;DAGB;DAT;FACT;JR;KDR;KOERS;KPL;OMSCHR;OPM;PN;REK;RG;STORNO;TEGREK;VAL;DOSSIER",
		WhereFields:    "AANT",
		WhereOperators: "=",
		WhereValues:    "0",
		MutDateVA:      "2020-04-01",
		OrderBy:        "BOEKSTUK",
		MaxResult:      0,
		PageSize:       10000,
		SelectPage:     1,
	}
	resp, err := client.GetStamtabelRecords(&imuis.GetStamtabelRecords{
		PartnerKey:    &partnerKey,
		Omgevingscode: &omgevingsCode,
		SessionId:     &sessionID,
		Selectie:      &selectie,
	}, &[]imuis.Journaalpost{})
	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetStamtabelRecordsBoekingenByDebiteur(t *testing.T) {
	// selectie := imuis.CData(`<NewDataSet>
	// <Table1>
	// 	<TABLE>DEB</TABLE>
	// 	<SELECTFIELDS>NR;ZKSL;NAAM;NAAM2;ADRES;PLAATS</SELECTFIELDS>
	// 	<WHEREFIELDS>NR</WHEREFIELDS>
	// 	<WHEREOPERATORS>=</WHEREOPERATORS>
	// 	<WHEREVALUES>10000</WHEREVALUES>
	// 	<MUTDATE_VA >01-01-2016</MUTDATE_VA > (Optioneel)
	// 	<ORDERBY>NR</ORDERBY>
	// 	<MAXRESULT>0</MAXRESULT>
	// 	<PAGESIZE>10000</PAGESIZE>
	// 	<SELECTPAGE>1</SELECTPAGE>
	// </Table1>
	// </NewDataSet>`)
	selectie := imuis.Selectie{
		Table:          "BOE",
		SelectFields:   "AANT;AANT2;AANT3;BEDRBOEK;BEDRBOEKVAL;BEDRBTW;BEDRBTWVAL;BOEKSTUK;BTW;DAGB;DAT;FACT;JR;KDR;KOERS;KPL;OMSCHR;OPM;PN;REK;RG;STORNO;TEGREK;VAL;DOSSIER",
		WhereFields:    "REK",
		WhereOperators: "=",
		WhereValues:    "12765",
		MutDateVA:      "2020-01-01",
		OrderBy:        "BOEKSTUK",
		MaxResult:      0,
		PageSize:       10000,
		SelectPage:     1,
	}
	resp, err := client.GetStamtabelRecords(&imuis.GetStamtabelRecords{
		PartnerKey:    &partnerKey,
		Omgevingscode: &omgevingsCode,
		SessionId:     &sessionID,
		Selectie:      &selectie,
	}, &[]imuis.Journaalpost{})
	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestCreateJournaalpost(t *testing.T) {
	// journaalpost := `NewDataSet>
	// <BOE>
	// <BEDRINCL>100</BEDRINCL>
	// <BOEKSTUK>BKST1000</BOEKSTUK>
	// <DAGB>10</DAGB>
	// <DAT>27-10-2018</DAT>
	// <JR>2018</JR>
	// <OMSCHR>Omschrijving</OMSCHR>
	// <PN>10</PN>
	// <OPM>Opmerking</OPM>
	// <BTW>12</BTW>
	// <BEDRBTW>-21</BEDRBTW>
	// <TEGREK>8000</TEGREK>
	// <REK>10000</REK>
	// </BOE>
	// </NewDataSet>`
	journaalpost := imuis.Journaalpost{
		BEDRBOEK: 100.0,
		BOEKSTUK: "TEST123",
		DAGB:     "10",
		DAT:      "27-10-2018",
		JR:       "2018",
		OMSCHR:   "omschrijving",
		PN:       "10",
		OPM:      "OPMERKING",
		BTW:      "12",
		BEDRBTW:  -21,
		TEGREK:   "8000",
		REK:      "10000",
	}
	resp, err := client.CreateJournaalpost(&imuis.CreateJournaalpost{
		PartnerKey:    &partnerKey,
		Omgevingscode: &omgevingsCode,
		SessionId:     &sessionID,
		Journaalpost:  &journaalpost,
	})
	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestCreateDebiteur(t *testing.T) {
	resp, err := client.CreateStamTabelRecord(&imuis.CreateStamTabelRecord{
		PartnerKey:    &partnerKey,
		Omgevingscode: &omgevingsCode,
		SessionId:     &sessionID,
		Stamtabel: &imuis.CreateStamTabelRecordStamtabel{
			NewDataSet: imuis.CreateStamTabelRecordNewDataSet{
				METADATA: imuis.CreateStamTabelRecordMetadata{
					TABLE: "DEB",
				},
				DATA: imuis.Debiteur{},
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

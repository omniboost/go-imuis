package imuis

import (
	"encoding/xml"
	"errors"
	"net/http"

	"github.com/omniboost/go-imuis/wsdl2go/soap"
)

type CData string

func (n CData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		S string `xml:",innerxml"`
	}{
		S: "<![CDATA[" + string(n) + "]]>",
	}, start)
}

type Selectie struct {
	Table          string `xml:"TABLE"`
	SelectFields   string `xml:"SELECTFIELDS"`
	WhereFields    string `xml:"WHEREFIELDS,omitempty"`
	WhereOperators string `xml:"WHEREOPERATORS,omitempty"`
	WhereValues    string `xml:"WHEREVALUES,omitempty"`
	MutDateVA      string `xml:"MUTDATE_VA"`
	OrderBy        string `xml:"ORDERBY"`
	MaxResult      int    `xml:"MAXRESULT"`
	PageSize       int    `xml:"PAGESIZE"`
	SelectPage     int    `xml:"SELECTPAGE"`
}

func (s Selectie) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Table1 Selectie
	a := Table1(s)
	b, err := xml.Marshal(a)
	if err != nil {
		return err
	}

	return e.EncodeElement(struct {
		S string `xml:",innerxml"`
	}{
		S: "<![CDATA[<NewDataSet>" + string(b) + "</NewDataSet>]]>",
	}, start)
}

func (jp Journaalpost) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// type Alias Journaalpost
	// type NewDataSet struct {
	// 	Boe Alias
	// }

	// wrapper := NewDataSet{Boe: Alias(jp)}
	// return e.EncodeElement(wrapper, start)
	type Boe Journaalpost
	a := Boe(jp)
	b, err := xml.Marshal(a)
	if err != nil {
		return err
	}

	return e.EncodeElement(struct {
		S string `xml:",innerxml"`
	}{
		S: "<![CDATA[<NewDataSet>" + string(b) + "</NewDataSet>]]>",
	}, start)
}

type DebiteurenTabel []Debiteur

type Debiteur struct {
	// Aanhef, datatype: CHAR, lengte: 20
	AANHEF string
	// Aanmaning, datatype: BOOLEAN
	AANM string `xml:",omitempty"`
	// Aanmaningen afdrukken of e-mailen, datatype: CHAR, lengte: 1
	// [A]fdrukken, [E]mailen
	AANMAFDRUK string `xml:",omitempty"`
	// Vaste aanmaning gebruiken, datatype: CHAR, lengte: 20
	AANMVAST string
	// Adres, datatype: CHAR, lengte: 50
	ADRES string
	// Uitgebreide adressering, datatype: MEMO, lengte: 1024
	ADRESSERING string
	// Betalingsplichtige, datatype: NUMERIC, lengte: 8
	BETALER string
	// Betalingsconditie, datatype: NUMERIC, lengte: 3, verplicht
	BETCOND string
	// Blokkeren, datatype: BOOLEAN
	BLOK string `xml:",omitempty"`
	// Blokkeren voor module declaraties/urenverantwoording vanaf, datatype: DATE
	BLOKDECLVA string
	// Blokkeren voor verkooporders vanaf, datatype: DATE
	BLOKVRKVA string
	// BIC van bankrekening, datatype: CHAR, lengte: 11
	BNKBNKREK string
	// BIC van 3e bankrekening, datatype: CHAR, lengte: 11
	BNKBNKREK2 string
	// BIC van 2e bankrekening, datatype: CHAR, lengte: 11
	BNKGIRO string
	// BIC van G-rekening, datatype: CHAR, lengte: 11
	BNKGREK string
	// Bankrekening IBAN, datatype: CHAR, lengte: 34
	BNKIBAN string
	// 3e Bankrekening IBAN, datatype: CHAR, lengte: 34
	BNKIBAN2 string
	// Bankrekening, datatype: CHAR, lengte: 15
	BNKREK string
	// 3e Bankrekening, datatype: CHAR, lengte: 15
	BNKREK2 string
	// Bankrekening numeriek, datatype: NUMERIC, lengte: 15, niet toegankelijk
	BNKREKNUM string
	// 3e Bankrekening numeriek, datatype: NUMERIC, lengte: 15, niet toegankelijk
	BNKREKNUM2 string
	// Banksoort voor incasso's, datatype: CHAR, lengte: 1
	BNKSRTINC string
	// BTW-nummer, datatype: CHAR, lengte: 14
	BTWNR string
	// BTW-plichtig, datatype: CHAR, lengte: 1, verplicht
	// {B (Plichtig),L (laag), I (Binnen EU), U (Buiten EU), N (niet plichtig)}
	BTWPL string
	// BTW-nummer verificatiedatum, datatype: DATE
	DATBTWNR string
	// Klant af, datatype: DATE
	DATKLANTAF string
	// Klant sinds, datatype: DATE
	DATKLANTSINDS string
	// Einddatum kredietlimiet, datatype: DATE
	DATKRLIMTM string
	// Begindatum kredietlimiet, datatype: DATE
	DATKRLIMVAN string
	// Kvk datum uittreksel, datatype: DATE
	DATKVKUITTR string
	// Datum laatste aanmaning, datatype: DATE, niet toegankelijk
	DATLSTAANM string
	// Datum laatste betaling, datatype: DATE, niet toegankelijk
	DATLSTBET string
	// Datum laatste factuur, datatype: DATE, niet toegankelijk
	DATLSTFACT string
	// Oprichtingsdatum/geboortedatum, datatype: DATE
	DATOPRICHTING string
	// Gebruik declaratiebudget verplicht, datatype: CHAR, lengte: 1
	DECBUDVERPL string `xml:",omitempty"`
	// EAN nummer, datatype: NUMERIC, lengte: 13
	EANNR string
	// E-mailadres, datatype: CHAR, lengte: 64
	EMAIL string
	// Nieuwsbrief e-mailen, datatype: CHAR, lengte: 1
	EMAILMAILINGJN string `xml:",omitempty"`
	// Aanbrenger extern (contactpersoon)relatie, datatype: NUMERIC, lengte: 8
	EXTAANBRENGREL string
	// Aanbrenger extern (contactpersoon)zoeksleutel, datatype: CHAR, lengte: 20
	EXTAANBRENGZKSL string
	// Facturen afdrukken of e-mailen, datatype: CHAR, lengte: 1
	// [A]fdrukken, [E]mailen
	FACTAFDRUK string `xml:",omitempty"`
	// Factoring, datatype: BOOLEAN
	FACTORING string `xml:",omitempty"`
	// Artikel/debiteur voorkeursafspraken van de betalingsplichtige overnemen naar onderliggende debiteuren, datatype: CHAR, lengte: 1
	FACVRKNAARDEB string `xml:",omitempty"`
	// Fax, datatype: CHAR, lengte: 15
	FAX string
	// Staffel gebruiken, datatype: CHAR, lengte: 1
	GEBRSTAFFEL string `xml:",omitempty"`
	// Kostprijs gebruiken als verkoopprijs bij orders, datatype: CHAR, lengte: 1
	GEBRVERKKOSTPR string `xml:",omitempty"`
	// 2e Bankrekening, datatype: CHAR, lengte: 15
	GIRO string
	// 2e Bankrekening IBAN, datatype: CHAR, lengte: 34
	GIROIBAN string
	// Tenaamstelling 2e bankrekening, datatype: CHAR, lengte: 40
	GIRONAAM string
	// 2e Bankrekening numeriek, datatype: NUMERIC, lengte: 15, niet toegankelijk
	GIRONUM string
	// G-rekening, datatype: CHAR, lengte: 15
	GREK string
	// G-rekening IBAN, datatype: CHAR, lengte: 34
	GREKIBAN string
	// G-rekening numeriek, datatype: NUMERIC, lengte: 15, niet toegankelijk
	GREKNUM string
	// Groepsdebiteur, datatype: NUMERIC, lengte: 8
	GRPDEB string
	// Heeft saldo, datatype: BOOLEAN, niet toegankelijk
	HEEFTSALDO string
	// Huisnummer, datatype: NUMERIC, lengte: 5
	HNR string
	// Huisnummertoevoeging, datatype: CHAR, lengte: 6
	HNRTV string
	// Homepage, datatype: CHAR, lengte: 64
	HOMEPAGE string
	// Specificatie incasso's afdrukken of e-mailen, datatype: CHAR, lengte: 1
	INCSELAFDRUK string `xml:",omitempty"`
	// Inkoopcombinatie, datatype: NUMERIC, lengte: 8
	INKCOMB string
	// Declaratie t/m jaar, datatype: NUMERIC, lengte: 4, niet toegankelijk
	JRDECLTM string
	// Declaratie vanaf jaar, datatype: NUMERIC, lengte: 4, niet toegankelijk
	JRDECLVAN string
	// KOSTENDRAGER, datatype: NUMERIC, lengte: 8
	KDR string
	// Kenmerk voor openstaande posten, datatype: CHAR, lengte: 25
	KENMOPP string
	// Kixcode, datatype: CHAR, lengte: 20, niet toegankelijk
	KIXCD string
	// Klant af tekst, datatype: CHAR, lengte: 20
	KLANTAFTEKST string
	// Klant sinds tekst, datatype: CHAR, lengte: 20
	KLANTSINDSTEKST string
	// KOSTENPLAATS, datatype: NUMERIC, lengte: 8
	KPL string
	// Kredietlimiet, datatype: MONEY
	KRLIM float64
	// Kvk-nummer, datatype: CHAR, lengte: 15
	KVKNR string
	// Kvk-plaats, datatype: CHAR, lengte: 24
	KVKPLAATS string
	// Laagste bedrag, datatype: BOOLEAN
	LAAGSTEBEDR string `xml:",omitempty"`
	// Land, datatype: CHAR, lengte: 3
	LAND string
	// Voorkeursleveringsconditie, datatype: CHAR, lengte: 20
	LEVCOND string
	// Verantwoordelijk medewerker, datatype: CHAR, lengte: 20
	MEDEW string
	// Aanbrenger intern (medewerker), datatype: CHAR, lengte: 20
	MEDEWAANBRENG string
	// Verantwoordelijk medewerker declaraties/urenverantwoording, datatype: CHAR, lengte: 20
	MEDEWDEC string
	// Fiscale medewerker, datatype: CHAR, lengte: 20
	MEDEWFISCAAL string
	// Loon medewerker, datatype: CHAR, lengte: 20
	MEDEWLOON string
	// Verantwoordelijke vennoot, datatype: CHAR, lengte: 20
	MEDEWVENNOOT string
	// Telefoon mobiel, datatype: CHAR, lengte: 15
	MOBIEL string
	// Naam, datatype: CHAR, lengte: 50
	NAAM string
	// Ter attentie van, datatype: CHAR, lengte: 50
	NAAM2 string
	// Nummer, datatype: NUMERIC, lengte: 8, verplicht
	NR string
	// Ons crediteurnummer bij debiteur, datatype: CHAR, lengte: 20
	NRBIJDEB string
	// Offerte afdrukken of e-mailen, datatype: CHAR, lengte: 1
	// [A]fdrukken, [E]mailen
	OFFAFDRUK string `xml:",omitempty"`
	// Voorkeursopdrachtwijze, datatype: CHAR, lengte: 20
	OPDRWZ string
	// Opmerking, datatype: MEMO, lengte: 1024
	OPM string
	// Orderbevestigingen afdrukken of e-mailen, datatype: CHAR, lengte: 1
	// [A]fdrukken, [E]mailen
	ORDBEVAFDRUK string `xml:",omitempty"`
	// Order compleet leveren, datatype: CHAR, lengte: 1
	ORDERCOMPLEET string `xml:",omitempty"`
	// Voorkeursordersoort, datatype: CHAR, lengte: 20
	ORDSRT string
	// Verzendbon ook e-mailen, datatype: CHAR, lengte: 1
	PAKBEMAIL string `xml:",omitempty"`
	// Perc. uit loonbestanddeel voor G-rekening, datatype: NUMERIC, lengte: 6
	PERCGREK string
	// Woonplaats, datatype: CHAR, lengte: 24
	PLAATS string
	// Bedrag plus/min, datatype: MONEY, niet toegankelijk
	PLM float64
	// Declaratie t/m periode, datatype: NUMERIC, lengte: 2, niet toegankelijk
	PNDECLTM string
	// Declaratie vanaf periode, datatype: NUMERIC, lengte: 2, niet toegankelijk
	PNDECLVAN string
	// Postcode, datatype: CHAR, lengte: 8
	POSTCD string
	// Gegenereerd door, datatype: CHAR, lengte: 12, niet toegankelijk
	PRG string
	// Was prospectnummer, datatype: NUMERIC, lengte: 8
	PROSP string
	// Artikelprijslijst, datatype: CHAR, lengte: 20
	PRSLST string
	// Rayon, datatype: NUMERIC, lengte: 8
	RAYON string
	// Rit, datatype: NUMERIC, lengte: 8
	RIT string
	// Rechtsvorm, datatype: CHAR, lengte: 3
	RVORM string
	// Saldo, datatype: MONEY, niet toegankelijk
	SALDO float64
	// Sluitrekening, datatype: NUMERIC, lengte: 8, verplicht
	SLUITREK string
	// Statutaire naam, datatype: CHAR, lengte: 40
	STATNAAM string
	// Statutaire plaats, datatype: CHAR, lengte: 24
	STATPLAATS string
	// Straat, datatype: CHAR, lengte: 37
	STRAAT string
	// Taal, datatype: CHAR, lengte: 3
	TAAL string
	// Voorkeurstegenrekening, datatype: NUMERIC, lengte: 8
	TEGREK string
	// Telefoon, datatype: CHAR, lengte: 15
	TEL string
	// Telefoon privé, datatype: CHAR, lengte: 15
	TELPRIVE string
	// Te ontvangen Incasso, datatype: MONEY, niet toegankelijk
	TEONTVINC float64
	// Termijn t.b.v. contracten, datatype: CHAR, lengte: 1
	TERM string `xml:",omitempty"`
	// Factuurtermijn, datatype: CHAR, lengte: 1
	TERMIJNFACT string `xml:",omitempty"`
	// Valuta, datatype: CHAR, lengte: 3
	VAL string
	// Vaste datum t.b.v. contracten, datatype: CHAR, lengte: 1
	VASTECTRDATUM string `xml:",omitempty"`
	// Verkoper, datatype: CHAR, lengte: 20
	VERKOPER string
	// Verrekenbaar, datatype: BOOLEAN
	VERR string `xml:",omitempty"`
	// Verzamelfactuur, datatype: CHAR, lengte: 1
	VERZFACT string `xml:",omitempty"`
	// Voorkeursverzendwijze, datatype: CHAR, lengte: 20
	VERZWZ string
	// Voorschotten, datatype: CHAR, lengte: 1
	VOORS string `xml:",omitempty"`
	// DEBVRIJVELD1, datatype: CHAR, lengte: 40
	VRIJVELD1 string
	// DEBVRIJVELD2, datatype: CHAR, lengte: 40
	VRIJVELD2 string
	// DEBVRIJVELD3, datatype: CHAR, lengte: 20
	VRIJVELD3 string
	// DEBVRIJVELD4, datatype: CHAR, lengte: 20
	VRIJVELD4 string
	// DEBVRIJVELD5, datatype: CHAR, lengte: 20
	VRIJVELD5 string
	// Zoeksleutel, datatype: CHAR, lengte: 20, primary key, verplicht
	ZKSL string
}

type Journaalpost struct {
	// Aantal, datatype: NUMERIC, lengte: 11
	AANT string `xml:",omitempty"`
	// Aantal2, datatype: NUMERIC, lengte: 11
	AANT2 string `xml:",omitempty"`
	// Aantal3, datatype: NUMERIC, lengte: 11
	AANT3 string `xml:",omitempty"`
	// Boekbedrag, datatype: MONEY
	BEDRBOEK string
	// Valuta boekbedrag, datatype: MONEY
	BEDRBOEKVAL string
	// BTW-bedrag, datatype: MONEY
	BEDRBTW string
	// Valuta BTW-bedrag, datatype: MONEY
	BEDRBTWVAL string
	// Betwiste factuur, datatype: BOOLEAN
	BETWIST string `xml:",omitempty"`
	// Optioneel: Alleen gebruiken bij debiteur/crediteur journaalpost om de factuur
	// (openstaande post) na aanmaken direct op betwist te zetten.
	// BOEKSTUK, datatype: CHAR, lengte: 20
	BOEKSTUK string
	// BTW-code, datatype: NUMERIC, lengte: 2
	BTW string `xml:",omitempty"`
	// Dagboek, NUMERIC, lengte: 4, primary key, verplicht
	DAGB string
	// Datum, datatype: DATE, verplicht
	DAT string
	// BOEKSTUK, datatype: CHAR, lengte: 20
	DATVERV string `xml:",omitempty"`
	// Optioneel: Alleen gebruiken bij debiteur/crediteur journaalpost om de factuur
	// (openstaande post) na aanmaken direct een vervaldatum te geven.
	// FACT, datatype: CHAR, lengte: 20
	FACT string `xml:",omitempty"`
	// Jaar, datatype: NUMERIC, lengte: 4, primary key, verplicht
	JR string
	// KOSTENDRAGER, datatype: NUMERIC, lengte: 8
	KDR string `xml:",omitempty"`
	// Koers, datatype: NUMERIC, lengte: 13
	KOERS string `xml:",omitempty"`
	// KOSTENPLAATS, datatype: NUMERIC, lengte: 8
	KPL string `xml:",omitempty"`
	// Omschrijving, datatype: CHAR, lengte: 40
	OMSCHR string
	// Opmerking, datatype: CHAR, lengte: 250
	OPM string
	// Periode, datatype: NUMERIC, lengte: 2, primary key
	PN string
	// Rekening, datatype: NUMERIC, lengte: 8
	REK string
	// Regel, datatype: NUMERIC, lengte: 8, primary key, verplicht
	RG string `xml:",omitempty"`
	// Storno, datatype: BOOLEAN
	STORNO string `xml:",omitempty"`
	// Tegenrekening, datatype: NUMERIC, lengte: 8, verplicht
	TEGREK string
	// Valuta, datatype: CHAR, lengte: 3
	VAL string `xml:",omitempty"`
	// Dossier, datatype: CHAR, lengte: 20
	DOSSIER string `xml:",omitempty"`

	BOE_WARNING string
}

type GetStamtabelRecordsResponseRecords struct {
	NewDataSet NewDataSet `xml:"NewDataSet"`
}

func (rr *GetStamtabelRecordsResponseRecords) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	b := xml.CharData{}
	err := d.DecodeElement(&b, &start)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(b, &rr.NewDataSet)
	if err != nil {
		return err
	}

	return nil
}

type NewDataSet struct {
	Data interface{} `xml:"DATA"`
}

type Client struct {
	Ws1_xmlSoap
}

func NewClient(h *http.Client) Client {
	if h == nil {
		h = http.DefaultClient
	}

	cli := &soap.Client{
		URL:         "https://api.kingfinance.nl/v1/ws1_xml.asmx",
		Namespace:   Namespace,
		Config:      h,
		ContentType: "text/xml; charset=utf-8; action=\"%s\"",
		Envelope:    "http://www.w3.org/2003/05/soap-envelope",
	}

	return Client{NewWs1_xmlSoap(cli)}
}

func (c *Client) SetDebug(debug bool) {
}

func (c Client) CreateJournaalpost(CreateJournaalpost *CreateJournaalpost) (*CreateJournaalpostResponse, error) {
	resp, err := c.Ws1_xmlSoap.CreateJournaalpost(CreateJournaalpost)
	if err != nil {
		return resp, err
	}

	if resp.Foutmelding != nil && *resp.Foutmelding != "" {
		return resp, errors.New(*resp.Foutmelding)
	}

	return resp, err
}

func (c Client) GetStamtabelRecords(getStamtabelRecords *GetStamtabelRecords, data interface{}) (*GetStamtabelRecordsResponse, error) {
	resp, err := c.Ws1_xmlSoap.GetStamtabelRecords(getStamtabelRecords, data)
	if err != nil {
		return resp, err
	}

	if resp.Foutmelding != nil && *resp.Foutmelding != "" {
		return resp, errors.New(*resp.Foutmelding)
	}

	return resp, err
}

func (c Client) CreateStamTabelRecord(createStamTabelRecord *CreateStamTabelRecord) (*CreateStamTabelRecordResponse, error) {
	resp, err := c.Ws1_xmlSoap.CreateStamTabelRecord(createStamTabelRecord)
	if err != nil {
		return resp, err
	}

	if resp.Foutmelding != nil && *resp.Foutmelding != "" {
		return resp, errors.New(*resp.Foutmelding)
	}

	return resp, err
}

type CreateJournaalpostResponseJournaalpost Journaalpost

func (jp *CreateJournaalpostResponseJournaalpost) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	b := xml.CharData{}
	err := d.DecodeElement(&b, &start)
	if err != nil {
		return err
	}

	t := struct {
		Boe Journaalpost
	}{}

	err = xml.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*jp = CreateJournaalpostResponseJournaalpost(t.Boe)
	return nil
}

type CreateStamTabelRecordStamtabel struct {
	NewDataSet CreateStamTabelRecordNewDataSet
}

func (st CreateStamTabelRecordStamtabel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	b, err := xml.Marshal(st.NewDataSet)
	if err != nil {
		return err
	}

	return e.EncodeElement(struct {
		S string `xml:",innerxml"`
	}{
		S: "<![CDATA[" + string(b) + "]]>",
	}, start)
}

type UpdateStamTabelRecordMutatie struct {
	NewDataSet UpdateStamTabelRecordNewDataSet
}

func (st UpdateStamTabelRecordMutatie) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	b, err := xml.Marshal(st.NewDataSet)
	if err != nil {
		return err
	}

	return e.EncodeElement(struct {
		S string `xml:",innerxml"`
	}{
		S: "<![CDATA[" + string(b) + "]]>",
	}, start)
}

type CreateStamTabelRecordNewDataSet struct {
	METADATA CreateStamTabelRecordMetadata
	DATA     interface{}
}

type CreateStamTabelRecordMetadata struct {
	TABLE string
}

type UpdateStamTabelRecordNewDataSet struct {
	SELECTION Selection
	DATA      interface{}
}

type UpdateStamTabelRecordMetadata struct {
	TABLE string
}

type Selection struct {
	Table          string `xml:"TABLE"`
	SelectFields   string `xml:"SELECTFIELDS"`
	WhereFields    string `xml:"WHEREFIELDS,omitempty"`
	WhereOperators string `xml:"WHEREOPERATORS,omitempty"`
	WhereValues    string `xml:"WHEREVALUES,omitempty"`
}

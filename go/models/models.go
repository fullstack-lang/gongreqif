package models

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	Name     string
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type Length string

type LinkTypes []NMTOKEN

type MediaDesc string

type MultiLength string

// type Number uint32

// type Pixels uint32

type Script string

type Color string

type Text string

type Character string

type Charset string

type Charsets []Charset

type ContentType string

type ContentTypes string

type Datetime time.Time

type FPI NormalizedString

type FrameTarget string

type LanguageCode Language

type LanguageCodes string

type URI AnyURI

type URIs []AnyURI

type URIREF string

type MultiLengths string

type CDATA string

type CURIE string

type CURIEs []CURIE

type SafeCURIE string

type SafeCURIEs []SafeCURIE

type URIorSafeCURIE string

type URIorSafeCURIEs []URIorSafeCURIE

type Xhtml_address_type struct {
	Name string
}

type Xhtml_blockquote_type struct {
	Name string
}

type Xhtml_pre_type struct {
	Name string
}

type Xhtml_heading_type struct {
	Name string
}

type Xhtml_h1_type struct {
	Name string
}

type Xhtml_h2_type struct {
	Name string
}

type Xhtml_h3_type struct {
	Name string
}

type Xhtml_h4_type struct {
	Name string
}

type Xhtml_h5_type struct {
	Name string
}

type Xhtml_h6_type struct {
	Name string
}

type Xhtml_div_type struct {
	Name string
}

type Xhtml_p_type struct {
	Name string
}

type Xhtml_abbr_type struct {
	Name string
}

type Xhtml_acronym_type struct {
	Name string
}

type Xhtml_cite_type struct {
	Name string
}

type Xhtml_code_type struct {
	Name string
}

type Xhtml_dfn_type struct {
	Name string
}

type Xhtml_em_type struct {
	Name string
}

type Xhtml_kbd_type struct {
	Name string
}

type Xhtml_samp_type struct {
	Name string
}

type Xhtml_strong_type struct {
	Name string
}

type Xhtml_var_type struct {
	Name string
}

type Xhtml_q_type struct {
	Name string
}

type Xhtml_br_type struct {
	Name string
}

type Xhtml_span_type struct {
	Name string
}

type Xhtml_a_type struct {
	Name string
}

type Xhtml_dt_type struct {
	Name string
}

type Xhtml_dd_type struct {
	Name string
}

type Xhtml_dl_type struct {
	Name string
}

type Xhtml_li_type struct {
	Name string
}

type Xhtml_ol_type struct {
	Name string
}

type Xhtml_ul_type struct {
	Name string
}

type Xhtml_edit_type struct {
	Name string
}

type Xhtml_hr_type struct {
	Name string
}

type Xhtml_InlPres_type struct {
	Name string
}

type Xhtml_param_type struct {
	Name string
}

type Xhtml_object_type struct {
	Name string
}

type Xhtml_td_type struct {
	Name string
}

type Xhtml_th_type struct {
	Name string
}

type Xhtml_tr_type struct {
	Name string
}

type Xhtml_col_type struct {
	Name string
}

type Xhtml_colgroup_type struct {
	Name string
}

type Xhtml_tbody_type struct {
	Name string
}

type Xhtml_tfoot_type struct {
	Name string
}

type Xhtml_thead_type struct {
	Name string
}

type Xhtml_caption_type struct {
	Name string
}

type Xhtml_table_type struct {
	Name string
}

type LOCAL_REF IDREF

type GLOBAL_REF string

type ALTERNATIVE_ID struct {
	Name       string
	IDENTIFIER string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`
}

type ATTRIBUTE_DEFINITION_BOOLEAN struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	DEFAULT_VALUE struct {
		ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN `xml:"ATTRIBUTE-VALUE-BOOLEAN,omitempty" json:"ATTRIBUTE-VALUE-BOOLEAN,omitempty"`
	} `xml:"DEFAULT-VALUE,omitempty" json:"DEFAULT-VALUE,omitempty"`

	TYPE struct {
		DATATYPE_DEFINITION_BOOLEAN_REF *LOCAL_REF `xml:"DATATYPE-DEFINITION-BOOLEAN-REF,omitempty" json:"DATATYPE-DEFINITION-BOOLEAN-REF,omitempty"`
	} `xml:"TYPE,omitempty" json:"TYPE,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	IS_EDITABLE bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IS-EDITABLE,attr,omitempty" json:"IS-EDITABLE,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type ATTRIBUTE_DEFINITION_DATE struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	DEFAULT_VALUE struct {
		ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE `xml:"ATTRIBUTE-VALUE-DATE,omitempty" json:"ATTRIBUTE-VALUE-DATE,omitempty"`
	} `xml:"DEFAULT-VALUE,omitempty" json:"DEFAULT-VALUE,omitempty"`

	TYPE struct {
		DATATYPE_DEFINITION_DATE_REF *LOCAL_REF `xml:"DATATYPE-DEFINITION-DATE-REF,omitempty" json:"DATATYPE-DEFINITION-DATE-REF,omitempty"`
	} `xml:"TYPE,omitempty" json:"TYPE,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	IS_EDITABLE bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IS-EDITABLE,attr,omitempty" json:"IS-EDITABLE,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type ATTRIBUTE_DEFINITION_ENUMERATION struct {
	Name          string
	DEFAULT_VALUE struct {
		ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION `xml:"ATTRIBUTE-VALUE-ENUMERATION,omitempty" json:"ATTRIBUTE-VALUE-ENUMERATION,omitempty"`
	} `xml:"DEFAULT-VALUE,omitempty" json:"DEFAULT-VALUE,omitempty"`

	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	TYPE struct {
		DATATYPE_DEFINITION_ENUMERATION_REF *LOCAL_REF `xml:"DATATYPE-DEFINITION-ENUMERATION-REF,omitempty" json:"DATATYPE-DEFINITION-ENUMERATION-REF,omitempty"`
	} `xml:"TYPE,omitempty" json:"TYPE,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	IS_EDITABLE bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IS-EDITABLE,attr,omitempty" json:"IS-EDITABLE,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`

	MULTI_VALUED bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd MULTI-VALUED,attr,omitempty" json:"MULTI-VALUED,omitempty"`
}

type ATTRIBUTE_DEFINITION_INTEGER struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	DEFAULT_VALUE struct {
		ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER `xml:"ATTRIBUTE-VALUE-INTEGER,omitempty" json:"ATTRIBUTE-VALUE-INTEGER,omitempty"`
	} `xml:"DEFAULT-VALUE,omitempty" json:"DEFAULT-VALUE,omitempty"`

	TYPE struct {
		DATATYPE_DEFINITION_INTEGER_REF *LOCAL_REF `xml:"DATATYPE-DEFINITION-INTEGER-REF,omitempty" json:"DATATYPE-DEFINITION-INTEGER-REF,omitempty"`
	} `xml:"TYPE,omitempty" json:"TYPE,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	IS_EDITABLE bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IS-EDITABLE,attr,omitempty" json:"IS-EDITABLE,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type ATTRIBUTE_DEFINITION_REAL struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	DEFAULT_VALUE struct {
		ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL `xml:"ATTRIBUTE-VALUE-REAL,omitempty" json:"ATTRIBUTE-VALUE-REAL,omitempty"`
	} `xml:"DEFAULT-VALUE,omitempty" json:"DEFAULT-VALUE,omitempty"`

	TYPE struct {
		DATATYPE_DEFINITION_REAL_REF *LOCAL_REF `xml:"DATATYPE-DEFINITION-REAL-REF,omitempty" json:"DATATYPE-DEFINITION-REAL-REF,omitempty"`
	} `xml:"TYPE,omitempty" json:"TYPE,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	IS_EDITABLE bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IS-EDITABLE,attr,omitempty" json:"IS-EDITABLE,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type ATTRIBUTE_DEFINITION_STRING struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	DEFAULT_VALUE struct {
		ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING `xml:"ATTRIBUTE-VALUE-STRING,omitempty" json:"ATTRIBUTE-VALUE-STRING,omitempty"`
	} `xml:"DEFAULT-VALUE,omitempty" json:"DEFAULT-VALUE,omitempty"`

	TYPE struct {
		DATATYPE_DEFINITION_STRING_REF *LOCAL_REF `xml:"DATATYPE-DEFINITION-STRING-REF,omitempty" json:"DATATYPE-DEFINITION-STRING-REF,omitempty"`
	} `xml:"TYPE,omitempty" json:"TYPE,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	IS_EDITABLE bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IS-EDITABLE,attr,omitempty" json:"IS-EDITABLE,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type ATTRIBUTE_DEFINITION_XHTML struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	DEFAULT_VALUE struct {
		ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML `xml:"ATTRIBUTE-VALUE-XHTML,omitempty" json:"ATTRIBUTE-VALUE-XHTML,omitempty"`
	} `xml:"DEFAULT-VALUE,omitempty" json:"DEFAULT-VALUE,omitempty"`

	TYPE struct {
		DATATYPE_DEFINITION_XHTML_REF *LOCAL_REF `xml:"DATATYPE-DEFINITION-XHTML-REF,omitempty" json:"DATATYPE-DEFINITION-XHTML-REF,omitempty"`
	} `xml:"TYPE,omitempty" json:"TYPE,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	IS_EDITABLE bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IS-EDITABLE,attr,omitempty" json:"IS-EDITABLE,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type ATTRIBUTE_VALUE_BOOLEAN struct {
	Name       string
	DEFINITION struct {
		ATTRIBUTE_DEFINITION_BOOLEAN_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-BOOLEAN-REF,omitempty" json:"ATTRIBUTE-DEFINITION-BOOLEAN-REF,omitempty"`
	} `xml:"DEFINITION,omitempty" json:"DEFINITION,omitempty"`

	THE_VALUE bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd THE-VALUE,attr,omitempty" json:"THE-VALUE,omitempty"`
}

type ATTRIBUTE_VALUE_DATE struct {
	Name       string
	DEFINITION struct {
		ATTRIBUTE_DEFINITION_DATE_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-DATE-REF,omitempty" json:"ATTRIBUTE-DEFINITION-DATE-REF,omitempty"`
	} `xml:"DEFINITION,omitempty" json:"DEFINITION,omitempty"`

	THE_VALUE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd THE-VALUE,attr,omitempty" json:"THE-VALUE,omitempty"`
}

type ATTRIBUTE_VALUE_ENUMERATION struct {
	Name       string
	DEFINITION struct {
		ATTRIBUTE_DEFINITION_ENUMERATION_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-ENUMERATION-REF,omitempty" json:"ATTRIBUTE-DEFINITION-ENUMERATION-REF,omitempty"`
	} `xml:"DEFINITION,omitempty" json:"DEFINITION,omitempty"`

	VALUES struct {
		ENUM_VALUE_REF *LOCAL_REF `xml:"ENUM-VALUE-REF,omitempty" json:"ENUM-VALUE-REF,omitempty"`
	} `xml:"VALUES,omitempty" json:"VALUES,omitempty"`
}

type ATTRIBUTE_VALUE_INTEGER struct {
	Name       string
	DEFINITION struct {
		ATTRIBUTE_DEFINITION_INTEGER_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-INTEGER-REF,omitempty" json:"ATTRIBUTE-DEFINITION-INTEGER-REF,omitempty"`
	} `xml:"DEFINITION,omitempty" json:"DEFINITION,omitempty"`

	THE_VALUE int32 `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd THE-VALUE,attr,omitempty" json:"THE-VALUE,omitempty"`
}

type ATTRIBUTE_VALUE_REAL struct {
	Name       string
	DEFINITION struct {
		ATTRIBUTE_DEFINITION_REAL_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-REAL-REF,omitempty" json:"ATTRIBUTE-DEFINITION-REAL-REF,omitempty"`
	} `xml:"DEFINITION,omitempty" json:"DEFINITION,omitempty"`

	THE_VALUE float64 `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd THE-VALUE,attr,omitempty" json:"THE-VALUE,omitempty"`
}

type ATTRIBUTE_VALUE_STRING struct {
	Name       string
	DEFINITION struct {
		ATTRIBUTE_DEFINITION_STRING_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-STRING-REF,omitempty" json:"ATTRIBUTE-DEFINITION-STRING-REF,omitempty"`
	} `xml:"DEFINITION,omitempty" json:"DEFINITION,omitempty"`

	THE_VALUE string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd THE-VALUE,attr,omitempty" json:"THE-VALUE,omitempty"`
}

type ATTRIBUTE_VALUE_XHTML struct {
	Name      string
	THE_VALUE []*XHTML_CONTENT `xml:"THE-VALUE,omitempty" json:"THE-VALUE,omitempty"`

	THE_ORIGINAL_VALUE []*XHTML_CONTENT `xml:"THE-ORIGINAL-VALUE,omitempty" json:"THE-ORIGINAL-VALUE,omitempty"`

	DEFINITION struct {
		ATTRIBUTE_DEFINITION_XHTML_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-XHTML-REF,omitempty" json:"ATTRIBUTE-DEFINITION-XHTML-REF,omitempty"`
	} `xml:"DEFINITION,omitempty" json:"DEFINITION,omitempty"`

	IS_SIMPLIFIED bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IS-SIMPLIFIED,attr,omitempty" json:"IS-SIMPLIFIED,omitempty"`
}

type DATATYPE_DEFINITION_BOOLEAN struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type DATATYPE_DEFINITION_DATE struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type DATATYPE_DEFINITION_ENUMERATION struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	SPECIFIED_VALUES struct {
		ENUM_VALUE []*ENUM_VALUE `xml:"ENUM-VALUE,omitempty" json:"ENUM-VALUE,omitempty"`
	} `xml:"SPECIFIED-VALUES,omitempty" json:"SPECIFIED-VALUES,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type DATATYPE_DEFINITION_INTEGER struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`

	MAX int32 `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd MAX,attr,omitempty" json:"MAX,omitempty"`

	MIN int32 `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd MIN,attr,omitempty" json:"MIN,omitempty"`
}

type DATATYPE_DEFINITION_REAL struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	ACCURACY int32 `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd ACCURACY,attr,omitempty" json:"ACCURACY,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`

	MAX float64 `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd MAX,attr,omitempty" json:"MAX,omitempty"`

	MIN float64 `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd MIN,attr,omitempty" json:"MIN,omitempty"`
}

type DATATYPE_DEFINITION_STRING struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`

	MAX_LENGTH int32 `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd MAX-LENGTH,attr,omitempty" json:"MAX-LENGTH,omitempty"`
}

type DATATYPE_DEFINITION_XHTML struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type EMBEDDED_VALUE struct {
	Name string
	KEY  int32 `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd KEY,attr,omitempty" json:"KEY,omitempty"`

	OTHER_CONTENT string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd OTHER-CONTENT,attr,omitempty" json:"OTHER-CONTENT,omitempty"`
}

type ENUM_VALUE struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	PROPERTIES struct {
		EMBEDDED_VALUE []*EMBEDDED_VALUE `xml:"EMBEDDED-VALUE,omitempty" json:"EMBEDDED-VALUE,omitempty"`
	} `xml:"PROPERTIES,omitempty" json:"PROPERTIES,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type RELATION_GROUP struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	SOURCE_SPECIFICATION struct {
		SPECIFICATION_REF []*GLOBAL_REF `xml:"SPECIFICATION-REF,omitempty" json:"SPECIFICATION-REF,omitempty"`
	} `xml:"SOURCE-SPECIFICATION,omitempty" json:"SOURCE-SPECIFICATION,omitempty"`

	SPEC_RELATIONS struct {
		SPEC_RELATION_REF *LOCAL_REF `xml:"SPEC-RELATION-REF,omitempty" json:"SPEC-RELATION-REF,omitempty"`
	} `xml:"SPEC-RELATIONS,omitempty" json:"SPEC-RELATIONS,omitempty"`

	TARGET_SPECIFICATION struct {
		SPECIFICATION_REF []*GLOBAL_REF `xml:"SPECIFICATION-REF,omitempty" json:"SPECIFICATION-REF,omitempty"`
	} `xml:"TARGET-SPECIFICATION,omitempty" json:"TARGET-SPECIFICATION,omitempty"`

	TYPE struct {
		RELATION_GROUP_TYPE_REF *LOCAL_REF `xml:"RELATION-GROUP-TYPE-REF,omitempty" json:"RELATION-GROUP-TYPE-REF,omitempty"`
	} `xml:"TYPE,omitempty" json:"TYPE,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type RELATION_GROUP_TYPE struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	SPEC_ATTRIBUTES struct {
		ATTRIBUTE_DEFINITION_BOOLEAN []*ATTRIBUTE_DEFINITION_BOOLEAN `xml:"ATTRIBUTE-DEFINITION-BOOLEAN,omitempty" json:"ATTRIBUTE-DEFINITION-BOOLEAN,omitempty"`

		ATTRIBUTE_DEFINITION_DATE []*ATTRIBUTE_DEFINITION_DATE `xml:"ATTRIBUTE-DEFINITION-DATE,omitempty" json:"ATTRIBUTE-DEFINITION-DATE,omitempty"`

		ATTRIBUTE_DEFINITION_ENUMERATION []*ATTRIBUTE_DEFINITION_ENUMERATION `xml:"ATTRIBUTE-DEFINITION-ENUMERATION,omitempty" json:"ATTRIBUTE-DEFINITION-ENUMERATION,omitempty"`

		ATTRIBUTE_DEFINITION_INTEGER []*ATTRIBUTE_DEFINITION_INTEGER `xml:"ATTRIBUTE-DEFINITION-INTEGER,omitempty" json:"ATTRIBUTE-DEFINITION-INTEGER,omitempty"`

		ATTRIBUTE_DEFINITION_REAL []*ATTRIBUTE_DEFINITION_REAL `xml:"ATTRIBUTE-DEFINITION-REAL,omitempty" json:"ATTRIBUTE-DEFINITION-REAL,omitempty"`

		ATTRIBUTE_DEFINITION_STRING []*ATTRIBUTE_DEFINITION_STRING `xml:"ATTRIBUTE-DEFINITION-STRING,omitempty" json:"ATTRIBUTE-DEFINITION-STRING,omitempty"`

		ATTRIBUTE_DEFINITION_XHTML []*ATTRIBUTE_DEFINITION_XHTML `xml:"ATTRIBUTE-DEFINITION-XHTML,omitempty" json:"ATTRIBUTE-DEFINITION-XHTML,omitempty"`
	} `xml:"SPEC-ATTRIBUTES,omitempty" json:"SPEC-ATTRIBUTES,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type REQ_IF struct {
	Name       string
	THE_HEADER struct {
		REQ_IF_HEADER *REQ_IF_HEADER `xml:"REQ-IF-HEADER,omitempty" json:"REQ-IF-HEADER,omitempty"`
	} `xml:"THE-HEADER,omitempty" json:"THE-HEADER,omitempty"`

	CORE_CONTENT struct {
		REQ_IF_CONTENT *REQ_IF_CONTENT `xml:"REQ-IF-CONTENT,omitempty" json:"REQ-IF-CONTENT,omitempty"`
	} `xml:"CORE-CONTENT,omitempty" json:"CORE-CONTENT,omitempty"`

	TOOL_EXTENSIONS struct {
		REQ_IF_TOOL_EXTENSION []*REQ_IF_TOOL_EXTENSION `xml:"REQ-IF-TOOL-EXTENSION,omitempty" json:"REQ-IF-TOOL-EXTENSION,omitempty"`
	} `xml:"TOOL-EXTENSIONS,omitempty" json:"TOOL-EXTENSIONS,omitempty"`

	Lang string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd lang,attr,omitempty" json:"lang,omitempty"`
}

type REQ_IF_CONTENT struct {
	Name      string
	DATATYPES struct {
		DATATYPE_DEFINITION_BOOLEAN []*DATATYPE_DEFINITION_BOOLEAN `xml:"DATATYPE-DEFINITION-BOOLEAN,omitempty" json:"DATATYPE-DEFINITION-BOOLEAN,omitempty"`

		DATATYPE_DEFINITION_DATE []*DATATYPE_DEFINITION_DATE `xml:"DATATYPE-DEFINITION-DATE,omitempty" json:"DATATYPE-DEFINITION-DATE,omitempty"`

		DATATYPE_DEFINITION_ENUMERATION []*DATATYPE_DEFINITION_ENUMERATION `xml:"DATATYPE-DEFINITION-ENUMERATION,omitempty" json:"DATATYPE-DEFINITION-ENUMERATION,omitempty"`

		DATATYPE_DEFINITION_INTEGER []*DATATYPE_DEFINITION_INTEGER `xml:"DATATYPE-DEFINITION-INTEGER,omitempty" json:"DATATYPE-DEFINITION-INTEGER,omitempty"`

		DATATYPE_DEFINITION_REAL []*DATATYPE_DEFINITION_REAL `xml:"DATATYPE-DEFINITION-REAL,omitempty" json:"DATATYPE-DEFINITION-REAL,omitempty"`

		DATATYPE_DEFINITION_STRING []*DATATYPE_DEFINITION_STRING `xml:"DATATYPE-DEFINITION-STRING,omitempty" json:"DATATYPE-DEFINITION-STRING,omitempty"`

		DATATYPE_DEFINITION_XHTML []*DATATYPE_DEFINITION_XHTML `xml:"DATATYPE-DEFINITION-XHTML,omitempty" json:"DATATYPE-DEFINITION-XHTML,omitempty"`
	} `xml:"DATATYPES,omitempty" json:"DATATYPES,omitempty"`

	SPEC_TYPES struct {
		RELATION_GROUP_TYPE []*RELATION_GROUP_TYPE `xml:"RELATION-GROUP-TYPE,omitempty" json:"RELATION-GROUP-TYPE,omitempty"`

		SPEC_OBJECT_TYPE []*SPEC_OBJECT_TYPE `xml:"SPEC-OBJECT-TYPE,omitempty" json:"SPEC-OBJECT-TYPE,omitempty"`

		SPEC_RELATION_TYPE []*SPEC_RELATION_TYPE `xml:"SPEC-RELATION-TYPE,omitempty" json:"SPEC-RELATION-TYPE,omitempty"`

		SPECIFICATION_TYPE []*SPECIFICATION_TYPE `xml:"SPECIFICATION-TYPE,omitempty" json:"SPECIFICATION-TYPE,omitempty"`
	} `xml:"SPEC-TYPES,omitempty" json:"SPEC-TYPES,omitempty"`

	SPEC_OBJECTS struct {
		SPEC_OBJECT []*SPEC_OBJECT `xml:"SPEC-OBJECT,omitempty" json:"SPEC-OBJECT,omitempty"`
	} `xml:"SPEC-OBJECTS,omitempty" json:"SPEC-OBJECTS,omitempty"`

	SPEC_RELATIONS struct {
		SPEC_RELATION []*SPEC_RELATION `xml:"SPEC-RELATION,omitempty" json:"SPEC-RELATION,omitempty"`
	} `xml:"SPEC-RELATIONS,omitempty" json:"SPEC-RELATIONS,omitempty"`

	SPECIFICATIONS struct {
		SPECIFICATION []*SPECIFICATION `xml:"SPECIFICATION,omitempty" json:"SPECIFICATION,omitempty"`
	} `xml:"SPECIFICATIONS,omitempty" json:"SPECIFICATIONS,omitempty"`

	SPEC_RELATION_GROUPS struct {
		RELATION_GROUP []*RELATION_GROUP `xml:"RELATION-GROUP,omitempty" json:"RELATION-GROUP,omitempty"`
	} `xml:"SPEC-RELATION-GROUPS,omitempty" json:"SPEC-RELATION-GROUPS,omitempty"`
}

type REQ_IF_HEADER struct {
	Name    string
	COMMENT string `xml:"COMMENT,omitempty" json:"COMMENT,omitempty"`

	CREATION_TIME time.Time `xml:"CREATION-TIME,omitempty" json:"CREATION-TIME,omitempty"`

	REPOSITORY_ID string `xml:"REPOSITORY-ID,omitempty" json:"REPOSITORY-ID,omitempty"`

	REQ_IF_TOOL_ID string `xml:"REQ-IF-TOOL-ID,omitempty" json:"REQ-IF-TOOL-ID,omitempty"`

	REQ_IF_VERSION string `xml:"REQ-IF-VERSION,omitempty" json:"REQ-IF-VERSION,omitempty"`

	SOURCE_TOOL_ID string `xml:"SOURCE-TOOL-ID,omitempty" json:"SOURCE-TOOL-ID,omitempty"`

	TITLE string `xml:"TITLE,omitempty" json:"TITLE,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`
}

type SPEC_HIERARCHY struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	CHILDREN struct {
		SPEC_HIERARCHY []*SPEC_HIERARCHY `xml:"SPEC-HIERARCHY,omitempty" json:"SPEC-HIERARCHY,omitempty"`
	} `xml:"CHILDREN,omitempty" json:"CHILDREN,omitempty"`

	EDITABLE_ATTS struct {
		ATTRIBUTE_DEFINITION_BOOLEAN_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-BOOLEAN-REF,omitempty" json:"ATTRIBUTE-DEFINITION-BOOLEAN-REF,omitempty"`

		ATTRIBUTE_DEFINITION_DATE_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-DATE-REF,omitempty" json:"ATTRIBUTE-DEFINITION-DATE-REF,omitempty"`

		ATTRIBUTE_DEFINITION_ENUMERATION_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-ENUMERATION-REF,omitempty" json:"ATTRIBUTE-DEFINITION-ENUMERATION-REF,omitempty"`

		ATTRIBUTE_DEFINITION_INTEGER_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-INTEGER-REF,omitempty" json:"ATTRIBUTE-DEFINITION-INTEGER-REF,omitempty"`

		ATTRIBUTE_DEFINITION_REAL_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-REAL-REF,omitempty" json:"ATTRIBUTE-DEFINITION-REAL-REF,omitempty"`

		ATTRIBUTE_DEFINITION_STRING_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-STRING-REF,omitempty" json:"ATTRIBUTE-DEFINITION-STRING-REF,omitempty"`

		ATTRIBUTE_DEFINITION_XHTML_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-XHTML-REF,omitempty" json:"ATTRIBUTE-DEFINITION-XHTML-REF,omitempty"`
	} `xml:"EDITABLE-ATTS,omitempty" json:"EDITABLE-ATTS,omitempty"`

	OBJECT struct {
		SPEC_OBJECT_REF *LOCAL_REF `xml:"SPEC-OBJECT-REF,omitempty" json:"SPEC-OBJECT-REF,omitempty"`
	} `xml:"OBJECT,omitempty" json:"OBJECT,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	IS_EDITABLE bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IS-EDITABLE,attr,omitempty" json:"IS-EDITABLE,omitempty"`

	IS_TABLE_INTERNAL bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IS-TABLE-INTERNAL,attr,omitempty" json:"IS-TABLE-INTERNAL,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type SPEC_OBJECT struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	VALUES struct {
		ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN `xml:"ATTRIBUTE-VALUE-BOOLEAN,omitempty" json:"ATTRIBUTE-VALUE-BOOLEAN,omitempty"`

		ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE `xml:"ATTRIBUTE-VALUE-DATE,omitempty" json:"ATTRIBUTE-VALUE-DATE,omitempty"`

		ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION `xml:"ATTRIBUTE-VALUE-ENUMERATION,omitempty" json:"ATTRIBUTE-VALUE-ENUMERATION,omitempty"`

		ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER `xml:"ATTRIBUTE-VALUE-INTEGER,omitempty" json:"ATTRIBUTE-VALUE-INTEGER,omitempty"`

		ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL `xml:"ATTRIBUTE-VALUE-REAL,omitempty" json:"ATTRIBUTE-VALUE-REAL,omitempty"`

		ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING `xml:"ATTRIBUTE-VALUE-STRING,omitempty" json:"ATTRIBUTE-VALUE-STRING,omitempty"`

		ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML `xml:"ATTRIBUTE-VALUE-XHTML,omitempty" json:"ATTRIBUTE-VALUE-XHTML,omitempty"`
	} `xml:"VALUES,omitempty" json:"VALUES,omitempty"`

	TYPE struct {
		SPEC_OBJECT_TYPE_REF *LOCAL_REF `xml:"SPEC-OBJECT-TYPE-REF,omitempty" json:"SPEC-OBJECT-TYPE-REF,omitempty"`
	} `xml:"TYPE,omitempty" json:"TYPE,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type SPEC_OBJECT_TYPE struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	SPEC_ATTRIBUTES struct {
		ATTRIBUTE_DEFINITION_BOOLEAN []*ATTRIBUTE_DEFINITION_BOOLEAN `xml:"ATTRIBUTE-DEFINITION-BOOLEAN,omitempty" json:"ATTRIBUTE-DEFINITION-BOOLEAN,omitempty"`

		ATTRIBUTE_DEFINITION_DATE []*ATTRIBUTE_DEFINITION_DATE `xml:"ATTRIBUTE-DEFINITION-DATE,omitempty" json:"ATTRIBUTE-DEFINITION-DATE,omitempty"`

		ATTRIBUTE_DEFINITION_ENUMERATION []*ATTRIBUTE_DEFINITION_ENUMERATION `xml:"ATTRIBUTE-DEFINITION-ENUMERATION,omitempty" json:"ATTRIBUTE-DEFINITION-ENUMERATION,omitempty"`

		ATTRIBUTE_DEFINITION_INTEGER []*ATTRIBUTE_DEFINITION_INTEGER `xml:"ATTRIBUTE-DEFINITION-INTEGER,omitempty" json:"ATTRIBUTE-DEFINITION-INTEGER,omitempty"`

		ATTRIBUTE_DEFINITION_REAL []*ATTRIBUTE_DEFINITION_REAL `xml:"ATTRIBUTE-DEFINITION-REAL,omitempty" json:"ATTRIBUTE-DEFINITION-REAL,omitempty"`

		ATTRIBUTE_DEFINITION_STRING []*ATTRIBUTE_DEFINITION_STRING `xml:"ATTRIBUTE-DEFINITION-STRING,omitempty" json:"ATTRIBUTE-DEFINITION-STRING,omitempty"`

		ATTRIBUTE_DEFINITION_XHTML []*ATTRIBUTE_DEFINITION_XHTML `xml:"ATTRIBUTE-DEFINITION-XHTML,omitempty" json:"ATTRIBUTE-DEFINITION-XHTML,omitempty"`
	} `xml:"SPEC-ATTRIBUTES,omitempty" json:"SPEC-ATTRIBUTES,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type SPEC_RELATION struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	VALUES struct {
		ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN `xml:"ATTRIBUTE-VALUE-BOOLEAN,omitempty" json:"ATTRIBUTE-VALUE-BOOLEAN,omitempty"`

		ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE `xml:"ATTRIBUTE-VALUE-DATE,omitempty" json:"ATTRIBUTE-VALUE-DATE,omitempty"`

		ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION `xml:"ATTRIBUTE-VALUE-ENUMERATION,omitempty" json:"ATTRIBUTE-VALUE-ENUMERATION,omitempty"`

		ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER `xml:"ATTRIBUTE-VALUE-INTEGER,omitempty" json:"ATTRIBUTE-VALUE-INTEGER,omitempty"`

		ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL `xml:"ATTRIBUTE-VALUE-REAL,omitempty" json:"ATTRIBUTE-VALUE-REAL,omitempty"`

		ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING `xml:"ATTRIBUTE-VALUE-STRING,omitempty" json:"ATTRIBUTE-VALUE-STRING,omitempty"`

		ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML `xml:"ATTRIBUTE-VALUE-XHTML,omitempty" json:"ATTRIBUTE-VALUE-XHTML,omitempty"`
	} `xml:"VALUES,omitempty" json:"VALUES,omitempty"`

	SOURCE struct {
		SPEC_OBJECT_REF *GLOBAL_REF `xml:"SPEC-OBJECT-REF,omitempty" json:"SPEC-OBJECT-REF,omitempty"`
	} `xml:"SOURCE,omitempty" json:"SOURCE,omitempty"`

	TARGET struct {
		SPEC_OBJECT_REF *GLOBAL_REF `xml:"SPEC-OBJECT-REF,omitempty" json:"SPEC-OBJECT-REF,omitempty"`
	} `xml:"TARGET,omitempty" json:"TARGET,omitempty"`

	TYPE struct {
		SPEC_RELATION_TYPE_REF *LOCAL_REF `xml:"SPEC-RELATION-TYPE-REF,omitempty" json:"SPEC-RELATION-TYPE-REF,omitempty"`
	} `xml:"TYPE,omitempty" json:"TYPE,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type SPEC_RELATION_TYPE struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	SPEC_ATTRIBUTES struct {
		ATTRIBUTE_DEFINITION_BOOLEAN []*ATTRIBUTE_DEFINITION_BOOLEAN `xml:"ATTRIBUTE-DEFINITION-BOOLEAN,omitempty" json:"ATTRIBUTE-DEFINITION-BOOLEAN,omitempty"`

		ATTRIBUTE_DEFINITION_DATE []*ATTRIBUTE_DEFINITION_DATE `xml:"ATTRIBUTE-DEFINITION-DATE,omitempty" json:"ATTRIBUTE-DEFINITION-DATE,omitempty"`

		ATTRIBUTE_DEFINITION_ENUMERATION []*ATTRIBUTE_DEFINITION_ENUMERATION `xml:"ATTRIBUTE-DEFINITION-ENUMERATION,omitempty" json:"ATTRIBUTE-DEFINITION-ENUMERATION,omitempty"`

		ATTRIBUTE_DEFINITION_INTEGER []*ATTRIBUTE_DEFINITION_INTEGER `xml:"ATTRIBUTE-DEFINITION-INTEGER,omitempty" json:"ATTRIBUTE-DEFINITION-INTEGER,omitempty"`

		ATTRIBUTE_DEFINITION_REAL []*ATTRIBUTE_DEFINITION_REAL `xml:"ATTRIBUTE-DEFINITION-REAL,omitempty" json:"ATTRIBUTE-DEFINITION-REAL,omitempty"`

		ATTRIBUTE_DEFINITION_STRING []*ATTRIBUTE_DEFINITION_STRING `xml:"ATTRIBUTE-DEFINITION-STRING,omitempty" json:"ATTRIBUTE-DEFINITION-STRING,omitempty"`

		ATTRIBUTE_DEFINITION_XHTML []*ATTRIBUTE_DEFINITION_XHTML `xml:"ATTRIBUTE-DEFINITION-XHTML,omitempty" json:"ATTRIBUTE-DEFINITION-XHTML,omitempty"`
	} `xml:"SPEC-ATTRIBUTES,omitempty" json:"SPEC-ATTRIBUTES,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type SPECIFICATION struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	VALUES struct {
		ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN `xml:"ATTRIBUTE-VALUE-BOOLEAN,omitempty" json:"ATTRIBUTE-VALUE-BOOLEAN,omitempty"`

		ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE `xml:"ATTRIBUTE-VALUE-DATE,omitempty" json:"ATTRIBUTE-VALUE-DATE,omitempty"`

		ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION `xml:"ATTRIBUTE-VALUE-ENUMERATION,omitempty" json:"ATTRIBUTE-VALUE-ENUMERATION,omitempty"`

		ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER `xml:"ATTRIBUTE-VALUE-INTEGER,omitempty" json:"ATTRIBUTE-VALUE-INTEGER,omitempty"`

		ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL `xml:"ATTRIBUTE-VALUE-REAL,omitempty" json:"ATTRIBUTE-VALUE-REAL,omitempty"`

		ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING `xml:"ATTRIBUTE-VALUE-STRING,omitempty" json:"ATTRIBUTE-VALUE-STRING,omitempty"`

		ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML `xml:"ATTRIBUTE-VALUE-XHTML,omitempty" json:"ATTRIBUTE-VALUE-XHTML,omitempty"`
	} `xml:"VALUES,omitempty" json:"VALUES,omitempty"`

	CHILDREN struct {
		SPEC_HIERARCHY []*SPEC_HIERARCHY `xml:"SPEC-HIERARCHY,omitempty" json:"SPEC-HIERARCHY,omitempty"`
	} `xml:"CHILDREN,omitempty" json:"CHILDREN,omitempty"`

	TYPE struct {
		SPECIFICATION_TYPE_REF *LOCAL_REF `xml:"SPECIFICATION-TYPE-REF,omitempty" json:"SPECIFICATION-TYPE-REF,omitempty"`
	} `xml:"TYPE,omitempty" json:"TYPE,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type SPECIFICATION_TYPE struct {
	Name           string
	ALTERNATIVE_ID struct {
		ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	} `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	SPEC_ATTRIBUTES struct {
		ATTRIBUTE_DEFINITION_BOOLEAN []*ATTRIBUTE_DEFINITION_BOOLEAN `xml:"ATTRIBUTE-DEFINITION-BOOLEAN,omitempty" json:"ATTRIBUTE-DEFINITION-BOOLEAN,omitempty"`

		ATTRIBUTE_DEFINITION_DATE []*ATTRIBUTE_DEFINITION_DATE `xml:"ATTRIBUTE-DEFINITION-DATE,omitempty" json:"ATTRIBUTE-DEFINITION-DATE,omitempty"`

		ATTRIBUTE_DEFINITION_ENUMERATION []*ATTRIBUTE_DEFINITION_ENUMERATION `xml:"ATTRIBUTE-DEFINITION-ENUMERATION,omitempty" json:"ATTRIBUTE-DEFINITION-ENUMERATION,omitempty"`

		ATTRIBUTE_DEFINITION_INTEGER []*ATTRIBUTE_DEFINITION_INTEGER `xml:"ATTRIBUTE-DEFINITION-INTEGER,omitempty" json:"ATTRIBUTE-DEFINITION-INTEGER,omitempty"`

		ATTRIBUTE_DEFINITION_REAL []*ATTRIBUTE_DEFINITION_REAL `xml:"ATTRIBUTE-DEFINITION-REAL,omitempty" json:"ATTRIBUTE-DEFINITION-REAL,omitempty"`

		ATTRIBUTE_DEFINITION_STRING []*ATTRIBUTE_DEFINITION_STRING `xml:"ATTRIBUTE-DEFINITION-STRING,omitempty" json:"ATTRIBUTE-DEFINITION-STRING,omitempty"`

		ATTRIBUTE_DEFINITION_XHTML []*ATTRIBUTE_DEFINITION_XHTML `xml:"ATTRIBUTE-DEFINITION-XHTML,omitempty" json:"ATTRIBUTE-DEFINITION-XHTML,omitempty"`
	} `xml:"SPEC-ATTRIBUTES,omitempty" json:"SPEC-ATTRIBUTES,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

type REQ_IF_TOOL_EXTENSION struct {
	Name  string
	Items []string `xml:",any" json:"items,omitempty"`
}

type XHTML_CONTENT struct {
	Name string
}

package enum

type BwSearchEngineEnum string

const (
	// Global
	BwsSearchEngineGoogle      BwSearchEngineEnum = "GOOGLE"
	BwsSearchEngineBing        BwSearchEngineEnum = "BING"
	BwsSearchEngineYahoo       BwSearchEngineEnum = "YAHOO"
	BwsSearchEngineDuckDuckGo  BwSearchEngineEnum = "DUCKDUCKGO"
	BwsSearchEngineEcosia      BwSearchEngineEnum = "ECOSIA"
	BwsSearchEngineBraveSearch BwSearchEngineEnum = "BRAVESEARCH"
	BwsSearchEngineQwant       BwSearchEngineEnum = "QWANT"
	BwsSearchEngineStartpage   BwSearchEngineEnum = "STARTPAGE"
	BwsSearchEngineAsk         BwSearchEngineEnum = "ASK"

	// Regional
	BwsSearchEngineBaidu  BwSearchEngineEnum = "BAIDU"
	BwsSearchEngineYandex BwSearchEngineEnum = "YANDEX"
	BwsSearchEngineNaver  BwSearchEngineEnum = "NAVER"
	BwsSearchEngineSeznam BwSearchEngineEnum = "SEZNAM"
	BwsSearchEngineDaum   BwSearchEngineEnum = "DAUM"
	BwsSearchEngineSogou  BwSearchEngineEnum = "SOGOU"

	// Academic
	BwsSearchEngineGoogleScholar   BwSearchEngineEnum = "GOOGLE_SCHOLAR"
	BwsSearchEngineSemanticScholar BwSearchEngineEnum = "SEMANTIC_SCHOLAR"
	BwsSearchEnginePubMed          BwSearchEngineEnum = "PUBMED"
	BwsSearchEngineWolframAlpha    BwSearchEngineEnum = "WOLFRAM_ALPHA"
	BwsSearchEngineArchiveOrg      BwSearchEngineEnum = "ARCHIVE_ORG"
)

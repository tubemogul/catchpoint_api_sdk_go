package alertsAPI

// This File contains all the struct definitions needed for the xml parser for
// catchpoint Alert Push API according to the xsd they provide as of 2015-07-06

import (
	"encoding/xml"
)

type NodeThreshold struct {
	TypeId     uint8
	NodeCount  int64
	NodePct    float32
	OrOperator string
}

type InsightDataSource struct {
	id           int `xml:"id,attr"`
	Name         string
	DataSourceId uint8
}

type TriggerInsight struct {
	Indicator  InsightDataSource
	Tracepoint InsightDataSource
}

type TriggerZone struct {
	id   int `xml:"id,attr"`
	Name string
}

type Trigger struct {
	TypeId             uint8
	HistoricalInterval uint16
	OperatorId         uint8
	Warning            float32
	Critical           float32
	RegexMatch         string
	Insight            TriggerInsight
	Zone               TriggerZone
	CompareId          uint8
	CompareModifier    float32
}

type Setting struct {
	AlertGroupId                  int
	AlertGroupItemId              int
	AlertTypeId                   uint8
	AlertSubTypeId                uint8
	AlertGroupItemFilterTypeId    uint8
	AlertGroupItemFilterTypeValue string
	NodeThreshold                 NodeThreshold
	TimeThreshold                 uint16 // Unit = minutes
	Trigger                       Trigger
}

type TimestampData struct {
	Local, Utc string
}

type AlertTimestamp struct {
	Create, ReportInterval TimestampData
	ProcessingUtc          string
}

type AlertTestDetail struct {
	Name          string
	MonitorTypeId uint8
	TypeId        uint8
	Path          string
	Url           string
	ClientName    string
	DivisionName  string
	ProductName   string
}

type AlertNodeTriggered struct {
	Counter, Max, Min      int
	Mean, Median, Trailing float32
}

type AlertNodePageFailure struct {
	ErrorCode, HttpStatusCode int
}

type AlertNodeSuspect struct {
	Url                          string
	ObjectResponse, PageResponse int // Unit = milliseconds
}

type AlertNodeHostFailure struct {
	HostsFailed                 uint32
	WorstHost, WorstHostDetails string
}

type AlertNode struct {
	XmlName                                      xml.Name `xml:"Node"`
	Id                                           int      `xml:"id,attr"`
	Name, IpAddress, RemoteIpAddress, IsCritical string
	TransactionStepIndex                         uint8
	ProbableCauseId                              int8
	Counter                                      int
	Mean                                         float32
	Triggered                                    AlertNodeTriggered   // Applies for 'Response', 'ByteLength' and 'Insight (Indicator)' alerts
	PageFailure                                  AlertNodePageFailure // Applies only for 'PageFaliure' alert
	Suspect                                      AlertNodeSuspect     // Applies only for 'ResponseTimeTotalPageLoadWithSuspect alert sub-type
	HostFailure                                  AlertNodeHostFailure // Applies only for 'HostFailure' alert
}

type ConditionRuns struct {
	Detected, Expected int
}

type AlertCondition struct {
	NodeCount                  int64
	NodePct                    float32
	AverageAcrossNodes         float64
	AverageAcrossNodesTrailing float64
	TransactionStepIndex       uint8
	Runs                       ConditionRuns
	Nodes                      []AlertNode `xml:"Nodes>Node"`
}

type PingGroupPacketSent struct {
	Total, Failed  int8
	RoundTripTimes []int `xml:"RoundTripTimes>RoundTripTime"` // Unit = milliseconds (version since Aug-2011)
}

type PingGroup struct {
	V                         uint `xml:"v,attr"`
	Address, Host, Asn        string
	BufferSize, FailureStatus int16
	FromDebugPrimaryHost      string
	Duration                  int // Unit = milliseconds
	PacketsSent               PingGroupPacketSent
}

type TraceRouteGroup struct {
	V                        uint `xml:"v,attr"`
	Address, Host, Timestamp string
	ErrorCode                int8
	FromDebugPrimaryHost     string
	Duration                 int         // Unit = milliseconds
	Hops                     []PingGroup `xml:"Hops>Hop"`
}

type DnsServer struct {
	Address, HostName, Name string
	Port                    uint8
}

type DnsQuery struct {
	XmlName      xml.Name `xml:"Query"`
	V            uint     `xml:"v,attr"`
	Server       DnsServer
	ReturnCode   uint
	ResponseTime int
	ErrorMessage string
	Ping         PingGroup
	TraceRoute   []TraceRouteGroup `xml:"TraceRoute>TraceRouteGroup"`
}

type DnsResponse struct {
	XmlName                xml.Name `xml:"Response"`
	V                      uint     `xml:"v,attr"`
	Name, Info, Address    string
	Ttl                    uint
	Class                  uint8
	Type, InnerResolveTime uint16
	InnerResolveErrorCode  int8
}
type DnsGroup struct {
	XmlName   xml.Name      `xml:"Group"`
	Queries   []DnsQuery    `xml:"Queries>Query"`
	Responses []DnsResponse `xml:"Responses>Response"`
}

type DiagnosticDnsTraversalLastLevel struct {
	// For 'QueryType' enumerations see:  http://www.iana.org/assignments/dns-parameters
	QueryType uint16
	V         uint `xml:"v,attr"`
	level     int8
	Groups    []DnsGroup `xml:"Groups>Group"`
}

type AlertDiagnostic struct {
	Ping                  PingGroup
	TraceRoute            TraceRouteGroup
	DnsTraversalLastLevel DiagnosticDnsTraversalLastLevel
}

type Alert struct {
	XmlName             xml.Name `xml:"Alert"`
	Version             uint     `xml:"version,attr"`
	V                   uint     `xml:"v,attr"`
	TestId              uint64   `xml:"testId,attr"`
	NotificationLevelId uint8    `xml:"notificationLevelId,attr"`
	DivisionId          int      `xml:"divisionId,attr"`
	ProductId           int      `xml:"productId,attr"`
	Setting             Setting
	Timestamp           AlertTimestamp
	TestDetail          AlertTestDetail
	Condition           AlertCondition
	Diagnostic          AlertDiagnostic
}

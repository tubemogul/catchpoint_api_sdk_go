package catchpoint

// This File contains all the helpers definitions by the different parsers
// that can decode the messages received from the Catchpoint Push API.

// This variable contains the translation of the Alert.Setting.AlertTypeId
// variable into its human-friendly definition
var AlertTypeIdLabel = map[uint8]string{
	2:  "ByteLength",
	3:  "ContentMatch",
	4:  "HostFailure",
	5:  "HttpHeaderMatch",
	7:  "ResponseTime",
	8:  "Traffic (RU only)",
	9:  "PageFailure",
	10: "Insight",
	11: "ScriptFailure (IE,Chrome only)",
	12: "Ping",
	13: "Requests",
	14: "Zone",
	15: "Availability",
	16: "Address",
}

// This variable contains the translation of the Alert.Setting.AlertSubTypeId
// variable into its human-friendly definition
var AlertSubTypeIdLabel = map[uint8]string{
	1:   "ByteLengthPage",
	2:   "ByteLengthTotalPage",
	3:   "ByteLengthPageResponseContent",
	10:  "ContentMatchRegularExpression",
	11:  "ContentMatchLastCheckedFile",
	12:  "ContentMatchUploadedFile",
	13:  "ContentMatchRegularExpressionWithNumericCheck",
	30:  "HttpHeaderMatchConnection",
	31:  "HttpHeaderMatchContentType",
	32:  "HttpHeaderMatchCustom",
	33:  "HttpHeaderMatchEncoding",
	34:  "HttpHeaderMatchRedirect",
	35:  "HttpHeaderMatchVersion",
	50:  "ResponseTimeDns",
	51:  "ResponseTimeConnect",
	52:  "ResponseTimeSend (for FTP, refers to 'Upload-Time')",
	53:  "ResponseTimeWait",
	54:  "ResponseTimeLoad - from First byte to Last byte. (for FTP, refers to 'Donwload-Time')",
	55:  "ResponseTimeTimeToFirstByte",
	56:  "ResponseTimePageLoad",
	57:  "ResponseTimePageElementsLoad: Web based synthetic agents -> Total-Time - \"Page's First-Byte Time Full (includes all redirects)\" + \"Page's last redirect Wait-Time\"",
	58:  "ResponseTimeTotalPageLoad: SN only -> \"Page's Load-Time Full (DNS + Conn + Send + Wait + Load for all redirects)\"",
	59:  "ResponseTimeTotalPageElementsLoad: RU -> Total-Time + Exit-To-Entry Time ; SN -> Total-Time",
	61:  "ResponseTimeDomLoad",
	63:  "ResponseTimeTotalPageLoadWithObjectSuspect (Web, Transaction, HTML only -> Total-Time, plus obtaining worst performing Object)",
	64:  "ResponseTimeServerResponse",
	65:  "ResponseTimeFtpDelete",
	66:  "ResponseTimeDocumentComplete",
	90:  "InsightTracepoint",
	91:  "InsightIndicator",
	100: "PingRoundTripTimeAvg",
	101: "PingPacketLoss",
	110: "RequestsObjects",
	111: "RequestsHosts",
	112: "RequestsConnections",
	113: "RequestsRedirects",
	114: "RequestsContentTypeOther",
	115: "RequestsContentTypeImage",
	116: "RequestsContentTypeScript",
	117: "RequestsContentTypeHtml",
	118: "RequestsContentTypeCss",
	119: "RequestsContentTypeXml",
	120: "RequestsContentTypeFlash",
	121: "RequestsContentTypeMedia",
	130: "ContentZoneStartToEndTime",
	131: "ContentZoneBottleneckTime",
	132: "ContentZoneBottleneckTimePct",
	133: "ContentZoneRequests",
	134: "ContentZoneRequestFailures",
	140: "AvailabilityPage",
	141: "AvailabilityContent",
	150: "AddressTestUrl",
	151: "AddressChild",
	152: "AddressPage",
}

// This variable contains the translation of the Alert.Setting.AlertGroupItemFilterTypeId
// variable into its human-friendly definition
var AlertGroupItemFilterTypeIdLabel = map[uint8]string{
	0: "None",
	1: "Index",
	2: "Name",
	3: "Address",
}

// This variable contains the translation of the
// Alert.Setting.NodeThreshold.TypeId
// variable into its human-friendly definition
var NodeThresholdTypeIdLabel = map[uint8]string{
	0: "NodeAgainstTotalAverage (Default)",
	1: "AverageAcrossNodes",
	2: "NodeAgainstNodeAverage",
}

// This variable contains the translation of the Alert.Setting.Trigger.TypeId
// into its human-friendly definition
var TriggerTypeIdLabel = map[uint8]string{
	1: "Default",
	2: "PercentageDeviationFromAverage",
}

// This variable contains the translation of the
// Alert.Setting.Trigger.OperatorId
// variable into its human-friendly definition
var TriggerOperatorIdLabel = map[uint8]string{
	0: "NotEquals",
	1: "Equals",
	2: "GreaterThan",
	3: "GreaterThanOrEquals",
	4: "LessThan",
	5: "LessThanOrEquals",
	6: "NotBetween",
	7: "Between",
}

// This variable contains the translation of the
// Alert.Setting.Trigger.Insight.Indicator.DataSourceId
// and Alert.Setting.Trigger.Insight.Tracepoint.DataSourceId
// variables into its human-friendly definition
var IndicatorDataSourceIdLabel = map[uint8]string{
	0: "HttpHeader",
	1: "HttpContent",
}

// This variable contains the translation of the Alert.Setting.Trigger.CompareId
// variable into its human-friendly definition
var IndicatorCompareIdLabel = map[uint8]string{
	0:  "Counter",
	1:  "Average",
	2:  "Median",
	3:  "StandardDeviation",
	4:  "Min",
	5:  "Max",
	6:  "GeometricMean",
	7:  "Percentile95",
	8:  "Percentile85",
	14: "InterQuartileRange",
	15: "Percentile75",
	16: "Percentile99",
	17: "GeometricStandardDeviation",
	18: "PercentileCustom",
	19: "Percentile25",
	21: "Total",
}

// This variable contains the translation of the Alert.TestDetail.MonitorTypeId
// variable into its human-friendly definition
var MonitorTypeIdLabel = map[uint8]string{
	0:  "IE",
	2:  "OBJECT",
	3:  "EMULATED",
	8:  "PING",
	9:  "TRACERT",
	10: "DNSTRAVERSAL",
	11: "PING-TCP",
	12: "DNSEXP",
	13: "DNSDIRECT",
	14: "TRACERT-UDP",
	15: "PORTTCP",
	16: "FTP",
	17: "DATAPUSHAPI",
	18: "CHROME",
	19: "PLAYBACK",
	20: "PLAYBACK-MOBILE",
	21: "SMTP",
	22: "PORT-UDP",
	23: "PING-UDP",
	24: "STREAMING",
	25: "API",
	26: "MOBILE",
	27: "SFTP",
	28: "SSH",
	29: "TRACERT-TCP",
}

// This variable contains the translation of the Alert.TestDetail.TypeId
// variable into its human-friendly definition
var TestDetailTypeIdLabel = map[uint8]string{
	0:  "Web",
	1:  "Transaction",
	2:  "HtmlCode",
	3:  "Ftp",
	4:  "Tcp",
	5:  "Dns",
	6:  "Ping",
	7:  "Smtp",
	8:  "Udp",
	9:  "API",
	10: "Streaming",
	11: "SSH",
	12: "TraceRoute",
}

// This variable contains the translation of the
// Alert.Diagnostic.DnsTraversalLastLevel.Groups.Group.Responses.Response.Class
// variable into its human-friendly definition
// For now there is only 1 value but keeping this as it can evolve
var ResponseClassLabel = map[uint8]string{
	1: "IN (Internet)",
}

// This variable contains the translation of the Alert.TestDetail.TypeId
// variable into its human-friendly definition
var AlertNotificationLevelIdLabel = map[uint8]string{
	0: "Warning",
	1: "Critical",
	2: "Information",
	3: "Improved",
}

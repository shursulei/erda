// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: trace.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*GetTraceQueryConditionsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetTraceQueryConditionsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TraceQueryConditions)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TraceQueryCondition)(nil)
var _ urlenc.URLValuesUnmarshaler = (*OtherTraceQueryCondition)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetTraceDebugRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateTraceDebugRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*StopTraceDebugRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetTraceDebugStatusByRequestIDRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetTraceDebugHistoriesRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetSpansRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetSpanDashboardsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetTracesRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetSpansResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*SpanAnalysis)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetSpanDashboardsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetTracesResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetTraceDebugHistoriesResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetTraceDebugResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateTraceDebugResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*StopTraceDebugResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetTraceDebugStatusByRequestIDResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TraceDebug)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TraceDebugStatus)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TraceDebugHistory)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Span)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Trace)(nil)
var _ urlenc.URLValuesUnmarshaler = (*SpanEvent)(nil)
var _ urlenc.URLValuesUnmarshaler = (*SpanEventRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*SpanEventResponse)(nil)

// GetTraceQueryConditionsRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetTraceQueryConditionsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "tenantID":
				m.TenantID = vals[0]
			}
		}
	}
	return nil
}

// GetTraceQueryConditionsResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetTraceQueryConditionsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &TraceQueryConditions{}
				}
			}
		}
	}
	return nil
}

// TraceQueryConditions implement urlenc.URLValuesUnmarshaler.
func (m *TraceQueryConditions) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// TraceQueryCondition implement urlenc.URLValuesUnmarshaler.
func (m *TraceQueryCondition) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "value":
				m.Value = vals[0]
			case "displayName":
				m.DisplayName = vals[0]
			}
		}
	}
	return nil
}

// OtherTraceQueryCondition implement urlenc.URLValuesUnmarshaler.
func (m *OtherTraceQueryCondition) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "value":
				m.Value = vals[0]
			case "displayName":
				m.DisplayName = vals[0]
			case "type":
				m.Type = vals[0]
			case "paramKey":
				m.ParamKey = vals[0]
			}
		}
	}
	return nil
}

// GetTraceDebugRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetTraceDebugRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "requestID":
				m.RequestID = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			}
		}
	}
	return nil
}

// CreateTraceDebugRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateTraceDebugRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "requestID":
				m.RequestID = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "url":
				m.Url = vals[0]
			case "body":
				m.Body = vals[0]
			case "status":
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Status = int32(val)
			case "statusName":
				m.StatusName = vals[0]
			case "responseCode":
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.ResponseCode = int32(val)
			case "responseBody":
				m.ResponseBody = vals[0]
			case "method":
				m.Method = vals[0]
			case "createTime":
				m.CreateTime = vals[0]
			case "updateTime":
				m.UpdateTime = vals[0]
			case "projectID":
				m.ProjectID = vals[0]
			case "name":
				m.Name = vals[0]
			}
		}
	}
	return nil
}

// StopTraceDebugRequest implement urlenc.URLValuesUnmarshaler.
func (m *StopTraceDebugRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "requestID":
				m.RequestID = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			}
		}
	}
	return nil
}

// GetTraceDebugStatusByRequestIDRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetTraceDebugStatusByRequestIDRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "requestID":
				m.RequestID = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			}
		}
	}
	return nil
}

// GetTraceDebugHistoriesRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetTraceDebugHistoriesRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scopeID":
				m.ScopeID = vals[0]
			case "limit":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Limit = val
			}
		}
	}
	return nil
}

// GetSpansRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetSpansRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "traceID":
				m.TraceID = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "limit":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Limit = val
			case "orgName":
				m.OrgName = vals[0]
			case "startTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.StartTime = val
			}
		}
	}
	return nil
}

// GetSpanDashboardsRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetSpanDashboardsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "tenantID":
				m.TenantID = vals[0]
			case "type":
				m.Type = vals[0]
			case "serviceInstanceID":
				m.ServiceInstanceID = vals[0]
			case "startTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.StartTime = val
			case "endTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.EndTime = val
			}
		}
	}
	return nil
}

// GetTracesRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetTracesRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "tenantID":
				m.TenantID = vals[0]
			case "status":
				m.Status = vals[0]
			case "startTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.StartTime = val
			case "endTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.EndTime = val
			case "limit":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Limit = val
			case "traceID":
				m.TraceID = vals[0]
			case "durationMin":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.DurationMin = val
			case "durationMax":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.DurationMax = val
			case "sort":
				m.Sort = vals[0]
			case "serviceName":
				m.ServiceName = vals[0]
			case "rpcMethod":
				m.RpcMethod = vals[0]
			case "httpPath":
				m.HttpPath = vals[0]
			case "pageNo":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageNo = val
			case "pageSize":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageSize = val
			}
		}
	}
	return nil
}

// GetSpansResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetSpansResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "duration":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Duration = val
			case "serviceCount":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ServiceCount = val
			case "depth":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Depth = val
			case "spanCount":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.SpanCount = val
			}
		}
	}
	return nil
}

// SpanAnalysis implement urlenc.URLValuesUnmarshaler.
func (m *SpanAnalysis) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "dashboardID":
				m.DashboardID = vals[0]
			case "conditions":
				m.Conditions = vals
			}
		}
	}
	return nil
}

// GetSpanDashboardsResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetSpanDashboardsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "callAnalysis":
				if m.CallAnalysis == nil {
					m.CallAnalysis = &SpanAnalysis{}
				}
			case "callAnalysis.dashboardID":
				if m.CallAnalysis == nil {
					m.CallAnalysis = &SpanAnalysis{}
				}
				m.CallAnalysis.DashboardID = vals[0]
			case "callAnalysis.conditions":
				if m.CallAnalysis == nil {
					m.CallAnalysis = &SpanAnalysis{}
				}
				m.CallAnalysis.Conditions = vals
			case "serviceAnalysis":
				if m.ServiceAnalysis == nil {
					m.ServiceAnalysis = &SpanAnalysis{}
				}
			case "serviceAnalysis.dashboardID":
				if m.ServiceAnalysis == nil {
					m.ServiceAnalysis = &SpanAnalysis{}
				}
				m.ServiceAnalysis.DashboardID = vals[0]
			case "serviceAnalysis.conditions":
				if m.ServiceAnalysis == nil {
					m.ServiceAnalysis = &SpanAnalysis{}
				}
				m.ServiceAnalysis.Conditions = vals
			}
		}
	}
	return nil
}

// GetTracesResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetTracesResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "pageNo":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageNo = val
			case "pageSize":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageSize = val
			case "total":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Total = val
			}
		}
	}
	return nil
}

// GetTraceDebugHistoriesResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetTraceDebugHistoriesResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &TraceDebug{}
				}
			case "data.limit":
				if m.Data == nil {
					m.Data = &TraceDebug{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.Limit = int32(val)
			case "data.offset":
				if m.Data == nil {
					m.Data = &TraceDebug{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Offset = val
			case "data.total":
				if m.Data == nil {
					m.Data = &TraceDebug{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.Total = int32(val)
			}
		}
	}
	return nil
}

// GetTraceDebugResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetTraceDebugResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &TraceDebugHistory{}
				}
			case "data.requestID":
				if m.Data == nil {
					m.Data = &TraceDebugHistory{}
				}
				m.Data.RequestID = vals[0]
			case "data.scopeID":
				if m.Data == nil {
					m.Data = &TraceDebugHistory{}
				}
				m.Data.ScopeID = vals[0]
			case "data.url":
				if m.Data == nil {
					m.Data = &TraceDebugHistory{}
				}
				m.Data.Url = vals[0]
			case "data.body":
				if m.Data == nil {
					m.Data = &TraceDebugHistory{}
				}
				m.Data.Body = vals[0]
			case "data.status":
				if m.Data == nil {
					m.Data = &TraceDebugHistory{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.Status = int32(val)
			case "data.statusName":
				if m.Data == nil {
					m.Data = &TraceDebugHistory{}
				}
				m.Data.StatusName = vals[0]
			case "data.responseCode":
				if m.Data == nil {
					m.Data = &TraceDebugHistory{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.ResponseCode = int32(val)
			case "data.responseBody":
				if m.Data == nil {
					m.Data = &TraceDebugHistory{}
				}
				m.Data.ResponseBody = vals[0]
			case "data.method":
				if m.Data == nil {
					m.Data = &TraceDebugHistory{}
				}
				m.Data.Method = vals[0]
			case "data.createTime":
				if m.Data == nil {
					m.Data = &TraceDebugHistory{}
				}
				m.Data.CreateTime = vals[0]
			case "data.updateTime":
				if m.Data == nil {
					m.Data = &TraceDebugHistory{}
				}
				m.Data.UpdateTime = vals[0]
			case "data.name":
				if m.Data == nil {
					m.Data = &TraceDebugHistory{}
				}
				m.Data.Name = vals[0]
			}
		}
	}
	return nil
}

// CreateTraceDebugResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateTraceDebugResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &TraceDebugStatus{}
				}
			case "data.requestID":
				if m.Data == nil {
					m.Data = &TraceDebugStatus{}
				}
				m.Data.RequestID = vals[0]
			case "data.status":
				if m.Data == nil {
					m.Data = &TraceDebugStatus{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.Status = int32(val)
			case "data.statusName":
				if m.Data == nil {
					m.Data = &TraceDebugStatus{}
				}
				m.Data.StatusName = vals[0]
			case "data.scopeID":
				if m.Data == nil {
					m.Data = &TraceDebugStatus{}
				}
				m.Data.ScopeID = vals[0]
			}
		}
	}
	return nil
}

// StopTraceDebugResponse implement urlenc.URLValuesUnmarshaler.
func (m *StopTraceDebugResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetTraceDebugStatusByRequestIDResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetTraceDebugStatusByRequestIDResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &TraceDebugStatus{}
				}
			case "data.requestID":
				if m.Data == nil {
					m.Data = &TraceDebugStatus{}
				}
				m.Data.RequestID = vals[0]
			case "data.status":
				if m.Data == nil {
					m.Data = &TraceDebugStatus{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.Status = int32(val)
			case "data.statusName":
				if m.Data == nil {
					m.Data = &TraceDebugStatus{}
				}
				m.Data.StatusName = vals[0]
			case "data.scopeID":
				if m.Data == nil {
					m.Data = &TraceDebugStatus{}
				}
				m.Data.ScopeID = vals[0]
			}
		}
	}
	return nil
}

// TraceDebug implement urlenc.URLValuesUnmarshaler.
func (m *TraceDebug) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "limit":
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Limit = int32(val)
			case "offset":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Offset = val
			case "total":
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Total = int32(val)
			}
		}
	}
	return nil
}

// TraceDebugStatus implement urlenc.URLValuesUnmarshaler.
func (m *TraceDebugStatus) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "requestID":
				m.RequestID = vals[0]
			case "status":
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Status = int32(val)
			case "statusName":
				m.StatusName = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			}
		}
	}
	return nil
}

// TraceDebugHistory implement urlenc.URLValuesUnmarshaler.
func (m *TraceDebugHistory) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "requestID":
				m.RequestID = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "url":
				m.Url = vals[0]
			case "body":
				m.Body = vals[0]
			case "status":
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Status = int32(val)
			case "statusName":
				m.StatusName = vals[0]
			case "responseCode":
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.ResponseCode = int32(val)
			case "responseBody":
				m.ResponseBody = vals[0]
			case "method":
				m.Method = vals[0]
			case "createTime":
				m.CreateTime = vals[0]
			case "updateTime":
				m.UpdateTime = vals[0]
			case "name":
				m.Name = vals[0]
			}
		}
	}
	return nil
}

// Span implement urlenc.URLValuesUnmarshaler.
func (m *Span) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "traceId":
				m.TraceId = vals[0]
			case "operationName":
				m.OperationName = vals[0]
			case "startTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.StartTime = val
			case "endTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.EndTime = val
			case "parentSpanId":
				m.ParentSpanId = vals[0]
			case "timestamp":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Timestamp = val
			case "duration":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Duration = val
			case "selfDuration":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.SelfDuration = val
			}
		}
	}
	return nil
}

// Trace implement urlenc.URLValuesUnmarshaler.
func (m *Trace) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "duration":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Duration = val
			case "services":
				m.Services = vals
			case "startTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.StartTime = val
			case "spanCount":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.SpanCount = val
			}
		}
	}
	return nil
}

// SpanEvent implement urlenc.URLValuesUnmarshaler.
func (m *SpanEvent) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "timestamp":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Timestamp = val
			}
		}
	}
	return nil
}

// SpanEventRequest implement urlenc.URLValuesUnmarshaler.
func (m *SpanEventRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "spanID":
				m.SpanID = vals[0]
			case "startTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.StartTime = val
			}
		}
	}
	return nil
}

// SpanEventResponse implement urlenc.URLValuesUnmarshaler.
func (m *SpanEventResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

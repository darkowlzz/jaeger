// Copyright (c) 2021 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package querysvc

import (
	"context"
	"errors"
	"time"

	"github.com/jaegertracing/jaeger/proto-gen/api_v2/metrics"
	"github.com/jaegertracing/jaeger/storage/metricsstore"
)

// MetricsQueryService contains the underlying reader required for querying the metrics store.
type MetricsQueryService struct {
	metricsReader metricsstore.Reader
}

var errNilReader = errors.New("no reader defined for MetricsQueryService")

// NewMetricsQueryService returns a new MetricsQueryService.
// A nil reader will result in a nil MetricsQueryService being returned.
func NewMetricsQueryService(reader metricsstore.Reader) *MetricsQueryService {
	return &MetricsQueryService{
		metricsReader: reader,
	}
}

// GetLatencies is the queryService implementation of metricsstore.Reader.
func (mqs MetricsQueryService) GetLatencies(ctx context.Context, params *metricsstore.LatenciesQueryParameters) (*metrics.MetricFamily, error) {
	if mqs.metricsReader == nil {
		return nil, errNilReader
	}
	return mqs.metricsReader.GetLatencies(ctx, params)
}

// GetCallRates is the queryService implementation of metricsstore.Reader.
func (mqs MetricsQueryService) GetCallRates(ctx context.Context, params *metricsstore.CallRateQueryParameters) (*metrics.MetricFamily, error) {
	if mqs.metricsReader == nil {
		return nil, errNilReader
	}
	return mqs.metricsReader.GetCallRates(ctx, params)
}

// GetErrorRates is the queryService implementation of metricsstore.Reader.
func (mqs MetricsQueryService) GetErrorRates(ctx context.Context, params *metricsstore.ErrorRateQueryParameters) (*metrics.MetricFamily, error) {
	if mqs.metricsReader == nil {
		return nil, errNilReader
	}
	return mqs.metricsReader.GetErrorRates(ctx, params)
}

// GetMinStepDuration is the queryService implementation of metricsstore.Reader.
func (mqs MetricsQueryService) GetMinStepDuration(ctx context.Context, params *metricsstore.MinStepDurationQueryParameters) (time.Duration, error) {
	if mqs.metricsReader == nil {
		return 0, errNilReader
	}
	return mqs.metricsReader.GetMinStepDuration(ctx, params)
}

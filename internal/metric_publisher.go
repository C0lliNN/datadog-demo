package internal

import (
	"log"

	"github.com/DataDog/datadog-go/statsd"
)

const (
	metricPrefix = "datadog_demo."
)

type MetricPublisher struct {
	client *statsd.Client 
}

func NewMetricPublisher(client *statsd.Client) *MetricPublisher {
	return &MetricPublisher{client: client}
}

func (p *MetricPublisher) PublishOrderCreated(order *Order) {
	if err := p.client.Incr(metricPrefix + "orders.created.count", nil, 1); err != nil {
		log.Println("failed to publish orders.created.count metric: ", err)
	}

	if err := p.client.Count(metricPrefix + "items.created.count", int64(len(order.Items)), nil, 1); err != nil {
		log.Println("failed to publish items.count metric: ", err)
	}

	if err := p.client.Count(metricPrefix + "tpv", int64(order.TotalPrice()), nil, 1); err != nil {
		log.Println("failed to publish tpv metric: ", err)
	}
}

func (p *MetricPublisher) PublishOrderRetrieved() {
	if err := p.client.Incr(metricPrefix + "orders.retrieved.count", nil, 1); err != nil {
		log.Println("failed to publish orders.retrieved.count metric: ", err)
	}
}


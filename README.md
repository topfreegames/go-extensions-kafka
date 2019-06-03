# go-extensions-kafka

This go library is a wrapper over [github.com/Shopify/sarama](https://github.com/Shopify/sarama) that adds OpenTracing support through a new method WithContext.

All the other methods are proxies to the wrapped instance.

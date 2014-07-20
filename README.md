kafka console
=============

Producer
--------

KC_TOPIC=topic KC_BROKER=host:port KC_PARTITION=0 go run producer.go

Consumer
--------

KC_TOPIC=topic KC_BROKER=host:port KC_PARTITION=0 go run consumer/consumer.go

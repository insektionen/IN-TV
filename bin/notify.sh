#!/bin/bash

# Tells all TVs to force reload website
mosquitto_pub -t "kistan/slideshow/reload" -m ""
